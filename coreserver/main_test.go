package main

import (
	"net/http"
	"testing"

	"gopkg.in/kataras/iris.v6/httptest"
)

func TestRoot(t *testing.T) {
	server := newCoreServer()
	e := httptest.New(server, t)
	// test if server started
	e.GET("/").
		Expect().
		Status(http.StatusOK).
		Body().Equal("Fanach core server")
}

func TestRegister(t *testing.T) {
	server := newCoreServer()
	e := httptest.New(server, t)

	// test user register
	regJSON := map[string]string{
		"username": "tom",
		"password": "secret",
	}
	respJSON := `
		{
			"success": true,
			"errno": 0,
			"errmsg": "",
			"id": "9f9d51bc70ef21ca5c14f307980a29d8"
		}
	`
	e.POST("/users").
		WithJSON(regJSON).
		Expect().
		Status(http.StatusOK).
		JSON().Schema(respJSON)

	// register again
	respJSONConflict := `
		{
			"success": false,
			"errno": 409,
			"errmsg": "duplicated username",
			"id": ""
		}
	`
	e.POST("/users").
		WithJSON(regJSON).
		Expect().
		Status(http.StatusConflict).
		JSON().Schema(respJSONConflict)
}
