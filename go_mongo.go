package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	fmt.Println("Hello MongoDB Test")
	log.Println("Starting MongoDB connection test")

	connectionString := "mongodb://enosAdmin:securePassword@localhost:27017/enoh?authSource=admin"
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database("enoh").Collection("users")
	log.Println("Collection instance created:", collection.Name(), collection.Database().Name(), collection.Database().Client())
	p, err := collection.Find(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Found documents:", p)
}
