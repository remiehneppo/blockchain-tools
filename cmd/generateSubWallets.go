/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/remiehneppo/go-binance-api/types"
	"github.com/remiehneppo/go-binance-api/utils"
	"github.com/spf13/cobra"
)

// generate-sub-walletsCmd represents the generate-sub-wallets command
var generateSubWalletsCmd = &cobra.Command{
	Use:   "generate-sub-wallets",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Here you can add the logic to generate sub wallets
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
		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			fmt.Println("Error getting num flag:", err)
			return
		}

		fmt.Printf("Generating %d sub wallets from %s and saving to %s\n", num, input, output)
		outputFile, err := os.Create(output + ".csv")
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer outputFile.Close()

		// Write the header to the CSV file
		_, err = outputFile.WriteString("ID,Mnemonic,Index,PrivateKey,Address\n")
		if err != nil {
			fmt.Println("Error writing header to output file:", err)
			return
		}
		csvExporter := utils.NewCSVExporter(input)

		for row := range csvExporter.GetRowsCount() {

			mnemonic, ok := csvExporter.GetCell("mnemonic", row)
			if !ok {
				fmt.Println("Error getting mnemonic from CSV file")
				return
			}

			for j := 0; j < num; j++ {
				wallet, err := hdwallet.NewFromMnemonic(mnemonic)
				if err != nil {
					fmt.Println("Error creating wallet from mnemonic:", err)
					return
				}

				path := hdwallet.MustParseDerivationPath(fmt.Sprintf("%s%d", types.HD_BASE_PATH, j))
				account, err := wallet.Derive(path, false)
				if err != nil {
					fmt.Println("Error deriving account:", err)
					return
				}
				privateKey, err := wallet.PrivateKeyHex(account)
				if err != nil {
					fmt.Println("Error getting private key:", err)
					return
				}
				address, err := wallet.AddressHex(account)
				if err != nil {
					fmt.Println("Error getting address:", err)
					return
				}
				// Write the wallet information to the CSV file
				_, err = outputFile.WriteString(fmt.Sprintf("%d,%s,%d,%s,%s\n", row, mnemonic, j, privateKey, address))
				if err != nil {
					fmt.Println("Error writing wallet to output file:", err)
					return
				}

			}
		}
		// Add your logic to generate sub wallets here
		fmt.Printf("Using HD path: %s\n", types.HD_BASE_PATH)
	},
}

func init() {
	rootCmd.AddCommand(generateSubWalletsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateSubWalletsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateSubWalletsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateSubWalletsCmd.Flags().StringP("input", "i", "wallets.csv", "Input file with root wallets")
	generateSubWalletsCmd.Flags().StringP("output", "o", "sub_wallets", "Output file for sub wallets")
	generateSubWalletsCmd.Flags().IntP("num", "n", 10, "Number of sub wallets to generate")
}
