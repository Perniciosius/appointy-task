package api

import (
	"appointy-task/model"
	"appointy-task/utils/router"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var postCollection = database.Collection("/posts")

// Create post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.Write([]byte("Unable to process request currently"))
		log.Println(err)
		return
	}
	if post.Caption == "" || post.ImageUrl == "" || post.UserId == "" {
		w.Write([]byte("Caption, User ID and Image URL cannot be empty"))
		return
	}
	if !checkUserExists(post.UserId) {
		w.Write([]byte("Invalid user id"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var result *mongo.InsertOneResult
	result, err = postCollection.InsertOne(ctx, post)
	if err != nil {
		log.Fatalln(err)
	}
	w.Write([]byte(fmt.Sprintf("Post id: %v", result.InsertedID)))
}

// Get post by id
func GetPost(w http.ResponseWriter, r *http.Request) {
	postId := router.GetParam(r, 0)
	id, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		w.Write([]byte("Invalid post id"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var post model.Post
	postCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&post)
	var result []byte
	result, err = json.Marshal(post)
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(result)
}

// Get list of posts by user id
func GetPostList(w http.ResponseWriter, r *http.Request) {
	userId := router.GetParam(r, 0)
	if !checkUserExists(userId) {
		w.Write([]byte("Invalid user id"))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := postCollection.Find(ctx, bson.M{"userId": userId})
	if err != nil {
		log.Fatalln(err)
	}
	var postList []model.Post
	for cursor.Next(context.TODO()) {
		var element model.Post
		err = cursor.Decode(&element)
		if err != nil {
			log.Fatalln(err)
		}
		postList = append(postList, element)
	}
	cursor.Close(context.TODO())
	var result []byte
	result, err = json.Marshal(postList)
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(result)
}

// Function to check if post user exists
func checkUserExists(userId string) bool {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return false
	}
	var user model.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		log.Fatalln(err)
	}
	return true
}
