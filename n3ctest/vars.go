package n3ctest

import (
	"fmt"

	u "github.com/cdutwhu/util"
)

var (
	PE  = u.PanicOnError
	PE1 = u.PanicOnError1
	PH  = u.PanicHandle
	PC  = u.PanicOnCondition
	pln = fmt.Println
	pf  = fmt.Printf
)
