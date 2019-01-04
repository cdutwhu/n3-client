package xjy

import (
	"fmt"
	"strings"

	u "github.com/cdutwhu/util"
)

var (
	// PE is
	PE = u.PanicOnError

	// PE1 is
	PE1 = u.PanicOnError1

	// PH is
	PH = u.PanicHandle

	// PC is
	PC = u.PanicOnCondition

	sI  = strings.Index
	sLI = strings.LastIndex
	sT  = strings.Trim
	sTL = strings.TrimLeft
	sTR = strings.TrimRight

	pln = fmt.Println
	pf  = fmt.Printf
	epf = fmt.Errorf
	spf = fmt.Sprintf
)
