package query

import (
	"fmt"
	"log"

	c "../config"
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

	e   error
	Cfg *c.Config
)

type (
	qType int
)

const (
	SIF  qType = 0
	XAPI qType = 1
)
