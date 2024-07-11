// Package cmd /*
package cmd

import (
	"github.com/gin-gonic/gin"
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Message": "Hello World"})
	})

	err := r.Run(":8082")
	if err != nil {
		return
	}
}
