package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"

	s "../.."
)

// PublishXAPI :
func PublishXAPI(c echo.Context) error {
	defer func() {
		s.PHE(recover(), s.Cfg.Global.ErrLog, false, func(msg string, others ...interface{}) {
			c.JSON(http.StatusBadRequest, msg)
		})
	}()

	body, e := ioutil.ReadAll(c.Request().Body)
	s.PE(e)
	n := s.XAPI(string(body))
	return c.JSON(http.StatusAccepted, s.Spf("%d tuples has been sent", n))
}
