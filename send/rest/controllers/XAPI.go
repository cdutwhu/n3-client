package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"

	send "../.."
)

// PublishXAPI :
func PublishXAPI(c echo.Context) error {
	defer func() {
		uPHE(recover(), "./log.txt", false, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	body, e := ioutil.ReadAll(c.Request().Body)
	uPE(e)
	n := send.XAPI(string(body))
	return c.JSON(http.StatusAccepted, fSpf("%d tuples has been sent", n))
}
