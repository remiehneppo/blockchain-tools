/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	binance_connector "github.com/binance/binance-connector-go"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// binanceWithdrawCmd represents the binanceWithdraw command
var binanceWithdrawCmd = &cobra.Command{
	Use:   "binance-withdraw",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("binance-withdraw called")

		input, err := cmd.Flags().GetString("input")
		if err != nil {
			fmt.Println("Error getting input flag:", err)
			return
		}
		token, err := cmd.Flags().GetString("token")
		if err != nil {
			fmt.Println("Error getting input flag:", err)
			return
		}
		network, err := cmd.Flags().GetString("network")
		if err != nil {
			fmt.Println("Error getting input flag:", err)
			return
		}

		err = godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")
			return
		}

		apiKey := os.Getenv("API_KEY")
		secretKey := os.Getenv("SECRET_KEY")

		client := binance_connector.NewClient(apiKey, secretKey)

		inputBytes, err := os.ReadFile(input)
		if err != nil {
			fmt.Println("error read input: ", err)
		}
		lines := strings.Split(string(inputBytes), "\n")
	},
}

func init() {
	rootCmd.AddCommand(binanceWithdrawCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// binanceWithdrawCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// binanceWithdrawCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	apiConnectCmd.Flags().StringP("input", "i", "wallets.csv", "Input file addresses")
	apiConnectCmd.Flags().StringP("token", "t", "ETH", "Token symbol")
	apiConnectCmd.Flags().StringP("network", "n", "base", "Network withdraw")
}
