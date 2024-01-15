/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package paradigm

import (
	"github.com/hoehwa/jongi/internal"
	"github.com/spf13/cobra"
)

// FPCmd represents the FP command.
var FPCmd = &cobra.Command{
	Use:   "FP",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.ShowListAt("/paradigm/FP")
	},
}

func init() {
	paradigmCmd.AddCommand(FPCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// FPCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// FPCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
