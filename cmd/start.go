// Package cmd /*
package cmd

import (
	"eve/app/route"
	"eve/internal/server"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start serve",
	Long:  `start serve`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func run() {
	http := server.New()
	http.GenRouter(route.New())
	http.Run()
}
