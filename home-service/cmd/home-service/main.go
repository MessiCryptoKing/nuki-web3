package main

import (
	"fmt"
	"os"

	"github.com/messicryptoking/nuki-web3/home-service/cmd/home-service/webserver"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "homeservice",
	Short: "Nuki home service",
}

func main() {
	rootCmd.AddCommand(webserver.Webservercommand)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
