/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package h2l

import (
	"fmt"

	"github.com/spf13/cobra"
)

// high2LowCmd represents the high2Low command
var High2LowCmd = &cobra.Command{
	Use:   "high2Low",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("high2Low called")
	},
}

func init() {
	// rootCmd.AddCommand(high2LowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// high2LowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// high2LowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
