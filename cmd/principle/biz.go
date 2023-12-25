/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package principle

import (
	utills "github.com/hoehwa/jongi/utills/print.go"
	"github.com/spf13/cobra"
)

// bizCmd represents the biz command.
var bizCmd = &cobra.Command{
	Use:   "biz",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utills.PrintContent("/principle/biz")
	},
}

func init() {
	principleCmd.AddCommand(bizCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bizCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bizCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
