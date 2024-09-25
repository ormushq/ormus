package otela

import (
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/otel/attribute"
)

func EchoRequestLoggerLogValuesFunc(packageName, functionName string) func(c echo.Context, v middleware.RequestLoggerValues) error {
	return func(c echo.Context, v middleware.RequestLoggerValues) error {
		errMsg := ""
		if v.Error != nil {
			errMsg = v.Error.Error()
		}

		attributes := make([]attribute.KeyValue, 0)

		t := time.Time{}
		if v.StartTime != t {
			attributes = append(attributes, attribute.String("StartTime", v.StartTime.String()))
		}
		if v.Latency != 0 {
			attributes = append(attributes, attribute.Float64("Latency", v.Latency.Seconds()))
		}

		if v.Protocol != "" {
			attributes = append(attributes, attribute.String("Protocol", v.Protocol))
		}

		if v.RemoteIP != "" {
			attributes = append(attributes, attribute.String("RemoteIP", v.RemoteIP))
		}

		if v.Host != "" {
			attributes = append(attributes, attribute.String("Host", v.Host))
		}

		if v.Method != "" {
			attributes = append(attributes, attribute.String("Method", v.Method))
		}

		if v.URI != "" {
			attributes = append(attributes, attribute.String("URI", v.URI))
		}

		if v.URIPath != "" {
			attributes = append(attributes, attribute.String("URIPath", v.URIPath))
		}

		if v.RoutePath != "" {
			attributes = append(attributes, attribute.String("RoutePath", v.RoutePath))
		}

		if v.RequestID != "" {
			attributes = append(attributes, attribute.String("RequestID", v.RequestID))
		}

		if v.Referer != "" {
			attributes = append(attributes, attribute.String("Referer", v.Referer))
		}

		if v.UserAgent != "" {
			attributes = append(attributes, attribute.String("UserAgent", v.UserAgent))
		}

		if v.Status != 0 {
			attributes = append(attributes, attribute.Int("Status", v.Status))
		}

		if errMsg != "" {
			attributes = append(attributes, attribute.String("ErrMsg", errMsg))
		}

		if v.ContentLength != "" {
			attributes = append(attributes, attribute.String("ContentLength", v.ContentLength))
		}

		if v.ResponseSize != 0 {
			attributes = append(attributes, attribute.Int64("ResponseSize", v.ResponseSize))
		}

		if v.Headers != nil {
			for key, value := range v.Headers {
				attributes = append(attributes, attribute.String("Headers."+key, strings.Join(value, ",")))
			}
		}

		if v.QueryParams != nil {
			for key, value := range v.QueryParams {
				attributes = append(attributes, attribute.String("QueryParams."+key, strings.Join(value, ",")))
			}
		}

		if v.FormValues != nil {
			for key, value := range v.FormValues {
				attributes = append(attributes, attribute.String("QueryParams."+key, strings.Join(value, ",")))
			}
		}

		ctx, _ := TraceBuilder(packageName, functionName,
			WithContext(c.Request().Context()),
			WithSpanOptionAttributes(attributes...),
		)
		c.SetRequest(c.Request().WithContext(ctx))

		return nil
	}
}
