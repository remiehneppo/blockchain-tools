/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/remiehneppo/go-binance-api/types"
	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip39"
)

// generate-walletsCmd represents the generate-wallets command
var generateWalletsCmd = &cobra.Command{
	Use:   "generate-wallets",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Here you can add the logic to generate wallets
		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			fmt.Println("Error getting num flag:", err)
			return
		}
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println("Error getting output flag:", err)
			return
		}

		fmt.Printf("Generating %d wallets and saving to %s\n", num, output)
		outputFile, err := os.Create(output + ".csv")
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer outputFile.Close()
		// Write the header to the CSV file
		_, err = outputFile.WriteString("Mnemonic,PrivateKey,Address\n")
		if err != nil {
			fmt.Println("Error writing header to output file:", err)
			return
		}

		for i := 0; i < num; i++ {
			entropy, err := bip39.NewEntropy(256)
			if err != nil {
				fmt.Println("Error generating entropy:", err)
				return
			}
			mnemonic, err := bip39.NewMnemonic(entropy)
			if err != nil {
				fmt.Println("Error generating mnemonic:", err)
				return
			}
			// generate the private key and address from the mnemonic
			wallet, err := hdwallet.NewFromMnemonic(mnemonic)
			if err != nil {
				fmt.Println("Error creating wallet from mnemonic:", err)
				return
			}
			path := hdwallet.MustParseDerivationPath(types.HD_PATH)
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
			_, err = outputFile.WriteString(fmt.Sprintf("%s,%s,%s\n", mnemonic, privateKey, address))
			if err != nil {
				fmt.Println("Error writing wallet to output file:", err)
				return
			}

		}

		fmt.Println("Wallets generated successfully")
	},
}

func init() {
	rootCmd.AddCommand(generateWalletsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateWalletsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateWalletsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateWalletsCmd.Flags().IntP("num", "n", 10, "Number of wallets to generate")
	generateWalletsCmd.Flags().StringP("output", "o", "wallets", "Output file for generated wallets")
}
