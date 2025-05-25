/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/remiehneppo/go-binance-api/contracts/launchpad"
	"github.com/remiehneppo/go-binance-api/utils"
	"github.com/spf13/cobra"
)

// callClaimCmd represents the callClaim command
var callClaimCmd = &cobra.Command{
	Use:   "call-claim",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("call-claim called")
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
		outputFile.WriteString("address,tx_hash,success\n") // Write header to output file
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
			privateKey, ok := csvExporter.GetCell("private_key", row)
			if !ok {
				fmt.Println("Error getting private key from CSV file")
				continue
			}
			transactOpts, account, err := utils.PrepareTransactor(
				cmd.Context(),
				*client,
				privateKey,
				chainId,
			)
			if err != nil {
				fmt.Println("Error preparing transaction options:", err)
				continue
			}
			fmt.Printf("Processing row %d for address %s\n", row, account.Hex())

			tx, err := contractInstance.Claim(transactOpts)
			if err != nil {
				fmt.Println("Error calling claim function:", err)
				continue
			}

			fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
			txReceipt, err := bind.WaitMined(cmd.Context(), client, tx)
			if err != nil {
				fmt.Println("Error waiting for transaction to be mined:", err)
				continue
			}
			_, err = outputFile.WriteString(fmt.Sprintf("%s,%s,%t\n",
				account.Hex(),
				tx.Hash().Hex(),
				txReceipt.Status == 1,
			))
			if err != nil {
				fmt.Println("Error writing to output file:", err)
				continue
			}
			fmt.Printf("Transaction mined: %s, success: %t\n", tx.Hash().Hex(), txReceipt.Status == 1)
		}

	},
}

func init() {
	rootCmd.AddCommand(callClaimCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// callClaimCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// callClaimCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	callClaimCmd.Flags().StringP("input", "i", "", "Input file with addresses to claim")
	callClaimCmd.Flags().StringP("output", "o", "claim_output.csv", "Output file to save results")
	callClaimCmd.Flags().StringP("contract", "c", "", "Contract address for the claim")
	callClaimCmd.Flags().StringP("rpc", "r", "", "RPC URL for the Ethereum client")
}
