package database

import (
	"log"
	"os"
	"time"

	"authentication/data"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var counts int

func NewPostgresDatabase() *gorm.DB {
	return ConnectToDB()
}

func ConnectToDB() *gorm.DB {
	//dsn := viper.GetString("postgresql.dsn")
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Printf("Postgres not yet ready... %s", err)
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func openDB(dsn string) (*gorm.DB, error) {
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := connection.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	connection.AutoMigrate(&data.User{})

	return connection, nil
}
