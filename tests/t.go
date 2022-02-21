package main

//
//import "C"
//import (
//	"bookstore_case/pkg"
//	"context"
//	"fmt"
//	"log"
//)
//
//func main() {
//	c := pkg.Connection()
//	collection := c.Database("test").Collection("trainers")
//
//	type Trainer struct {
//		Name string
//		Age  int
//		City string
//	}
//	ash := Trainer{"Ash", 10, "Pallet Town"}
//	insertResult, err := collection.InsertOne(context.TODO(), ash)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(insertResult)
//}
