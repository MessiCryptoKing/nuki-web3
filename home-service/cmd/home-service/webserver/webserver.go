package webserver

import "github.com/spf13/cobra"

var Webservercommand = &cobra.Command{
	Use:   "server",
	Short: "web server functionality",
}

func init() {
	Webservercommand.AddCommand(start)
}
