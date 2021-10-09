package router_test

import (
	"appointy-task/utils/router"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var routes = []router.Route{
	router.NewRoute("POST", "/users", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Success"))
	}),
	router.NewRoute("GET", "/users/([^/]+)", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(router.GetParam(r, 0)))
	}),
}

func TestRouteHandler1(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: routes})
	defer ts.Close()
	res, err := http.Post(ts.URL+"/users", "application/json", nil)
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	actual := string(body)
	expected := "Success"
	if !strings.Contains(actual, expected) {
		t.Errorf("Expected response: %v, got response: %v", expected, actual)
	}
}

func TestRouteHandler2(t *testing.T) {
	ts := httptest.NewServer(router.RouteHandler{Routes: routes})
	defer ts.Close()
	res, err := http.Get(ts.URL + "/users/1234")
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	actual := string(body)
	expected := "1234"
	if !strings.Contains(actual, expected) {
		t.Errorf("Expected response: %v, got response: %v", expected, actual)
	}
}
