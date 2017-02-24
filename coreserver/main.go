package main

import (
	"github.com/zyfdegh/fanach/coreserver/api"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	server := newCoreServer()
	server.Listen(":8080")
}

func newCoreServer() *iris.Framework {
	ir := iris.New()
	ir.Adapt(httprouter.New())

	ir.Get("/", api.GetRoot)

	ir.Post("/users", api.CreateUser)
	ir.Get("/users", api.GetUsers)
	ir.Get("/users/:id", api.GetUser)
	ir.Put("/users/:id", api.ModifyUser)
	ir.Delete("/users/:id", api.DeleteUser)

	ir.Post("/session", api.PostSession)
	ir.Delete("/session/:id", api.DeleteSession)

	// save sessions to LevelDB(with GC)
	// db := leveldb.New(leveldb.Config{
	// 	Path:         db.SessDBFile,
	// 	CleanTimeout: 1 * time.Hour,
	// 	MaxAge:       7 * 24 * time.Hour,
	// })
	// ir.UseSessionDB(db)

	return ir
}
