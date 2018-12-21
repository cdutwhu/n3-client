package main

// func SendXmlToDataStore(filename string) {
// 	fi, err := os.Lstat(filename)
// }

import (
	"fmt"

	u "github.com/cdutwhu/util"
)

var PE = u.PanicOnError
var PE1 = u.PanicOnError1
var PC = u.PanicOnCondition
var PH = u.PanicHandle
var PHE = u.PanicHandleEx
var LE = u.LogOnError

// func SendXmlToDataStore(filename string) {
// 	defer PH("", false)
// 	// defer PHE("", false, func(emsg string, params ...interface{}) {
// 	// 	fmt.Println(emsg)
// 	// 	fmt.Println(params[0])
// 	// }, "do more things?")

// 	fi, err := os.Lstat(filename)
// 	PE(err)
// 	PC(fi.Mode().IsDir(), fmt.Errorf("%s is a directory", filename))
// 	PC(!strings.HasSuffix(filename, ".xml"), fmt.Errorf("%s is not an XML file", filename))

// 	file, err := os.Open(filename)
// 	PE1(err, fmt.Sprintf("Cannot read in file %s\n", filename))
// }

type S struct {
	s1 string
	s2 string
	s3 string
}

// func printArrAsync(ss []<-chan string) {
// 	for _, s := range ss {
// 		fmt.Println(<-s)
// 	}
// 	fmt.Println("over11")
// }

// func printAsync(s <-chan string) {
// 	fmt.Println(<-s)
// 	fmt.Println("over1")
// }

// func printSAsync(s <-chan S) {
// 	//fmt.Println(<-s)
// 	for i := range s {
// 		fmt.Println(i)
// 	}
// 	fmt.Println("over1")
// }

func main() {
	// SendXmlToDataStore("C:\\")

	// s := make(chan S)
	// go printSAsync(s)

	// time.Sleep(2 * time.Second)

	// s <- S{"a", "b", "c"}
	// close(s)

	// //s <- "abc"
	// time.Sleep(10 * time.Second)
	// fmt.Println("over")

	chans := make(chan int, 5)
	chans <- 1
	chans <- 2
	chans <- 3
	chans <- 4
	chans <- 5
	//chans <- 6

	fmt.Println(<-chans)
	fmt.Println(<-chans)
	fmt.Println(<-chans)
	fmt.Println(<-chans)
	fmt.Println(<-chans)

	chans <- 6

	// close(chans)

	// for c := range chans {
	// 	fmt.Println(c)

	// }

}
