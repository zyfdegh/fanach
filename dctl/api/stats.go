package api

import (
	"github.com/kataras/iris"
	"github.com/zyfdegh/fanach/dctl/entity"
	"github.com/zyfdegh/fanach/dctl/service"
)

// GetDockerStats get usage infomation of a container
func GetDockerStats(ctx *iris.Context) {
	resp := &entity.RespGetStats{}

	id := ctx.Param("id")
	if len(id) == 0 {
		resp.ErrNo = iris.StatusBadRequest
		resp.Errmsg = "invalid param id"
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}

	stats, err := service.DockerStats(id)
	if err != nil {
		resp.ErrNo = iris.StatusInternalServerError
		resp.Errmsg = err.Error()
		ctx.JSON(iris.StatusInternalServerError, resp)
		return
	}

	if stats != nil {
		resp.Stats = *stats
		resp.Success = true
	}
	ctx.JSON(iris.StatusOK, resp)
	return
}
