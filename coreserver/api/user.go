package api

import (
	"github.com/kataras/iris"
	"github.com/zyfdegh/fanach/coreserver/entity"
	"github.com/zyfdegh/fanach/coreserver/service"
)

// CreateUser handles POST /users
func CreateUser(ctx *iris.Context) {
	resp := entity.RespPostUser{}

	user := &entity.User{}
	err := ctx.ReadJSON(user)
	if err != nil {
		resp.Errmsg = err.Error()
		resp.ErrNo = iris.StatusBadRequest
		ctx.JSON(resp.ErrNo, resp)
		return
	}

	user, err = service.CreateUser(*user)
	if err != nil {
		resp.Errmsg = err.Error()
		resp.ErrNo = iris.StatusInternalServerError
		if err == service.ErrUserExist {
			resp.ErrNo = iris.StatusConflict
		}
		ctx.JSON(resp.ErrNo, resp)
		return
	}
	resp.Success = true
	resp.ID = user.ID
	ctx.JSON(iris.StatusOK, resp)
	return
}

// GetUser handles GET /users/:id
func GetUser(ctx *iris.Context) {
	resp := entity.RespGetUser{}

	userID := ctx.Param("id")
	user, err := service.GetUser(userID)
	if err != nil {
		resp.Errmsg = err.Error()
		if err == service.ErrUserNotFound {
			resp.ErrNo = iris.StatusNotFound
		}
		ctx.JSON(resp.ErrNo, resp)
		return
	}

	resp.Success = true
	resp.User = *user
	resp.User.Password = entity.HidenString
	ctx.JSON(iris.StatusOK, resp)
	return
}

// GetUsers handles GET /users
// Return users in bulk
func GetUsers(ctx *iris.Context) {
	resp := entity.RespGetUsers{}

	users, err := service.GetUsers()
	if err != nil {
		resp.Errmsg = err.Error()
		resp.ErrNo = iris.StatusInternalServerError
		ctx.JSON(resp.ErrNo, resp)
		return
	}

	resp.Success = true
	if users != nil {
		for _, user := range *users {
			user.Password = entity.HidenString
			resp.Users = append(resp.Users, user)
		}
	}
	ctx.JSON(iris.StatusOK, resp)
	return
}

// ModifyUser handles PUT /users/:id
func ModifyUser(ctx *iris.Context) {}

// DeleteUser handles DELETE /users/:id
func DeleteUser(ctx *iris.Context) {}
