package n3pub

import (
	"fmt"

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

	pln = fmt.Println
	pf  = fmt.Printf
)
