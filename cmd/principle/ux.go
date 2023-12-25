/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package principle

import (
	utills "github.com/hoehwa/jongi/utills/print.go"
	"github.com/spf13/cobra"
)

// uxCmd represents the ux command.
var uxCmd = &cobra.Command{
	Use:   "ux",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utills.PrintContent("/principle/ux")
	},
}

func init() {
	principleCmd.AddCommand(uxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
