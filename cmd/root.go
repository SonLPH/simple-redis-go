package cmd

import (
	"fmt"
	"log"
	"os"
	"simple-redis-go/config"
	"simple-redis-go/db"
	"simple-redis-go/internal"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "simple-redis-go",
	Long:  "simple-redis-go",
	Run: func(_ *cobra.Command, _ []string) {
		c := config.GetConfig()
		_db, err := db.DBConnection(c)
		if err != nil {
			log.Fatal(err)
		}
		_rdb, err := db.RedisConnection(c)
		if err != nil {
			log.Fatal(err)
		}

		app := fiber.New()
		group := app.Group("/api/v1")
		internal.Router(group, _db, _rdb)
		addr := fmt.Sprintf(":%s", c.ServerPort)
		app.Listen(addr)
	},
}

func Execute() {
	rootCmd.AddCommand(databaseCmd)
	rootCmd.AddCommand(reidsCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
