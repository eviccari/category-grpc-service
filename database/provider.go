package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/eviccari/category-grpc-service/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Provider - Describe Provider structure
type Provider struct {}

// NewProvider - Create new Provider instance
func NewProvider() (provider *Provider){
	return &Provider{}
}

// MongoDB - Get database instance for MongoDB
func (p *Provider) MongoDB() (db *mongo.Collection, err error) {
	var ctx = context.TODO()

	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
		utils.PanicIfError(err)
	}

	credential := options.Credential{
		Username: os.Getenv("MONGODB_USER"),
		Password: os.Getenv("MONGODB_PASS"),
	}

	uri := fmt.Sprintf("mongodb://%s:%s", os.Getenv("MONGODB_HOST"), os.Getenv("MONGODB_PORT"))
	clientOptions := options.Client().ApplyURI(uri).SetAuth(credential)
	client, openError := mongo.Connect(ctx, clientOptions)
	utils.PanicIfError(openError)

	collection := client.Database(os.Getenv("MONGODB_REPO")).Collection(os.Getenv("MONGODB_COLL"))
	return collection, nil
}