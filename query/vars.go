package query

import (
	"fmt"
	"log"

	u "github.com/cdutwhu/util"
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

	e error
)
