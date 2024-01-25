package projecthandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/param"
	"github.com/ormushq/ormus/pkg/echomsg"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/httpmsg"
)

func (h Handler) Create(ctx echo.Context) error {
	var req param.CreateProjectRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echomsg.DefaultMessage(errmsg.ErrBadRequest))
	}

	vErr := h.projectValidator.ValidateCreateRequest(req)
	if vErr != nil {
		msg, code := httpmsg.Error(vErr.Err)

		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  vErr.Fields,
		})
	}

	newProject, err := h.projectSvc.Create(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echomsg.DefaultMessage(errmsg.ErrBadRequest))
	}

	resp := param.CreateProjectResponse{
		Name: newProject.Name,
		ID:   newProject.ID,
	}

	return ctx.JSON(http.StatusCreated, resp)
}
