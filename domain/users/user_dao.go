package users

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

var collection *mongo.Collection

func ConnectDb() {

	// connect to MongoDB

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://kumar:Kvivzrohl538cjVR@cluster0.momx3.gcp.mongodb.net/user-db?retryWrites=true&w=majority"))

	// Ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	// Connect to the database
	db := client.Database("go_mongo")
	fmt.Println("db:", db)

	collection = client.Database("user-db").Collection("user-col")

}

func (user *User) FindByEmailAndPassword() error {

	ConnectDb()

	filter := bson.M{"email": user.Email}

	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	fmt.Println("err:", err)

	fmt.Println("user in dao:", user)

	if err != nil {
		return err
	}

	/*
		if err != nil {

			return restErrDb
		}*/

	return nil
}

func (user *User) UpdateToken() error {

	ConnectDb()

	filter := bson.M{"email": user.Email}

	update := bson.M{"$set": bson.M{"accessToken": user.AccessToken, "date_created": user.DateCreated}}

	_, err := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	if err != nil {
		return err
	}

	return nil
}
