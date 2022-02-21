package config

import (
	"bookstore_case/models"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Configure() models.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var c models.Config
	c.JWTSecret = os.Getenv("SECRET_KEY")
	c.MongoUserName = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	c.MongoPassword = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	c.MongoDatabase = os.Getenv("MONGO_INITDB_DATABASE")
	c.MongoURL = os.Getenv("MONGO_URL")
	c.MongoUserCollection = "users"
	c.MongoBookCollection = "books"
	c.MongoBookOrderCollection = "orders"
	return c
}
