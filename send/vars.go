package send

import (
	"fmt"
	"log"

	u "github.com/cdutwhu/go-util"
)

var (
	uPE  = u.PanicOnError
	uPE1 = u.PanicOnError1
	uPC  = u.PanicOnCondition
	uPH  = u.PanicHandle
	uPHE = u.PanicHandleEx
	uLE  = u.LogOnError

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSpf = fmt.Sprintf
	lPln = log.Println

	verSIF  = int64(1)
	verXAPI = int64(1)

	e error
)
