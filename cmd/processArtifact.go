/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// processArtifactCmd represents the processArtifact command
var processArtifactCmd = &cobra.Command{
	Use:   "processArtifact",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("processArtifact called")
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			fmt.Println("Error getting input flag:", err)
			return
		}
		data, err := os.ReadFile(input)
		if err != nil {
			fmt.Println("Error reading input file:", err)
			return
		}

		var artifact map[string]interface{}
		if err := json.Unmarshal(data, &artifact); err != nil {
			fmt.Println("Error unmarshalling input file:", err)
			return
		}
		abi := artifact["abi"]
		if abi == nil {
			fmt.Println("ABI not found in the input file")
			return
		}
		abiJson, err := json.Marshal(abi)
		if err != nil {
			fmt.Println("Error marshalling ABI:", err)
			return
		}
		artifact["abi"] = string(abiJson)
		// overwrite the input file with the processed artifact
		outputString, err := json.MarshalIndent(artifact, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling processed artifact:", err)
			return
		}
		if err := os.WriteFile(input, outputString, 0644); err != nil {
			fmt.Println("Error writing processed artifact to file:", err)
			return
		}
		fmt.Println("Processed artifact successfully written to", input)

	},
}

func init() {
	rootCmd.AddCommand(processArtifactCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// processArtifactCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// processArtifactCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	processArtifactCmd.Flags().StringP("input", "i", "", "Input file path")
}
