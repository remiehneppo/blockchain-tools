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

// checkTokenDexCmd represents the checkTokenDex command
var checkTokenDexCmd = &cobra.Command{
	Use:   "check-token-dex",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check-token-dex called")
		token, err := cmd.Flags().GetString("token")
		if err != nil {
			fmt.Println("Error getting token flag:", err)
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

		res, err := client.NewGetAllCoinsInfoService().Do(cmd.Context())
		if err != nil {
			fmt.Println("Error connecting to Binance API:", err)
			return
		}

		for _, coin := range res {
			if coin.Coin == token {
				fmt.Println("Coin:", coin.Coin)
				fmt.Println("Name:", coin.Name)
				for _, network := range coin.NetworkList {
					fmt.Println("================")
					fmt.Println("Network:", network.Network)
					fmt.Println("Withdraw Fee:", network.WithdrawFee)
					fmt.Println("Min Withdraw Amount:", network.WithdrawMin)
					fmt.Println("Is default:", network.IsDefault)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkTokenDexCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkTokenDexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkTokenDexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	checkTokenDexCmd.Flags().StringP("token", "t", "", "token")
}
