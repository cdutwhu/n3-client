package xjy

import (
	"io/ioutil"
	"testing"
)

func TestXMLScanObjects(t *testing.T) {

	defer func() { PH(recover(), "./log.txt", true) }()

	xmlbytes, err := ioutil.ReadFile("./files/staffpersonal.xml")
	PE(err)

	done := make(chan int)
	go XMLStructAsync(string(xmlbytes), "RefId", true, func(p, v string) {
		pf("%-60s:: %s\n", p, v)
	}, done)

	pf("finish: %d\n", <-done)

	ids, objtags, psarr := XMLScanObjects(string(xmlbytes), "RefId")
	pln(len(objtags))
	for _, objtag := range objtags {
		pln(objtag)
	}
	for i := range ids {
		pf("%s -- %s -- %d\n", objtags[i], ids[i], psarr[i])
	}

	//fmt.Print(string(xmlbytes[psarr[1]:psarr[2]]))

	// xmlobj := XMLObjStrByID(string(xmlbytes), "RefId", "D3E34F41-9D75-101A-8C3D-00AA001A1652")
	// pln(xmlobj)
	// pln()
	// pln(XMLFindAttributes(xmlobj))
}
