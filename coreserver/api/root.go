package api

import "github.com/kataras/iris"

// GetRoot handles GET /
func GetRoot(ctx *iris.Context) {
	ctx.WriteString("Fanach core server")
}
