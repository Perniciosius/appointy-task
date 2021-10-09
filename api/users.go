package api

import (
	"appointy-task/db"
	"appointy-task/model"
	utils "appointy-task/utils/hashing"
	"appointy-task/utils/router"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var database = db.GetDatabase()
var userCollection = database.Collection("/users")

// Create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user, tempUser model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Write([]byte("Unable to process request currently"))
		log.Println(err)
		return
	}
	if user.Name == "" {
		w.Write([]byte("Name must not be empty"))
		return
	}
	if user.Password == "" {
		w.Write([]byte("Password must not be empty"))
	}
	if _, err := mail.ParseAddress(user.Email); err != nil {
		w.Write([]byte("Provide a valid email"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&tempUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			var result *mongo.InsertOneResult
			user.Password = utils.HashPassword(user.Password, nil)
			result, err = userCollection.InsertOne(ctx, user)
			if err != nil {
				log.Fatalln(err)
			}
			w.Write([]byte(fmt.Sprintf("User id: %v", result.InsertedID)))
			return
		}
		log.Fatalln(err)
	}
	w.Write([]byte("User already exists with given email"))
}

// Get user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := router.GetParam(r, 0)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		w.Write([]byte("Invalid user id"))
		return
	}
	var user model.User
	userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	var result []byte
	result, err = json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(result)
}
