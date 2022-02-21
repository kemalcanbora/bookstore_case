package models

type Config struct {
	JWTSecret                string `json:"jwt_secret"`
	MongoUserName            string `json:"mongo_user_name"`
	MongoPassword            string `json:"mongo_password"`
	MongoDatabase            string `json:"mongo_database"`
	MongoURL                 string `json:"mongo_url"`
	MongoUserCollection      string `json:"mongo_user_collection"`
	MongoBookCollection      string `json:"mongo_book_collection"`
	MongoBookOrderCollection string `json:"mongo_book_order_collection"`
}
