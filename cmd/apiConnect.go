/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	binance_connector "github.com/binance/binance-connector-go"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// apiConnectCmd represents the apiConnect command
var apiConnectCmd = &cobra.Command{
	Use:   "api-connect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load environment variables from .env file
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")
			return
		}

		apiKey := os.Getenv("API_KEY")
		secretKey := os.Getenv("SECRET_KEY")

		client := binance_connector.NewClient(apiKey, secretKey)

		client.Debug = true

		err = client.NewPingService().Do(cmd.Context())
		if err != nil {
			fmt.Println("Error connecting to Binance API:", err)
			return
		}
		fmt.Println("Successfully connected to Binance API")
		walletBalance, err := client.NewWalletBalanceService().Do(cmd.Context())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(binance_connector.PrettyPrint(walletBalance))
	},
}

func init() {
	rootCmd.AddCommand(apiConnectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiConnectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiConnectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
