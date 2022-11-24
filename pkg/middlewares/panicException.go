package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"user-api/pkg/errors"
)

func PanicExceptionHandling() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					switch v := err.(type) {
					case *errors.Error:
						v.Log()
						c.JSON(v.StatusCode, v.Public)
					default:
						logrus.Error(v)
						c.JSON(http.StatusInternalServerError, struct {
							Error interface{} `json:"error"`
						}{
							Error: err,
						})
					}
				}
			}()
			return next(c)
		}
	}
}
