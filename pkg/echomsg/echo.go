package echomsg

import "github.com/labstack/echo/v4"

func DefaultMessage(message string) echo.Map {
	return echo.Map{"message": message}
}
