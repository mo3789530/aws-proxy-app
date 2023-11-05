/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"aws-proxy-app/cmd"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	cmd.Execute()
}
