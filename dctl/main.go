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
	ir.Delete("/ssc/:id", api.RmContainer)
	ir.Get("/stats/:id", api.GetDockerStats)
	ir.Put("/pause/:id", api.PauseContainer)
	ir.Put("/unpause/:id", api.UnpauseContainer)

	ir.Listen(":8080")
}
