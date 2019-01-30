package send

import (
	"fmt"
	"log"
	"strings"

	c "../config"
	u "github.com/cdutwhu/go-util"
)

var (
	uPE  = u.PanicOnError
	PE   = uPE
	uPE1 = u.PanicOnError1
	PE1  = uPE1
	uPC  = u.PanicOnCondition
	PC   = uPC
	uPH  = u.PanicHandle
	PH   = uPH
	uPHE = u.PanicHandleEx
	PHE  = uPHE
	uLE  = u.LogOnError
	LE   = uLE
	fPln = fmt.Println
	Pln  = fPln
	fPf  = fmt.Printf
	Pf   = fPf
	fEf  = fmt.Errorf
	Ef   = fEf
	fSpf = fmt.Sprintf
	Spf  = fSpf
	lPln = log.Println
	LPln = lPln

	sC = strings.Contains
	SC = sC

	e   error
	Cfg *c.Config
)

const (
	TERMMARK = "ENDENDEND"
	HEADTRIM = "sif."
	DELAY    = 200
)

type (
	sType int
)

const (
	SIF  sType = 0
	XAPI sType = 1
)
