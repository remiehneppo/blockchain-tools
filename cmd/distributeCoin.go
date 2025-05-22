/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// distributeCoinCmd represents the distributeCoin command
var distributeCoinCmd = &cobra.Command{
	Use:   "distributeCoin",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("distributeCoin called")
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
	distributeCoinCmd.Flags().StringP("input", "i", "wallet.csv", "input file")
	distributeCoinCmd.Flags().StringP("token", "t", "", "token")
	distributeCoinCmd.Flags().StringP("network", "n", "", "network")
	distributeCoinCmd.Flags().StringP("min", "m", "0.01", "min amount")
	distributeCoinCmd.Flags().StringP("max", "x", "0.1", "max amount")
	distributeCoinCmd.Flags().IntP("num", "n", 10, "number of subaccounts")
}
