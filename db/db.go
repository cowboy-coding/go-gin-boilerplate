package db

import (
	"fmt"
	"log"

	"github.com/cowboy-coding/go-gin-boilerplate/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Init() {
	c := config.GetConfig()
	config_string := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", c.GetString("db.host"), c.GetString("db.port"), c.GetString("db.user"), c.GetString("db.dbname"), c.GetString("db.password"))
	log.Print(config_string)
	db, err := gorm.Open("postgres", config_string)
	if err == nil {
	}

	defer db.Close()
}

func GetDB() *gorm.DB {
	return db
}
