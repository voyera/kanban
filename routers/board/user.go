package board

import (
	"net/http"

	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/middleware"
)

// ListMembers gets a list of member on board accessible by the authenticated user.
func ListMembers(ctx *middleware.Context) {
	members, err := ctx.DataSource.ListMembers(ctx.Query("project_id"))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, members)
}
