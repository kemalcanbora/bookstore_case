package pkg

import (
	c "bookstore_case/config"
	"bookstore_case/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
	"time"
)

type Mongo interface {
	Insert(data interface{}, collectionName string) (*mongo.InsertOneResult, error)
	FindUserWithEmail(email string) (models.User, error)
	UpdateBook(book models.Book) error
	DeleteBook(id string) error
	FindBook(id string) (models.Book, error)
	Order(order models.Order) error
}

type MongoClient struct {
	Client  *mongo.Client
	Context context.Context
	Cancel  func()
}

func Connection() *MongoClient {
	var err error
	var mongoClient MongoClient

	credential := options.Credential{
		Username: c.Configure().MongoUserName,
		Password: c.Configure().MongoPassword,
	}
	mongoClient.Client, err = mongo.NewClient(options.Client().ApplyURI(c.Configure().MongoURL).SetAuth(credential))
	mongoClient.Context, mongoClient.Cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer mongoClient.Cancel()

	err_ := mongoClient.Client.Connect(mongoClient.Context)
	if err_ != nil {
		log.Fatal(err)
	}
	return &mongoClient
}

func (m *MongoClient) Insert(data interface{}, collectionName string) (*mongo.InsertOneResult, error) {
	collection := m.Client.Database(c.Configure().MongoDatabase).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func (m *MongoClient) FindUserWithEmail(email string) (models.User, error) {
	var dbUser models.User
	collection := m.Client.Database(c.Configure().MongoDatabase).Collection(c.Configure().MongoUserCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&dbUser)
	if err != nil {
		log.Println(err)
	}
	return dbUser, err
}

func (m *MongoClient) UpdateBook(book models.Book) error {
	collection := m.Client.Database(c.Configure().MongoDatabase).Collection(c.Configure().MongoBookCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.UpdateOne(ctx, bson.M{"_id": book.ID}, bson.M{"$set": book})
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *MongoClient) DeleteBook(id string) error {
	collection := m.Client.Database(c.Configure().MongoDatabase).Collection(c.Configure().MongoBookCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		fmt.Println("DeleteOne() document not found:", res)
	} else {
		// Print the results of the DeleteOne() method
		fmt.Println("DeleteOne Result:", res)
		// *mongo.DeleteResult object returned by API call
		fmt.Println("DeleteOne TYPE:", reflect.TypeOf(res))
	}
	return err
}

func (m *MongoClient) FindBook(id string) (models.Book, error) {
	var dbBook models.Book
	collection := m.Client.Database(c.Configure().MongoDatabase).Collection(c.Configure().MongoBookCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&dbBook)
	if err != nil {
		log.Println(err)
	}
	return dbBook, err
}

func (m *MongoClient) Order(order models.Order) error {
	var dbBook models.Book
	bookCollection := m.Client.Database(c.Configure().MongoDatabase).Collection(c.Configure().MongoBookCollection)
	orderCollection := m.Client.Database(c.Configure().MongoDatabase).Collection(c.Configure().MongoBookOrderCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := bookCollection.FindOne(ctx, bson.M{"_id": order.ProductID}).Decode(&dbBook)
	if dbBook.ID == "" {
		log.Println("Book not found with id:", order.ProductID)
	}

	quantity := dbBook.Quantity - order.Quantity
	if quantity < 0 {
		return fmt.Errorf("not enough quantity")
	}
	dbBook.Quantity = quantity
	_, err = bookCollection.UpdateOne(ctx, bson.M{"_id": dbBook.ID}, bson.M{"$set": dbBook})
	if err != nil {
		return err
	}
	_, err = orderCollection.InsertOne(ctx, order)
	if err != nil {
		return err
	}
	if err != nil {
		log.Println(err)
	}
	return err
}
