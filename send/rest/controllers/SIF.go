package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"

	send "../.."
)

// PublishSIF :
func PublishSIF(c echo.Context) error {
	defer func() {
		uPHE(recover(), "./log.txt", false, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	body, e := ioutil.ReadAll(c.Request().Body)
	uPE(e)
	nV, nS := send.SIF(string(body))
	return c.JSON(http.StatusAccepted, fSpf("<%d> value tuples, <%d> struct tuples have been sent", nV, nS))
}
