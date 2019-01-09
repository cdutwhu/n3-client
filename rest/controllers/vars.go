package controllers

import (
	"fmt"

	"github.com/nsip/n3-messages/n3grpc"
)

var (
	pln = fmt.Println
	pf  = fmt.Printf

	verSIF1 = int64(1)
	verSIF2 = int64(1)
	verXAPI = int64(1)

	n3pub       *n3grpc.Publisher
	nameSpace   = "Aa5fKf2UmyfCufY6JFmQpX12j1jjDFSUfbFUEE92t2nx"
	ctxNameSIF  = "abc-sif"
	ctxNameXAPI = "abc-xapi"
	sendTo      = "192.168.76.10" //"localhost"
)
