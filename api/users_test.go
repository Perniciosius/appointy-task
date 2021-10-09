package api_test

import (
	"appointy-task/api"
	"appointy-task/utils/router"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader("{\"name\":\"aman\", \"email\":\"amaaa@aman.com\", \"password\":\"abc123\"}"))
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(body), "ObjectID") && !strings.Contains(string(body), "User already exists with given email") {
		t.Error("Failed")
	}
}

func TestCreateUserEmptyName(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader("{\"email\":\"amaaa@aman.com\", \"password\":\"abc123\"}"))
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(body), "Name must not be empty") {
		t.Error("Failed")
	}
}

func TestCreateUserInvalidEmail(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader("{\"name\":\"someone\", \"email\":\"someone\", \"password\":\"abc123\"}"))
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(body), "Provide a valid email") {
		t.Error("Failed")
	}
}

func TestCreateUserEmptyPassword(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader("{\"name\":\"someone\", \"email\":\"someone@s.com\"}"))
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(body), "Password must not be empty") {
		t.Error("Failed")
	}
}

func TestGetUserInvalidUserId(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users/sdlfjlsdfj342j4l234")
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(body), "Invalid user id") {
		t.Error("Failed")
	}
}

func TestGetUser(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: api.Routes})
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users/6160b5e731b11bc9fb96d866")
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(body), "{\"id\":\"6160b5e731b11bc9fb96d866\",\"name\":\"aman\",\"email\":\"aman@aman.com\",\"password\":\"gSZ0GrznooMOgqEkOeZQcpIcqSgrkZF9ZDp-FU2OW-H9asN_SpFSsuU1ETUpq4G1T5kG9FmU8HKeKZU-5FaThCIQHi3dIM3BAEaV12_uoSs=\"}") {
		t.Error("Failed")
	}
}
