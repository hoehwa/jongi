/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package law

import (
	utills "github.com/hoehwa/jongi/utills/print.go"
	"github.com/spf13/cobra"
)

// decideCmd represents the decide command.
var decideCmd = &cobra.Command{
	Use:   "decide",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utills.PrintContent("/law/decide")
	},
}

func init() {
	lawCmd.AddCommand(decideCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decideCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decideCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
