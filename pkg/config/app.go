package config

import (
	"fmt"
	"log"

	"github.com/TwiN/go-color"
	"github.com/egacheru/services/pkg/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// godotenv package
	get := utils.GetEnvWithKey
	USER := get("DB_USER")
	PASS := get("DB_PASSWORD")
	HOST := get("DB_HOST")
	PORT := get("DB_PORT")
	DBNAME := get("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	d, err := gorm.Open("mysql", URL)
	log.Print(color.Ize(color.Yellow, "Connecting to the database..."))

	if err != nil {
		panic(err.Error())
	}
	log.Print(color.Ize(color.Green, "Successfully Connected to the database!"))
	db = d
}
func GetDB() *gorm.DB {
	return db
}
