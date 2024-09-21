package projecthandler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/managerparam/projectparam"
	"github.com/ormushq/ormus/manager/service/authservice"
	"github.com/ormushq/ormus/manager/validator"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/httpmsg"
	"github.com/ormushq/ormus/pkg/httputil"
)

// Delete godoc
//
//	@Summary		Delete project
//	@Description	Delete project
//	@Tags			Project
//	@Accept			json
//	@Produce		json
//	@Param			project_id	path		string	true	"Project identifier"
//	@Success		200			{object}	projectparam.DeleteResponse
//	@Failure		400			{object}	httputil.HTTPError
//	@Failure		401			{object}	httputil.HTTPError
//	@Failure		500			{object}	httputil.HTTPError
//	@Security		JWTToken
//	@Router			/projects/{project_id} [delete]
func (h Handler) Delete(ctx echo.Context) error {
	claim, ok := ctx.Get(h.authSvc.GetConfig().ContextKey).(*authservice.Claims)
	if !ok {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid auth token",
		})
	}

	var req projectparam.DeleteRequest
	if err := ctx.Bind(&req); err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, errmsg.ErrBadRequest)
	}

	req.UserID = claim.UserID

	resp, err := h.projectSvc.Delete(req)

	logger.LogError(err)
	var vErr *validator.Error
	if errors.As(err, &vErr) {
		msg, code := httpmsg.Error(vErr.Err)

		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  vErr.Fields,
		})
	}

	if err != nil {
		return httputil.NewErrorWithError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, resp)
}
