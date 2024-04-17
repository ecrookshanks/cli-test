package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Name string

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greet is a command to say hello",
	Long: `A Simple test to exercise the cobra libary,
	       and to say hello.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Name == "" {
			Name = "World!!"
		}
		msg := fmt.Sprintf("Hello %s!", Name)
		fmt.Println(msg)
	},
}

func init() {
	greetCmd.Flags().StringVarP(&Name, "name", "n", "", "Hello name to print.")
	rootCmd.AddCommand(greetCmd)
}
