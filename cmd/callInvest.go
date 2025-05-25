/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/remiehneppo/go-binance-api/contracts/launchpad"
	"github.com/remiehneppo/go-binance-api/utils"
	"github.com/spf13/cobra"
)

// callInvestCmd represents the callInvest command
var callInvestCmd = &cobra.Command{
	Use:   "call-invest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("call-invest called")
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			fmt.Println("Error getting input flag:", err)
			return
		}
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println("Error getting output flag:", err)
			return
		}
		outputFile, err := os.Create(output)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer outputFile.Close()
		outputFile.WriteString("address,amount,tx_hash,success\n") // Write header to output file
		contract, err := cmd.Flags().GetString("contract")
		if err != nil {
			fmt.Println("Error getting contract flag:", err)
			return
		}
		rpc, err := cmd.Flags().GetString("rpc")
		if err != nil {
			fmt.Println("Error getting rpc flag:", err)
			return
		}
		client, err := ethclient.DialContext(cmd.Context(), rpc)
		if err != nil {
			fmt.Println("Error connecting to Ethereum client:", err)
			return
		}
		chainId, err := client.ChainID(cmd.Context())
		if err != nil {
			fmt.Println("Error getting chain ID:", err)
			return
		}

		contractInstance, err := launchpad.NewLaunchpad(common.HexToAddress(contract), client)

		csvExporter := utils.NewCSVExporter(input)
		for row := range csvExporter.GetRowsCount() {
			privateKeyStr, ok := csvExporter.GetCell("privatekey", row)
			if !ok {
				fmt.Println("Error getting private key from input file")
				continue
			}
			privateKey, err := crypto.HexToECDSA(privateKeyStr)
			if err != nil {
				fmt.Println("Error converting private key:", err)
				continue
			}
			amount, ok := csvExporter.GetCell("amount", row)
			if !ok {
				fmt.Println("Error getting amount from input file")
				continue
			}
			pubKey := crypto.PubkeyToAddress(privateKey.PublicKey)
			fmt.Printf("Processing row %d for address %s\n", row, pubKey.Hex())
			transactOpts, err := bind.NewKeyedTransactorWithChainID(
				privateKey,
				chainId,
			)
			if err != nil {
				fmt.Println("Error creating transaction options:", err)
				continue
			}
			transactOpts.Context = cmd.Context()
			transactOpts.GasLimit = 1000000 // Set a reasonable gas limit
			gasPrice, err := client.SuggestGasPrice(cmd.Context())
			if err != nil {
				fmt.Println("Error suggesting gas price:", err)
				continue
			}
			transactOpts.GasPrice = new(big.Int).Mul(gasPrice, big.NewInt(1))
			transactOpts.Value, ok = new(big.Int).SetString(amount, 10)
			if !ok {
				fmt.Println("Error converting amount to big.Int:", amount)
				continue
			}

			tx, err := contractInstance.Invest(transactOpts)
			if err != nil {
				fmt.Println("Error calling Invest function:", err)
				continue
			}
			fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
			// check if the transaction was executed successfully
			txReceipt, err := bind.WaitMined(cmd.Context(), client, tx)
			if err != nil {
				fmt.Println("Error waiting for transaction to be mined:", err)
				continue
			}
			outputFile.WriteString(fmt.Sprintf("%s,%s,%s,%t\n",
				pubKey.Hex(),
				amount,
				tx.Hash().Hex(),
				txReceipt.Status == 1, // 1 means success, 0 means failure
			))
			// Here you can add logic to write the transaction hash to the output file

		}

	},
}

func init() {
	rootCmd.AddCommand(callInvestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// callInvestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// callInvestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	callInvestCmd.Flags().StringP("input", "i", "", "Input file")
	callInvestCmd.Flags().StringP("output", "o", "invest_output.csv", "Output file")
	callInvestCmd.Flags().StringP("contract", "c", "", "Contract address")
	callInvestCmd.Flags().StringP("rpc", "r", "", "RPC URL")
}
