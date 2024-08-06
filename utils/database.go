package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection

func ConnectDB() error {
	clientOptions := options.Client().ApplyURI("mongodb+srv://itachi:Anoman2002@cluster0.aiemcbo.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	UserCollection = client.Database("MiniFoot").Collection("Utilisateur")
	return nil
}
