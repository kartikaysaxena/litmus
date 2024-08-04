package main

import (
	"context"
	"fmt"
	// "fmt"
	"log"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	deleteDB()
	// indexVersion()

}

func deleteDB() {
	// // MongoDB connection URL
	uri := "mongodb://m1:27015,m2:27016,m3:27017/?replicaSet=rs0"

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Check the connection
	// err = client.Ping(context.Background(), nil)
	// if err != nil {
	//     log.Fatalf("Error pinging MongoDB: %v", err)
	// }
	// fmt.Println("Connected to MongoDB!")

	// List databases
	// databases, err := client.ListDatabaseNames(context.Background(), nil)
	// if err != nil {
	//     log.Fatalf("Error listing databases: %v", err)
	// }

	// Drop each database
	client.Database("auth").Drop(context.Background())
	client.Database("litmus").Drop(context.Background())

	// for _, db := range databases {
	//     if db == "admin" || db == "local" || db == "config" {
	//         // Skip system databases
	//         continue
	//     }
	//     err := client.Database("auth").Drop(context.Background())
	//     if err != nil {
	//         log.Printf("Error dropping database %s: %v", db, err)
	//     } else {
	//         fmt.Printf("Dropped database: %s\n", db)
	//     }
	// }

	// Disconnect from MongoDB
	err = client.Disconnect(context.Background())
	if err != nil {
	    log.Fatalf("Error disconnecting from MongoDB: %v", err)
	}
	fmt.Println("Disconnected from MongoDB.")
}

func indexVersion() {

	uri := "mongodb://m1:27015,m2:27016,m3:27017/?replicaSet=rs0"

    // Set client options
    clientOptions := options.Client().ApplyURI(uri)

    // Connect to MongoDB
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatalf("Error connecting to MongoDB: %v", err)
    }

    // Check the connection
    // err = client.Ping(context.Background(), nil)
    // if err != nil {
    //     log.Fatalf("Error pinging MongoDB: %v", err)
    // }
    // fmt.Println("Connected to MongoDB!")

    // // List databases
    // // databases, err := client.ListDatabaseNames(context.Background(), nil)
    // // if err != nil {
    // //     log.Fatalf("Error listing databases: %v", err)
    // // }

    // // Drop each database
	// client.Database("auth").Drop(context.Background())
	// client.Database("litmus").Drop(context.Background())

    // // for _, db := range databases {
    // //     if db == "admin" || db == "local" || db == "config" {
    // //         // Skip system databases
    // //         continue
    // //     }
    // //     err := client.Database("auth").Drop(context.Background())
    // //     if err != nil {
    // //         log.Printf("Error dropping database %s: %v", db, err)
    // //     } else {
    // //         fmt.Printf("Dropped database: %s\n", db)
    // //     }
    // // }

    // // Disconnect from MongoDB
    // err = client.Disconnect(context.Background())
    // if err != nil {
    //     log.Fatalf("Error disconnecting from MongoDB: %v", err)
    // }
    // fmt.Println("Disconnected from MongoDB.")

	collection := client.Database("litmus").Collection("serverConfig")

	// Prepare the document to insert
	doc := bson.D{
		{"key", "version"},
		{"value", "3.8.0"},
	}

	// Insert the document
	_, err = collection.InsertOne(context.Background(), doc)
	if err != nil {
		return
	}

}
