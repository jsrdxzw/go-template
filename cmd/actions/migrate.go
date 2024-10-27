package actions

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"example/be/gen/entschema/migrate"
	"example/be/service/app"
	"github.com/spf13/cobra"
	"log"
)

func runMigrate(cmd *cobra.Command, args []string) {
	env, err := cmd.Flags().GetString("env")
	if err != nil {
		log.Fatalf("get env failed: %v\n", err)
	}

	envConfig := app.LoadConfig(env)
	err = app.InitExtensions(envConfig, env)
	if err != nil {
		log.Panicf("init extensions failed: %v\n", err)
	}

	client := app.EntClient
	if err := client.Debug().Schema.Create(
		context.Background(),
		schema.WithHooks(func(next schema.Creator) schema.Creator {
			return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
				return next.Create(ctx, tables...)
			})
		}),
		migrate.WithForeignKeys(false),
	); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("migrate successfully")
}

func init() {
	cmd := &cobra.Command{
		Use: "migrate",
		Run: runMigrate,
	}
	rootCmd.AddCommand(cmd)
}
