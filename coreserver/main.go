package main

import (
	"time"

	"github.com/kataras/go-sessions/sessiondb/leveldb"
	"github.com/kataras/iris"
	"github.com/zyfdegh/fanach/coreserver/api"
	"github.com/zyfdegh/fanach/coreserver/db"
)

func main() {
	ir := iris.New()
	ir.Get("/", api.GetRoot)

	ir.Post("/users", api.CreateUser)
	ir.Get("/users", api.GetUsers)
	ir.Get("/users/:id", api.GetUser)
	ir.Put("/users/:id", api.ModifyUser)
	ir.Delete("/users/:id", api.DeleteUser)

	ir.Post("/session", api.PostSession)
	ir.Delete("/session/:id", api.DeleteSession)

	// save sessions to LevelDB(with GC)
	db := leveldb.New(leveldb.Config{
		Path:         db.SessDBFile,
		CleanTimeout: 1 * time.Hour,
		MaxAge:       7 * 24 * time.Hour,
	})
	ir.UseSessionDB(db)

	ir.Listen(":8080")
}
