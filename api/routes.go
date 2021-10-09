package api

import (
	"appointy-task/utils/router"
)

var Routes = []router.Route{
	router.NewRoute("POST", "/users", CreateUser),
	router.NewRoute("GET", "/users/([^/]+)", GetUser),
	router.NewRoute("POST", "/posts", CreatePost),
	router.NewRoute("GET", "/posts/([^/]+)", GetPost),
	router.NewRoute("GET", "/posts/users/([^/]+)", GetPostList),
}
