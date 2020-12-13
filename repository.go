package main

import (
	"context"
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Init() error
	FindById(id string) (interface{}, error)
	InsertOne(id shortUrl) error
}

type repository struct {
	collection *mongo.Collection
}

func (repo *repository) FindById(id string) (interface{}, error) {
	fmt.Println("Fetch by " + id)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	// _, err := repo.collection.InsertOne(ctx, thing)

	res := struct{ Fieldint int }{}

	objectId, err := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objectId}
	// filter := bson.D{{"fieldstr", id}}

	err = repo.collection.FindOne(ctx, filter).Decode(&res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo repository) InsertOne(su shortUrl) error {
	fmt.Println("Insert one")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	// thing := struct{ Fieldint int }{Fieldint: 12}

	_, err := repo.collection.InsertOne(ctx, su)

	if err != nil {
		return err
	}

	return nil
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

	return nil
}
