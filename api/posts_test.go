package api_test

import (
	"appointy-task/api"
	"appointy-task/model"
	"appointy-task/utils/router"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostCreate(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Post(ts.URL+"/posts", "application/json", strings.NewReader("{\"userId\":\"6160b5e731b11bc9fb96d866\", \"imageUrl\":\"https://images.unsplash.com/photo-1592564630984-7410f94db184?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1446&q=80\"}"))
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	actual := string(body)
	expected := "ObjectID"
	if !strings.Contains(actual, expected) {
		t.Errorf("Exptected: %v, got %v", expected, actual)
	}
}

func TestPostCreateWithCaption(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Post(ts.URL+"/posts", "application/json", strings.NewReader("{\"userId\":\"6160b5e731b11bc9fb96d866\", \"caption\":\"Rick\", \"imageUrl\":\"https://images.unsplash.com/photo-1592564630984-7410f94db184?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1446&q=80\"}"))
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	actual := string(body)
	expected := "ObjectID"
	if !strings.Contains(actual, expected) {
		t.Errorf("Exptected: %v, got %v", expected, actual)
	}
}

func TestPostCreateWithoutUserId(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Post(ts.URL+"/posts", "application/json", strings.NewReader("{\"caption\":\"Rick\", \"imageUrl\":\"https://images.unsplash.com/photo-1592564630984-7410f94db184?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1446&q=80\"}"))
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	actual := string(body)
	expected := "User ID and Image URL cannot be empty"
	if !strings.Contains(actual, expected) {
		t.Errorf("Exptected: %v, got %v", expected, actual)
	}
}

func TestPostCreateWithoutImageUrl(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Post(ts.URL+"/posts", "application/json", strings.NewReader("{\"userId\":\"6160b5e731b11bc9fb96d866\", \"caption\":\"Rick\"}"))
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	actual := string(body)
	expected := "User ID and Image URL cannot be empty"
	if !strings.Contains(actual, expected) {
		t.Errorf("Exptected: %v, got %v", expected, actual)
	}
}

func TestPostCreateInvalidUserId(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Post(ts.URL+"/posts", "application/json", strings.NewReader("{\"userId\":\"60b5e731b11bc9fb96d8\", \"caption\":\"Rick\", \"imageUrl\":\"https://images.unsplash.com/photo-1592564630984-7410f94db184?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1446&q=80\"}"))
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	actual := string(body)
	expected := "Invalid user id"
	if !strings.Contains(actual, expected) {
		t.Errorf("Exptected: %v, got %v", expected, actual)
	}
}

func TestGetPost(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Get(ts.URL + "/posts/61613679f9e22c430aefd713")
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	actual := string(body)
	expected := "61613679f9e22c430aefd713"
	if !strings.Contains(actual, expected) {
		t.Errorf("Exptected: %v, got %v", expected, actual)
	}
}

func TestGetPostInvalidId(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Get(ts.URL + "/posts/13679f9e22c430aefd713")
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	actual := string(body)
	expected := "Invalid post id"
	if !strings.Contains(actual, expected) {
		t.Errorf("Exptected: %v, got %v", expected, actual)
	}
}

func TestGetPostList(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Get(ts.URL + "/posts/users/6160b5e731b11bc9fb96d866")
	if err != nil {
		t.Error(err)
	}

	var actual []model.Post

	err = json.NewDecoder(res.Body).Decode(&actual)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	if len(actual) == 0 {
		t.Error("Empty post list received")
	}
}

func TestGetPostListInvalidUserId(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Get(ts.URL + "/posts/users/b5e731b11bc9fb96d866")
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	actual := string(body)
	expected := "Invalid user id"
	if !strings.Contains(actual, expected) {
		t.Error("Empty post list received")
	}
}

func TestGetPostListPaginationLimit(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Get(ts.URL + "/posts/users/6160b5e731b11bc9fb96d866?limit=2")
	if err != nil {
		t.Error(err)
	}

	var actual []model.Post

	err = json.NewDecoder(res.Body).Decode(&actual)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	if len(actual) > 2 {
		t.Errorf("Expected less than 2 posts, received %v", len(actual))
	}
}

func TestGetPostListPaginationSkip(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Get(ts.URL + "/posts/users/6160b5e731b11bc9fb96d866?skip=10000")
	if err != nil {
		t.Error(err)
	}

	var actual []model.Post

	err = json.NewDecoder(res.Body).Decode(&actual)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	if len(actual) != 0 {
		t.Errorf("Expected 0 posts, received %v", len(actual))
	}
}
