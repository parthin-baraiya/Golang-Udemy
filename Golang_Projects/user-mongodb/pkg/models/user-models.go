package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

type User struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Name   string             `json:"name" bson:"name"`
	Gender string             `json:"gender" bson:"gender"`
	Age    int                `json:"age" bson:"age"`
}

func getSession() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://parthin-baraiya:parthin123@cluster0.r3qqdpk.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client
}

func init() {
	db = getSession().Database("user-mongodb")
	fmt.Println(db)
}

func GetUserByID(ID string) User {
	userID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		panic(err)
	}

	var user User
	db.Collection("users").FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)
	return user
}

func GetUsers() []User {
	var users []User
	cursor, err := db.Collection("users").Find(context.TODO(), bson.M{})

	if err != nil {
		log.Println(err)
	}

	if err = cursor.All(context.TODO(), &users); err != nil {
		log.Println(err)
	}
	return users
}

func (user *User) CreateUser() {
	user.ID = primitive.NewObjectID()
	db.Collection("users").InsertOne(context.TODO(), user)
}

func DeleteUser(ID string) User {
	var user User
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		panic(err)
	}
	db.Collection("users").FindOneAndDelete(context.TODO(), bson.M{"_id": id}).Decode(&user)
	return user
}

func (user *User) UpdateUser(ID string) {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		panic(err)
	}
	// db.Collection("users").FindOneAndUpdate(context.TODO(), bson.M{"_id": id}, bson.M{"$set": user}).Decode(&newUser)
	// println(newUser.Name)
	db.Collection("users").UpdateByID(context.TODO(), id, bson.M{"$set": user})
}
