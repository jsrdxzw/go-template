package actions

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

const Service = "example"

var (
	rootCmd = &cobra.Command{Use: Service}
)

func init() {
	rootCmd.PersistentFlags().String("env", os.Getenv("ENV"), "environment. eg: dev, prod")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
