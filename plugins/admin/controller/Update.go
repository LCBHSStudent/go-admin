package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"net/http"
)

func Update(ctx *context.Context) {

	param := guard.GetUpdateParam(ctx)

	err := param.Panel.UpdateDataFromDatabase(param.Value)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": "fail",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg": "ok",
	})
}
