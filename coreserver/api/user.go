package api

import (
	"github.com/kataras/iris"
	"github.com/zyfdegh/fanach/coreserver/entity"
	"github.com/zyfdegh/fanach/coreserver/service"
)

// CreateUser handles POST /users
func CreateUser(ctx *iris.Context) {
	resp := entity.RespPostUser{}

	user := entity.User{}
	err := ctx.ReadJSON(user)
	if err != nil {
		resp.Errmsg = err.Error()
		resp.ErrNo = iris.StatusBadRequest
		ctx.JSON(resp.ErrNo, resp)
		return
	}

	user, err = service.CreateUser(user)
	if err != nil {
		resp.Errmsg = err.Error()
		resp.ErrNo = iris.StatusInternalServerError
		ctx.JSON(resp.ErrNo, resp)
		return
	}
	resp.Success = true
	resp.ID = user.ID
	ctx.JSON(iris.StatusOK, resp)
	return
}

// GetUser handles GET /users/:id
func GetUser(ctx *iris.Context) {}

// ModifyUser handles PUT /users/:id
func ModifyUser(ctx *iris.Context) {}

// DeleteUser handles DELETE /users/:id
func DeleteUser(ctx *iris.Context) {}
