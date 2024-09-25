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

// Create godoc
//
//	@Summary		Create project
//	@Description	Create project
//	@Tags			Project
//	@Accept			json
//	@Produce		json
//	@Param			request	body		projectparam.CreateRequest	true	"Create project request body"
//	@Success		201		{object}	projectparam.CreateResponse
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		401		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Security		JWTToken
//	@Router			/projects [post]
func (h Handler) Create(ctx echo.Context) error {
	// get user id from context
	claim, ok := ctx.Get(h.authSvc.GetConfig().ContextKey).(*authservice.Claims)
	if !ok {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid auth token",
		})
	}

	var req projectparam.CreateRequest
	if err := ctx.Bind(&req); err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, errmsg.ErrBadRequest)
	}

	req.UserID = claim.UserID

	resp, err := h.projectSvc.Create(req)

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

	return ctx.JSON(http.StatusCreated, resp)
}
