package send

import (
	"fmt"
	"log"

	u "github.com/cdutwhu/util"
	"github.com/nsip/n3-messages/n3grpc"
)

var (
	PE  = u.PanicOnError
	PE1 = u.PanicOnError1
	PC  = u.PanicOnCondition
	PH  = u.PanicHandle
	PHE = u.PanicHandleEx
	LE  = u.LogOnError

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSpf = fmt.Sprintf
	lPln = log.Println	

	verSIF1 = int64(1)
	verSIF2 = int64(1)
	verXAPI = int64(1)

	n3pub       *n3grpc.Publisher
	e           error
	nameSpace   = "Aa5fKf2UmyfCufY6JFmQpX12j1jjDFSUfbFUEE92t2nx"
	ctxNameSIF  = "abc-sif"
	ctxNameXAPI = "abc-xapi"
	sendTo      = "192.168.76.10" //"localhost"
	sendToPort  = 5777
)
