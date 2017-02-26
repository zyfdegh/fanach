package main

import (
	"net/http"
	"testing"

	"gopkg.in/kataras/iris.v6/httptest"
)

func TestCoreServer(t *testing.T) {
	server := newCoreServer()
	e := httptest.New(server, t)

	// test if server started
	e.GET("/").
		Expect().
		Status(http.StatusOK).
		Body().Equal("Fanach core server")

	// register user "tom"
	regJSONReg := map[string]string{
		"username":  "tom",
		"password":  "secret",
		"wechat_id": "tomwechat",
		"email":     "tom@email.com",
	}
	respJSONReg := `
		{
			"success": true,
			"errno": 0,
			"errmsg": "",
			"id": "34b7da764b21d298ef307d04d8152dc5"
		}
	`
	e.POST("/users").
		WithJSON(regJSONReg).
		Expect().
		Status(http.StatusOK).
		JSON().Schema(respJSONReg)

	// update user "tom"
	regJSONUpdate := map[string]string{
		"password":  "strongpassword",
		"wechat_id": "tom123",
		"email":     "tom@outlook.com",
	}
	respJSONUpdate := `
		{
			"success": true,
			"errno": 0,
			"errmsg": "",
			"user": {
				"username": "tom",
				"password": "***",
				"wechat_id": "tom123",
				"email": "tom@outlook.com",
				"id": "34b7da764b21d298ef307d04d8152dc5"
			}
		}
	`
	e.PUT("/users/34b7da764b21d298ef307d04d8152dc5").
		WithJSON(regJSONUpdate).
		Expect().
		Status(http.StatusOK).
		JSON().Schema(respJSONUpdate)

	// register again with name "tom"
	regJSONConflict := map[string]string{
		"username": "tom",
		"password": "secret",
	}
	respJSONConflict := `
		{
			"success": false,
			"errno": 409,
			"errmsg": "duplicated username",
			"id": ""
		}
	`
	e.POST("/users").
		WithJSON(regJSONConflict).
		Expect().
		Status(http.StatusConflict).
		JSON().Schema(respJSONConflict)

	// test delete user "tom"
	e.DELETE("/users/34b7da764b21d298ef307d04d8152dc5").
		Expect().
		Status(http.StatusOK)
}
