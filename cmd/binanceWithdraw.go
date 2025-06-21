/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	binance_connector "github.com/binance/binance-connector-go"
	"github.com/joho/godotenv"
	"github.com/remiehneppo/go-binance-api/utils"
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
		fmt.Println("Output file created:", output)
		// Get the token and network flags
		outputFile.Write([]byte("address,amount,withdraw_id,error\n")) // Write header to output file

		token, err := cmd.Flags().GetString("token")
		if err != nil {
			fmt.Println("Error getting token flag:", err)
			return
		}
		network, err := cmd.Flags().GetString("network")
		if err != nil {
			fmt.Println("Error getting network flag:", err)
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

		csvExporter := utils.NewCSVExporter(input)

		for row := range csvExporter.GetRowsCount() {
			address, ok := csvExporter.GetCell("address", row)
			if !ok {
				fmt.Println("Error getting address cell")
				continue
			}
			amount, ok := csvExporter.GetCell("amount", row)
			if !ok {
				fmt.Println("Error getting amount cell")
				continue
			}
			amountFloat, err := strconv.ParseFloat(amount, 64)
			if err != nil {
				fmt.Println("error parse amount: ", err)
				return
			}
			fmt.Println("Withdraw request for address: ", address, " amount: ", amountFloat)
			withdraw, err := client.NewWithdrawService().Coin(token).Network(network).Address(address).
				Amount(amountFloat).Do(cmd.Context())
			if err != nil {
				outputFile.Write([]byte(fmt.Sprintf("%s,%f,,%s\n", address, amountFloat, err.Error())))
				fmt.Println("Error sending withdraw request:", err)
				continue
			}
			fmt.Println(binance_connector.PrettyPrint(withdraw))
			fmt.Println("Withdraw request sent successfully. address: ", address, " amount: ", amountFloat)
			fmt.Println("Withdraw ID: ", withdraw.Id)
			outputFile.Write([]byte(fmt.Sprintf("%s,%f,%s,\n", address, amountFloat, withdraw.Id)))
			// sleep for 5 seconds to avoid rate limit
			time.Sleep(5 * time.Second)
		}

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

	binanceWithdrawCmd.Flags().StringP("input", "i", "withdraw.csv", "Input file addresses")
	binanceWithdrawCmd.Flags().StringP("token", "t", "ETH", "Token symbol")
	binanceWithdrawCmd.Flags().StringP("network", "n", "BASE", "Network withdraw")
	binanceWithdrawCmd.Flags().StringP("output", "o", "withdraw_result.csv", "Output file for withdraw results")
}
