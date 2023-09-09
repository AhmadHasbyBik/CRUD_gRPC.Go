package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/grpc_go"))
	if err != nil {
		log.Fatalf("Database load failed %v", err.Error())
	}

	return db
}
