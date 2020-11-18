package database

import (
	"github.com/luisgomez29/api_lol/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DBURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	return db
}
