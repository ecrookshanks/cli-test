package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hello",
	Short: "Hello program",
	Long:  `hello`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Hello World!")
	// },
}

func init() {
	//rootCmd.Flags().StringVarP(&Message, "message", "m", "", "Hello message to print.")
}

func GetCmd() *cobra.Command {
	return rootCmd
}

func Execute(cmd *cobra.Command) {
	if err := cmd.Execute(); err != nil {
		fmt.Println("There was an error: " + err.Error())
		os.Exit(1)
	}
}
