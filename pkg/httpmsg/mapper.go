package httpmsg

import (
	"errors"
	"net/http"

	"github.com/ormushq/ormus/pkg/errmsg"
	"github.com/ormushq/ormus/pkg/richerror"
)

func Error(err error) (message string, code int) {
	serverErrCode := 500

	var richError richerror.RichError
	switch {
	case errors.As(err, &richError):
		var re richerror.RichError
		errors.As(err, &re)

		msg := re.Message()
		code := mapKindToHTTPStatusCode(re.Kind())

		if code >= serverErrCode {
			msg = errmsg.ErrSomeThingWentWrong
		}

		return msg, code
	default:
		return err.Error(), http.StatusBadRequest
	}
}

func mapKindToHTTPStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindInvalid:
		return http.StatusUnprocessableEntity
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindForbidden:
		return http.StatusForbidden
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
