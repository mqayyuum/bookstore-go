package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
	config := getDBConfig()

	d, err := gorm.Open("mysql", config.generateConnectionString())
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}

func getEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("error loading .env file")
	}

	return os.Getenv(key)
}

type DBConfig struct {
	Username string
	Password string
	Database string
}

func getDBConfig() DBConfig {
	username := getEnv("MYSQL_USERNAME")
	password := getEnv("MYSQL_PASSWORD")
	database := getEnv("MYSQL_DATABASE")

	return DBConfig{
		Username: username,
		Password: password,
		Database: database,
	}
}

func (c *DBConfig) generateConnectionString() string {
	config := getDBConfig()

	return fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=True&loc=Local", config.Username, config.Password, config.Database)
}
