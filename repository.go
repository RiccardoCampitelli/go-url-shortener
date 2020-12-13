package main

import (
	"context"
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Init() error
	FindById(id string) (interface{}, error)
}

type repository struct {
	collection *mongo.Collection
}

func (repo *repository) FindById(id string) (interface{}, error) {
	fmt.Println("Fetch by " + id)

	thing := MongoFields{
		FieldStr:  "Some Value",
		FieldInt:  12345,
		FieldBool: true,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	_, err := repo.collection.InsertOne(ctx, thing)

	if err != nil {
		return nil, err
	}

	return thing, nil
}

type MongoFields struct {
	FieldStr  string `json:"Field Str"`
	FieldInt  int    `json:"Field Int"`
	FieldBool bool   `json:"Field Bool"`
}

func (repo *repository) Init() error {

	envs, err := godotenv.Read(".env")

	if err != nil {
		return err
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(envs["MONGO_URL"]))

	if err != nil {
		return err
	}

	client.Connect(context.TODO())

	collection := client.Database("root").Collection("test2")

	repo.collection = collection
	// /mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass%20Community&ssl=false

	return nil
}
