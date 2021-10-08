package db

import "fmt"

const username = "aman"
const password = "qYHbfWVwpp3sXHx4"

var DB_URL = fmt.Sprintf("mongodb+srv://%v:%v@cluster0.6wjjz.mongodb.net/Instagram?retryWrites=true&w=majority", username, password)
