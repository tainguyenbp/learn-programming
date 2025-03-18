package db

import (
    "context"
    "fmt"
    "time"
    "os"
    "errors"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() error {
    uri := os.Getenv("MONGODB_URI")
    if uri == "" {
        return errors.New("MONGODB_URI environment variable is not set")
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    var err error
    Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
    if err != nil {
        return err
    }
    err = Client.Ping(context.TODO(), nil)
    if err != nil {
        return err
    }
    fmt.Println("Connected to MongoDB!")
    return nil
}

func GetDatabase() *mongo.Database {
    return Client.Database("invoice")
}

func Close() {
    if Client != nil {
        Client.Disconnect(context.TODO())
    }
}
