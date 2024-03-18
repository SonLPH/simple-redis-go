package cmd

import (
	"log"
	"os"
	"simple-redis-go/config"
	"simple-redis-go/db"

	"github.com/spf13/cobra"
)

var databaseCmd = &cobra.Command{
	Use:   "check_connect_db",
	Short: "check db",
	Long:  "check connection to db",
	Run: func(cmd *cobra.Command, args []string) {
		c := config.GetConfig()
		_, err := db.DBConnection(c)
		if err != nil {
			log.Fatalln("Error connecting to db", err)
			os.Exit(1)
		}
		log.Println("connected db")
		os.Exit(0)
	},
}

var reidsCmd = &cobra.Command{
	Use:   "check_connect_redis",
	Short: "check redis",
	Long:  "check connection to redis",
	Run: func(cmd *cobra.Command, args []string) {
		c := config.GetConfig()
		_, err := db.RedisConnection(c)
		if err != nil {
			log.Fatalln("Error connecting to redis", err)
			os.Exit(1)
		}
		log.Println("connected redis")
		os.Exit(0)
	},
}
