/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/hoehwa/jongi/cmd"
	_ "github.com/hoehwa/jongi/cmd/fndmn"
	_ "github.com/hoehwa/jongi/cmd/law"
	_ "github.com/hoehwa/jongi/cmd/paradigm"
	_ "github.com/hoehwa/jongi/cmd/principle"
)

func main() {
	cmd.Execute()
}
