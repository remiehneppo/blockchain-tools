/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/remiehneppo/go-binance-api/utils"
	"github.com/spf13/cobra"
)

// allTransferCmd represents the allTransfer command
var allTransferCmd = &cobra.Command{
	Use:   "all-transfer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("all-transfer called")
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			fmt.Println("Error getting input flag:", err)
			return
		}
		address, err := cmd.Flags().GetString("address")
		if err != nil {
			fmt.Println("error get address destination")
			return
		}
		token, err := cmd.Flags().GetString("token")
		if err != nil {
			fmt.Println("Error getting token flag:", err)
			return
		}
		rpc, err := cmd.Flags().GetString("rpc")
		if err != nil {
			fmt.Println("Error getting rpc flag:", err)
			return
		}

		client, err := ethclient.DialContext(cmd.Context(), rpc)
		if err != nil {
			panic(err)
		}

		csvExporter := utils.NewCSVExporter(input)

		for row := range csvExporter.GetRowsCount() {
			privateKey, ok := csvExporter.GetCell("PrivateKey", row)
			if !ok {
				fmt.Println("error get private key at row", row)
			}
			if token == "" {
				txHash, err := utils.TransferNative(
					privateKey,
					address,
					nil,
					client,
				)
				if err != nil {
					fmt.Printf("error transfer native with account at row %d\n", row)
					continue
				}
				fmt.Printf("transfer native successfully, tx hash: %s\n", txHash)
			} else {
				txHash, err := utils.TransferToken(
					privateKey,
					address,
					token,
					nil,
					client,
				)
				if err != nil {
					fmt.Printf("error transfer %s with account at row %d\n", token, row)
					continue
				}
				fmt.Printf("transfer %s successfully, tx hash: %s\n", token, txHash)
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(allTransferCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allTransferCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allTransferCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	allTransferCmd.Flags().StringP("input", "i", "", "input file")
	allTransferCmd.Flags().StringP("token", "t", "", "token")
	allTransferCmd.Flags().StringP("rpc", "r", "", "rpc endpoint")
	allTransferCmd.Flags().StringP("address", "a", "", "address")
}
