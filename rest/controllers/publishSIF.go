package controllers

import (
	"io/ioutil"
	"log"
	"net/http"

	"../../xjy"
	u "github.com/cdutwhu/util"
	"github.com/labstack/echo"
	"github.com/nsip/n3-messages/messages"
	"github.com/nsip/n3-messages/n3grpc"
)

// PublishSIF :
func PublishSIF(c echo.Context) error {
	body, e := ioutil.ReadAll(c.Request().Body)
	if e == nil {
		content := u.Str(string(body))
		if content.L() > 0 && content.IsXMLSegSimple() {
			if n3pub == nil {
				n3pub, e = n3grpc.NewPublisher(sendTo, 5777)
			}

			done := make(chan int, 2)

			go xjy.YAMLAllValuesAsync(xjy.Xstr2Y(content.V()), "RefId", true, true, func(p, v, id string) {
				tuple, _ := messages.NewTuple(id, p, v)
				tuple.Version = verSIF1
				verSIF1++
				e = n3pub.Publish(tuple, nameSpace, ctxNameSIF)
				pln("---", *tuple)
			}, done)

			go xjy.XMLStructAsync(content.V(), "RefId", true, func(p, v string) {
				tuple, _ := messages.NewTuple(p, "::", v)
				tuple.Version = verSIF2
				verSIF2++
				e = n3pub.Publish(tuple, nameSpace, ctxNameSIF)
			}, done)

			pf("sif sent 1: %d\n", <-done)
			pf("sif sent 2: %d\n", <-done)
			log.Println("messages sent")

			return c.JSON(http.StatusAccepted, content.V())
		}
		return c.JSON(http.StatusBadRequest, "Request Body is invalid xml segment")
	}
	return c.JSON(http.StatusBadRequest, e)
}
