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

// Update godoc
//
//	@Summary		Update source
//	@Description	Update source
//	@Tags			Source
//	@Accept			json
//	@Produce		json
//	@Param			source_id	path		string						true	"Source identifier"
//	@Param			request		body		sourceparam.UpdateRequest	true	"Update source request body"
//	@Success		201			{object}	sourceparam.UpdateResponse
//	@Failure		400			{object}	httputil.HTTPError
//	@Failure		401			{object}	httputil.HTTPError
//	@Failure		500			{object}	httputil.HTTPError
//	@Security		JWTToken
//	@Router			/sources/{source_id} [post]
func (h Handler) Update(ctx echo.Context) error {
	// get user id from context
	claim, ok := ctx.Get(h.authSvc.GetConfig().ContextKey).(*authservice.Claims)
	if !ok {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid auth token",
		})
	}

	var req sourceparam.UpdateRequest
	if err := ctx.Bind(&req); err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, errmsg.ErrBadRequest)
	}

	req.UserID = claim.UserID

	resp, err := h.sourceSvc.Update(req)
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
