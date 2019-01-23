package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"

	s "../.."
)

// PublishSIF :
func PublishSIF(c echo.Context) error {
	defer func() {
		s.PHE(recover(), s.Cfg.Global.ErrLog, false, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	body, e := ioutil.ReadAll(c.Request().Body)
	s.PE(e)
	nV, nS, nA := s.SIF(string(body))
	return c.JSON(http.StatusAccepted, s.Spf("<%d> value tuples, <%d> struct tuples, <%d> array info tuples have been sent", nV, nS, nA))
}
