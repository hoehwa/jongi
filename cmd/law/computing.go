/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package law

import (
	utills "github.com/hoehwa/jongi/utills/print.go"
	"github.com/spf13/cobra"
)

// computingCmd represents the computing command.
var computingCmd = &cobra.Command{
	Use:   "computing",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utills.PrintContent("/law/computing")
	},
}

func init() {
	lawCmd.AddCommand(computingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// computingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// computingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
