package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TransferNative(
	privateKey string,
	dest string,
	amount *big.Int,
	client *ethclient.Client,
) (string, error) {
	// Parse private key
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %v", err)
	}

	// Get the public address from the private key
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Parse destination address
	toAddress := common.HexToAddress(dest)

	// Get nonce for the sender address
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %v", err)
	}

	// Get current gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to suggest gas price: %v", err)
	}
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(2)) // Double the gas price for safety
	// Calculate gas limit (use default value for native token transfer)
	gasLimit := uint64(30000)

	// Calculate transaction fee
	txFee := new(big.Int).Mul(gasPrice, big.NewInt(int64(gasLimit)))
	if amount == nil {
		// Get account balance
		balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
		if err != nil {
			return "", fmt.Errorf("failed to get balance: %v", err)
		}

		// Check if balance is sufficient
		if balance.Cmp(txFee) <= 0 {
			return "", fmt.Errorf("insufficient balance for transaction fee")
		}

		// Calculate amount to send (balance - txFee)
		amount = new(big.Int).Sub(balance, txFee)
	}

	// Create transaction
	tx := ethtypes.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)

	// Get chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get chain ID: %v", err)
	}

	// Sign transaction
	signedTx, err := ethtypes.SignTx(tx, ethtypes.NewEIP155Signer(chainID), privateKeyECDSA)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Send transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %v", err)
	}

	return signedTx.Hash().Hex(), nil
}

func TransferToken(
	privatekey string,
	dest string,
	tokenAddress string,
	amount *big.Int,
	client *ethclient.Client,
) (string, error) {
	// Parse private key
	privateKeyECDSA, err := crypto.HexToECDSA(privatekey)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %v", err)
	}

	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	toAddress := common.HexToAddress(dest)
	tokenContractAddress := common.HexToAddress(tokenAddress)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(2)) // Double the gas price for safety

	if amount == nil {
		balanceOfFnSignature := []byte("balanceOf(address)")
		balanceMethodID := crypto.Keccak256(balanceOfFnSignature)[:4]
		paddedFromAddress := common.LeftPadBytes(fromAddress.Bytes(), 32)
		balanceData := append(balanceMethodID, paddedFromAddress...)
		// Make a call to the token contract to get the balance
		result, err := client.CallContract(context.Background(), ethereum.CallMsg{
			To:   &tokenContractAddress,
			Data: balanceData,
		}, nil)
		if err != nil {
			return "", fmt.Errorf("failed to get token balance: %v", err)
		}
		amount = new(big.Int).SetBytes(result)
		if amount.Cmp(big.NewInt(0)) == 0 {
			return "", fmt.Errorf("insufficient token balance")
		}
	}

	transferFnSignature := []byte("transfer(address,uint256)")
	methodID := crypto.Keccak256(transferFnSignature)[:4]

	// Pad address to 32 bytes
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	// Pad amount to 32 bytes
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// Estimate gas limit
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     fromAddress,
		To:       &tokenContractAddress,
		Gas:      0,
		GasPrice: gasPrice,
		Value:    big.NewInt(0),
		Data:     data,
	})
	if err != nil {
		return "", fmt.Errorf("failed to estimate gas limit: %v", err)
	}
	gasLimit = gasLimit + (gasLimit / 5) // Add 20%
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get balance: %v", err)
	}
	gasCost := new(big.Int).Mul(gasPrice, big.NewInt(int64(gasLimit)))
	if balance.Cmp(gasCost) < 0 {
		return "", fmt.Errorf("insufficient funds for gas: have %s, need %s",
			balance.String(), gasCost.String())
	}

	// Create transaction
	tx := ethtypes.NewTransaction(
		nonce,
		tokenContractAddress,
		big.NewInt(0), // No native tokens being sent
		gasLimit,
		gasPrice,
		data,
	)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get chain ID: %v", err)
	}

	// Sign transaction
	signedTx, err := ethtypes.SignTx(tx, ethtypes.NewEIP155Signer(chainID), privateKeyECDSA)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Send transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %v", err)
	}

	return signedTx.Hash().Hex(), nil
}

func CallMethod(
	privateKey string,
	contractAddress string,
	methodName string,
	methodArgs []interface{},
	valueToSend *big.Int,
	client *ethclient.Client,
) (string, error) {
	// Parse private key
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %v", err)
	}

	// Get the public address from the private key
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Parse contract address
	contractAddr := common.HexToAddress(contractAddress)

	// Get nonce for the sender address
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %v", err)
	}

	// Get current gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to suggest gas price: %v", err)
	}
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(2)) // Double the gas price for safety

	// Build the method signature
	var methodSig string
	if len(methodArgs) > 0 {
		argTypes := make([]string, len(methodArgs))
		for i, arg := range methodArgs {
			switch arg.(type) {
			case string:
				if common.IsHexAddress(arg.(string)) {
					argTypes[i] = "address"
				} else {
					argTypes[i] = "string"
				}
			case *big.Int, int, int64, uint, uint64:
				argTypes[i] = "uint256"
			case bool:
				argTypes[i] = "bool"
			case []byte:
				argTypes[i] = "bytes"
			default:
				return "", fmt.Errorf("unsupported argument type: %T", arg)
			}
		}
		methodSig = fmt.Sprintf("%s(%s)", methodName, strings.Join(argTypes, ","))
	} else {
		methodSig = fmt.Sprintf("%s()", methodName)
	}

	// Create method ID
	methodID := crypto.Keccak256([]byte(methodSig))[:4]

	// Encode arguments
	var data []byte
	data = append(data, methodID...)

	for _, arg := range methodArgs {
		var encodedArg []byte

		switch v := arg.(type) {
		case string:
			if common.IsHexAddress(v) {
				// Handle address type
				addr := common.HexToAddress(v)
				encodedArg = common.LeftPadBytes(addr.Bytes(), 32)
			} else {
				// Handle string type - this is more complex and requires dynamic types
				// For simplicity, we're not implementing full ABI encoding here
				return "", fmt.Errorf("string arguments not supported in this simplified implementation")
			}
		case *big.Int:
			encodedArg = common.LeftPadBytes(v.Bytes(), 32)
		case int, int64:
			val := big.NewInt(reflect.ValueOf(v).Int())
			encodedArg = common.LeftPadBytes(val.Bytes(), 32)
		case uint, uint64:
			val := big.NewInt(int64(reflect.ValueOf(v).Uint()))
			encodedArg = common.LeftPadBytes(val.Bytes(), 32)
		case bool:
			if v {
				encodedArg = common.LeftPadBytes([]byte{1}, 32)
			} else {
				encodedArg = common.LeftPadBytes([]byte{0}, 32)
			}
		default:
			return "", fmt.Errorf("unsupported argument type: %T", arg)
		}

		data = append(data, encodedArg...)
	}

	// Check account balance
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get balance: %v", err)
	}

	// Estimate gas limit
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     fromAddress,
		To:       &contractAddr,
		Gas:      0,
		GasPrice: gasPrice,
		Value:    valueToSend,
		Data:     data,
	})
	if err != nil {
		return "", fmt.Errorf("failed to estimate gas limit: %v", err)
	}

	// Add 20% buffer to gas limit
	gasLimit = gasLimit + (gasLimit / 5)

	// Calculate total cost (gas fee + value to send)
	gasCost := new(big.Int).Mul(gasPrice, big.NewInt(int64(gasLimit)))
	totalCost := new(big.Int).Add(gasCost, valueToSend)

	// Check if balance is sufficient
	if balance.Cmp(totalCost) < 0 {
		return "", fmt.Errorf("insufficient balance for transaction: have %s, need %s",
			balance.String(), totalCost.String())
	}

	// Create transaction
	tx := ethtypes.NewTransaction(
		nonce,
		contractAddr,
		valueToSend,
		gasLimit,
		gasPrice,
		data,
	)

	// Get chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get chain ID: %v", err)
	}

	// Sign transaction
	signedTx, err := ethtypes.SignTx(tx, ethtypes.NewEIP155Signer(chainID), privateKeyECDSA)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Send transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %v", err)
	}

	return signedTx.Hash().Hex(), nil
}

func PrepareTransactor(ctx context.Context, client ethclient.Client, privateKey string, chainId *big.Int) (*bind.TransactOpts, *common.Address, error) {

	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid private key: %v", err)
	}
	address := crypto.PubkeyToAddress(ecdsaPrivateKey.PublicKey)
	transactOpts, err := bind.NewKeyedTransactorWithChainID(
		ecdsaPrivateKey,
		chainId,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating transactor: %v", err)
	}
	transactOpts.Context = ctx
	transactOpts.GasLimit = 1000000 // Set a reasonable gas limit
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("error suggesting gas price")
	}
	transactOpts.GasPrice = new(big.Int).Mul(gasPrice, big.NewInt(1))

	return transactOpts, &address, nil
}
