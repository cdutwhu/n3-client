package main

import (
	"fmt"
	"strings"

	u "github.com/cdutwhu/go-util"
)

var (
	uPE  = u.PanicOnError
	uPE1 = u.PanicOnError1
	uPH  = u.PanicHandle
	uPC  = u.PanicOnCondition

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

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEpf = fmt.Errorf
	fSpf = fmt.Sprintf
)
