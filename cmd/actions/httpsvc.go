package actions

import (
	service "example/be/service/api"
	"github.com/spf13/cobra"
	"log"
)

func RunHTTPSvc(cmd *cobra.Command, args []string) {
	env, err := cmd.Flags().GetString("env")
	if err != nil {
		log.Fatalf("get env failed: %v\n", err)
	}
	service.StartHTTP(env)
}

func init() {
	cmd := &cobra.Command{
		Use: "httpsvc",
		Run: RunHTTPSvc,
	}
	rootCmd.AddCommand(cmd)
}
