package sourcehandler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/logger"
	"github.com/ormushq/ormus/manager/managerparam/sourceparam"
	"github.com/ormushq/ormus/manager/service/authservice"
	"github.com/ormushq/ormus/manager/validator"
	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/httpmsg"
	"github.com/ormushq/ormus/pkg/httputil"
)

// Delete godoc
//
//	@Summary		Delete source
//	@Description	Delete source
//	@Tags			Source
//	@Accept			json
//	@Produce		json
//	@Param			source_id	path		string	true	"Source identifier"
//	@Success		200			{object}	sourceparam.DeleteResponse
//	@Failure		400			{object}	httputil.HTTPError
//	@Failure		401			{object}	httputil.HTTPError
//	@Failure		500			{object}	httputil.HTTPError
//	@Security		JWTToken
//	@Router			/sources/{source_id} [delete]
func (h Handler) Delete(ctx echo.Context) error {
	claim, ok := ctx.Get(h.authSvc.GetConfig().ContextKey).(*authservice.Claims)
	if !ok {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid auth token",
		})
	}

	var req sourceparam.DeleteRequest
	if err := ctx.Bind(&req); err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, errmsg.ErrBadRequest)
	}

	req.UserID = claim.UserID

	resp, err := h.sourceSvc.Delete(req)

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
