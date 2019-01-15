package filewatcher

import (
	"fmt"
	"log"
	"strings"

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

	sC = strings.Contains
)
