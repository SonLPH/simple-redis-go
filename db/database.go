package db

import (
	"fmt"
	"log"
	"simple-redis-go/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection(conf config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresHost,
		conf.PostgresPort,
		conf.PostgresDatabase)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection to database failed", err)
	}

	return db, nil
}
