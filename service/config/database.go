package config

import (
 "gorm.io/driver/postgres"
 "gorm.io/gorm"
 "log"
)

var DB *gorm.DB

func ConnectDatabase() {
 dsn := "host=localhost user=usr password=pwd dbname=db port=5432 sslmode=disable"
 database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
 if err != nil {
  log.Fatal("Failed to connect to the database:", err)
 }

 DB = database
}