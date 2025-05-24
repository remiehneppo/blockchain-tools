/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/remiehneppo/go-binance-api/utils"
	"github.com/spf13/cobra"
)

// distributeCoinCmd represents the distributeCoin command
var distributeCoinCmd = &cobra.Command{
	Use:   "distribute-coin",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("distribute-coin called")
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
		token, err := cmd.Flags().GetString("token")
		if err != nil {
			fmt.Println("Error getting token flag:", err)
			return
		}
		rpc, err := cmd.Flags().GetString("rpc")
		if err != nil {
			fmt.Println("Error getting rpc")
			return
		}
		min, err := cmd.Flags().GetString("min")
		if err != nil {
			fmt.Println("Error getting min")
			return
		}
		max, err := cmd.Flags().GetString("max")
		if err != nil {
			fmt.Println("Error getting max")
			return
		}

		minAmount, ok := new(big.Int).SetString(min, 10)
		if !ok {
			fmt.Println("Error converting min to big.Int")
			return
		}
		maxAmount, ok := new(big.Int).SetString(max, 10)
		if !ok {
			fmt.Println("Error converting max to big.Int")
			return
		}

		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			fmt.Println("Error getting num flag:", err)
			return
		}

		fmt.Printf("Distributing %s to %s with min %s and max %s\n", token, input, min, max)
		outputFile, err := os.Create(output)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer outputFile.Close()
		// Write the header to the CSV file
		_, err = outputFile.WriteString("ID,Mnemonic,Index,PrivateKey,Address,Token,Amount\n")
		if err != nil {
			fmt.Println("Error writing header to output file:", err)
			return
		}

		// Here you can add the logic to distribute coins
		csvExporter := utils.NewCSVExporter(input)
		client, err := ethclient.DialContext(cmd.Context(), rpc)
		if err != nil {
			fmt.Println("Error connecting to Ethereum client:", err)
			return
		}
		defer client.Close()

		for row := range csvExporter.GetRowsCount() {
			mnemonic, ok := csvExporter.GetCell("mnemonic", row)
			if !ok {
				fmt.Println("Error getting mnemonic from CSV file")
				continue
			}
			parentKey, _, err := utils.CreateHDWalletByMnemonic(mnemonic, 0)
			if err != nil {
				fmt.Println("Error creating parent wallet from mnemonic:", err)
				continue
			}
			for j := 1; j <= num; j++ {
				childKey, childAddress, err := utils.CreateHDWalletByMnemonic(mnemonic, j)
				if err != nil {
					fmt.Println("Error creating child wallet from mnemonic:", err)
					continue
				}
				amount := utils.RandomBigInt(minAmount, maxAmount)
				if token == "" {
					txHash, err := utils.TransferNative(
						parentKey,
						childAddress,
						amount,
						client,
					)
					if err != nil {
						fmt.Printf("Error transferring native with account at row %d\n", row)
						continue
					}
					fmt.Printf("Transfer native successfully, tx hash: %s\n", txHash)
				} else {
					txHash, err := utils.TransferToken(
						parentKey,
						childAddress,
						token,
						amount,
						client,
					)
					if err != nil {
						fmt.Printf("Error transferring %s with account at row %d\n", token, row)
						continue
					}
					fmt.Printf("Transfer %s successfully, tx hash: %s\n", token, txHash)
				}
				// Write the wallet information to the CSV file
				_, err = outputFile.WriteString(fmt.Sprintf("%d,%s,%d,%s,%s,%s,%s\n", row, mnemonic, j, childKey, childAddress, token, amount.String()))
				if err != nil {
					fmt.Println("Error writing wallet to output file:", err)
					continue
				}

			}

		}
	},
}

func init() {
	rootCmd.AddCommand(distributeCoinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// distributeCoinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// distributeCoinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	distributeCoinCmd.Flags().StringP("input", "i", "wallets.csv", "input file")
	distributeCoinCmd.Flags().StringP("token", "t", "", "token")
	distributeCoinCmd.Flags().StringP("rpc", "r", "", "network")
	distributeCoinCmd.Flags().StringP("min", "m", "1000000000000000", "min amount")
	distributeCoinCmd.Flags().StringP("max", "x", "2000000000000000", "max amount")
	distributeCoinCmd.Flags().Int("num", 10, "number of subaccounts")
	distributeCoinCmd.Flags().StringP("output", "o", "distribute.csv", "output file")
}
