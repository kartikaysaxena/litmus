package main

import (
    "context"
    "fmt"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    // MongoDB server URL
    dbServerURL := "mongodb://m1:27015,m2:27016,m3:27017/?replicaSet=rs0"

    // Database names to be deleted
    dbNames := []string{"litmus", "auth"}

    // Set client options
    clientOptions := options.Client().ApplyURI(dbServerURL)

    // Connect to the MongoDB server
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")

    // Drop the databases
    for _, dbName := range dbNames {
        err = client.Database(dbName).Drop(context.TODO())
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Database '%s' has been deleted.\n", dbName)
    }

    // Disconnect from the MongoDB server
    err = client.Disconnect(context.TODO())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Disconnected from MongoDB.")
}
