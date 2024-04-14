/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/bazdalaz/cpi/cmd"
	"github.com/bazdalaz/cpi/internal/config"
)

const (
	// Version of the application
	Version  = "0.0.1"
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoadConfig()
	fmt.Printf("Loaded config: %+v\n", cfg)

	cmd.Execute()
}
