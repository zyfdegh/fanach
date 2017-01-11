package api

import (
	"github.com/kataras/iris"
	"github.com/zyfdegh/fanach/dctl/entity"
	"github.com/zyfdegh/fanach/dctl/service"
)

// PauseContainer pauses a docker container
func PauseContainer(ctx *iris.Context) {
	resp := &entity.Resp{}

	id := ctx.Param("id")
	if len(id) == 0 {
		resp.ErrNo = iris.StatusBadRequest
		resp.Errmsg = "invalid param id"
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}

	err := service.DockerPause(id)
	if err != nil {
		resp.ErrNo = iris.StatusInternalServerError
		resp.Errmsg = err.Error()
		ctx.JSON(iris.StatusInternalServerError, resp)
		return
	}
	resp.Success = true
	ctx.JSON(iris.StatusOK, resp)
	return
}
