package main

import (
	"net/http"
	"testing"
	"time"

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
			"user": {
				"id": "34b7da764b21d298ef307d04d8152dc5",
				"username": "tom",
				"password": "***",
				"wechat_id": "tomwechat",
				"type": "",
				"email": "tom@email.com"
			}
		}
	`
	e.POST("/users").
		WithJSON(regJSONReg).
		Expect().
		Status(http.StatusOK).
		JSON().Schema(respJSONReg)

	// query user "tom"
	respJSONQuery := `
		{
			"success": true,
			"errno": 0,
			"errmsg": "",
			"user": {
				"id": "34b7da764b21d298ef307d04d8152dc5",
				"username": "tom",
				"password": "***",
				"wechat_id": "tomwechat",
				"type": "",
				"email": "tom@email.com"
			}
		}
	`
	e.GET("/users/34b7da764b21d298ef307d04d8152dc5").
		Expect().
		Status(http.StatusOK).
		JSON().Schema(respJSONQuery)

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
				"id": "34b7da764b21d298ef307d04d8152dc5",
				"username": "tom",
				"password": "***",
				"wechat_id": "tom123",
				"type": "",
				"email": "tom@outlook.com"
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

	// register user "bob"
	regJSONReg2 := map[string]string{
		"username":  "bob",
		"password":  "password",
		"wechat_id": "bob123",
		"email":     "bob@email.com",
	}
	respJSONReg2 := `
			{
				"success": true,
				"errno": 0,
				"errmsg": "",
				"id": "9f9d51bc70ef21ca5c14f307980a29d8"
			}
		`
	e.POST("/users").
		WithJSON(regJSONReg2).
		Expect().
		Status(http.StatusOK).
		JSON().Schema(respJSONReg2)

	// query all users
	respJSONQuery2 := `
			{
				"success": true,
				"errno": 0,
				"errmsg": "",
				"users": [
					{
						"id": "34b7da764b21d298ef307d04d8152dc5",
						"username": "tom",
						"password": "***",
						"wechat_id": "tom123",
						"type": "",
						"email": "tom@outlook.com"
					},
					{
						"id": "9f9d51bc70ef21ca5c14f307980a29d8",
						"username": "bob",
						"password": "***",
						"wechat_id": "bob123",
						"type": "",
						"email": "bob@email.com"
					}
				]
			}
		`
	e.GET("/users").
		Expect().
		Status(http.StatusOK).
		JSON().Schema(respJSONQuery2)

	// create session(login)
	now := time.Now()

	regJSONSess := map[string]string{
		"username": "bob",
		"password": "password",
	}
	cookie := e.POST("/sess").
		WithJSON(regJSONSess).
		Expect().
		Status(http.StatusOK).
		Cookie("sessid")

	// check cookie
	cookie.Expires().InRange(now, now.Add(24*time.Hour))
	cookie.Path().Equal("/")

	// delete session (logout)
	e.DELETE("/sess/sessid").
		Expect().
		Status(http.StatusOK)

	// clean up
	// test delete user "tom"
	e.DELETE("/users/34b7da764b21d298ef307d04d8152dc5").
		Expect().
		Status(http.StatusOK)

	// test delete user "bob"
	e.DELETE("/users/9f9d51bc70ef21ca5c14f307980a29d8").
		Expect().
		Status(http.StatusOK)
}
