package server

import (
	"github.com/kataras/iris"
	"github.com/zyfdegh/fanach/deployer/server/api"
)

// Start launches http server
func Start() {
	ir := iris.New()

	ir.Get("/", iris.ToHandlerFunc(api.GetRoot))
	ir.Post("/deploy", iris.ToHandlerFunc(api.PostDeploy))
	ir.Post("/test", iris.ToHandlerFunc(api.PostTest))

	ir.Listen(":8080")
}
