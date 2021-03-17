package models

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//DB Connection instance
var DB *gorm.DB

//DbConnect Is used to connect to postgresql databases. This is backed by Gorm module. Single DB connection can be used to serve multiple requests.
func DbConnect(conf *Config) error {

	dbName := conf.Harepd.Repmgr.Db
	dsn := fmt.Sprintf("dbname=%s sslmode=disable", dbName)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		Log.Error(err)
		return err
	}
	//database.AutoMigrate(&Master{})

	DB = database
	return nil
}
