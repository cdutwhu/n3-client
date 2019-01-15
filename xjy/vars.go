package xjy

import (
	"fmt"
	"strings"

	u "github.com/cdutwhu/go-util"
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
	sHP = strings.HasPrefix
	sHS = strings.HasSuffix
	sFF = strings.FieldsFunc
	sC  = strings.Contains
	sJ  = strings.Join

	pln = fmt.Println
	pf  = fmt.Printf
	epf = fmt.Errorf
	spf = fmt.Sprintf
)

type DataType int

const (
	XML  DataType = 0
	JSON DataType = 1
	SIF  DataType = 0
	XAPI DataType = 1
)
