package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../../xjy"
	u "github.com/cdutwhu/util"
	"github.com/labstack/echo"
	"github.com/nsip/n3-messages/messages"
	"github.com/nsip/n3-messages/n3grpc"
)

var (
	pln = fmt.Println
	pf  = fmt.Printf

	verSIF1 = int64(1)
	verSIF2 = int64(1)
	verXAPI = int64(1)
)

var n3pub *n3grpc.Publisher

// PublishSIF :
func PublishSIF(c echo.Context) error {
	body, e := ioutil.ReadAll(c.Request().Body)
	if e == nil {
		content := u.Str(string(body))
		if content.L() > 0 && content.IsXMLSegSimple() {
			nameSpace := "Aa5fKf2UmyfCufY6JFmQpX12j1jjDFSUfbFUEE92t2nx"
			contextName := "abc-sif"
			to := "192.168.76.10" //"localhost"
			if n3pub == nil {
				n3pub, e = n3grpc.NewPublisher(to, 5777)
			}

			/*******************************************/
			done := make(chan int, 2)

			go xjy.YAMLAllValuesAsync(xjy.Xstr2Y(content.V()), "RefId", true, true, func(p, v, id string) {
				tuple, _ := messages.NewTuple(id, p, v)
				tuple.Version = verSIF1
				verSIF1++
				e = n3pub.Publish(tuple, nameSpace, contextName)
				pln("---", *tuple)
			}, done)

			/*****/

			go xjy.XMLStructAsync(content.V(), "RefId", true, func(p, v string) {
				tuple, _ := messages.NewTuple(p, "::", v)
				tuple.Version = verSIF2
				verSIF2++
				e = n3pub.Publish(tuple, nameSpace, contextName)
			}, done)

			pf("finish1: %d\n", <-done)
			pf("finish2: %d\n", <-done)
			log.Println("messages sent")

			return c.JSON(http.StatusAccepted, content.V())
		}
		return c.JSON(http.StatusBadRequest, "Request Body is invalid xml segment")
	}
	return c.JSON(http.StatusBadRequest, e)
}

// PublishXAPI :
func PublishXAPI(c echo.Context) error {
	body, e := ioutil.ReadAll(c.Request().Body)
	if e == nil {
		content := u.Str(string(body))
		return c.JSON(http.StatusAccepted, content.V())
	}
	return c.JSON(http.StatusBadRequest, e)
}
