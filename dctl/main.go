package main

import (
	"github.com/kataras/iris"
	"github.com/zyfdegh/fanach/dctl/api"
)

func main() {
	ir := iris.New()
	ir.Get("/", iris.ToHandlerFunc(api.GetRoot))
	ir.Get("/dver", iris.ToHandlerFunc(api.GetDockerVersion))
	ir.Post("/ssc", iris.ToHandlerFunc(api.AddSsContainer))
	ir.Delete("/ssc", iris.ToHandlerFunc(api.RmContainer))
	ir.Get("/stats", iris.ToHandlerFunc(api.GetDockerStats))
	ir.Put("/pause", iris.ToHandlerFunc(api.PauseContainer))
	ir.Put("/unpause", iris.ToHandlerFunc(api.UnpauseContainer))

	ir.Listen(":8080")
}
