package rest

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"./controllers"
)

// HostHTTPForPubAsync : Host a HTTP Server for publishing inbound SIF(xml) or XAPI(json) string(request body) to <n3-transport> grpc Server
func HostHTTPForPubAsync() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	// Route
	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "n3client is alive\n") })
	e.POST("/sif", controllers.PublishSIF)
	e.POST("/xapi", controllers.PublishXAPI)

	// Server
	e.Start(":1323")
}
