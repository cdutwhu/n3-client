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

// PublishXAPI :
func PublishXAPI(c echo.Context) error {
	body, e := ioutil.ReadAll(c.Request().Body)
	if e == nil {
		content := u.Str(string(body))
		if content.L() > 0 && content.IsJSON() {
			if n3pub == nil {
				n3pub, e = n3grpc.NewPublisher(sendTo, 5777)
			}

			done := make(chan int)
			go xjy.YAMLAllValuesAsync(xjy.Jstr2Y(content.V()), "id", false, true, func(p, v, id string) {
				tuple, _ := messages.NewTuple(id, p, v)
				tuple.Version = verXAPI
				verXAPI++
				e = n3pub.Publish(tuple, nameSpace, ctxNameXAPI)
				pln("---", *tuple)
			}, done)

			pf("xapi sent : %d\n", <-done)
			log.Println("messages sent")

			return c.JSON(http.StatusAccepted, content.V())
		}
		return c.JSON(http.StatusBadRequest, "Request Body is invalid json string")
	}
	return c.JSON(http.StatusBadRequest, e)
}
