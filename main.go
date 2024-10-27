package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var name string
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "A simple CLI application",
	}

	var helloCmd = &cobra.Command{
		Use:   "hello",
		Short: "Prints Hello",
		Run: func(cmd *cobra.Command, args []string) {
			if name != "" {
				fmt.Printf("Hello, %s!\n", name)
			} else {
				fmt.Printf("Hello, World!")
			}
		},
	}

	helloCmd.Flags().StringVarP(&name, "name", "n", "world", "Name to greet")
	rootCmd.AddCommand(helloCmd)
	rootCmd.Execute()
}
