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

// List godoc
//
//	@Summary		List sources
//	@Description	List sources
//	@Tags			Source
//	@Accept			json
//	@Produce		json
//	@Param			last_token_id	query		string	false	"Last token fetched"
//	@Param			per_page		query		int		false	"Per page count"
//	@Success		200				{object}	sourceparam.ListResponse
//	@Failure		400				{object}	httputil.HTTPError
//	@Failure		401				{object}	httputil.HTTPError
//	@Failure		500				{object}	httputil.HTTPError
//	@Security		JWTToken
//	@Router			/sources [get]
func (h Handler) List(ctx echo.Context) error {
	// get user id from context
	claim, ok := ctx.Get(h.authSvc.GetConfig().ContextKey).(*authservice.Claims)
	if !ok {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid auth token",
		})
	}

	var req sourceparam.ListRequest
	if err := ctx.Bind(&req); err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, errmsg.ErrBadRequest)
	}

	req.UserID = claim.UserID

	if req.PerPage == 0 {
		req.PerPage = 10
	}
	var lastTokenID int64 = -9223372036854775808
	if req.LastTokenID == 0 {
		req.LastTokenID = lastTokenID
	}

	// call save method in service
	resp, err := h.sourceSvc.List(req)

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
