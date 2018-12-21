package xjy

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestXMLScanObjects(t *testing.T) {

	defer func() { PH(recover(), "./log.txt", true) }()

	xmlbytes, err := ioutil.ReadFile("./files/staffpersonal.xml")
	PE(err)

	done := make(chan int)
	go XMLStructAsync(string(xmlbytes), "RefId", true, func(p, v string) {
		fmt.Printf("%-60s:: %s\n", p, v)
	}, done)

	fmt.Printf("finish: %d\n", <-done)

	//ids, objs, psarr := XMLScanObjects(string(xmlbytes), "RefId")
	// fmt.Println(len(objs))
	// for _, obj := range objs {
	// 	fmt.Println(obj)
	// }
	// for i := range ids {
	// 	fmt.Printf("%s -- %s -- %d\n", objs[0], ids[0], psarr[0])
	// }

	//fmt.Print(string(xmlbytes[psarr[1]:psarr[2]]))

	//for _, id := range ids {
	//D3F5B90C-D85D-4728-8C6F-0D606070606C
	//D0E7421A-38AE-48D0-985F-D5525D32B56D
	//xmlobj := XMLObjStrByID(string(xmlbytes), "D3E34F41-9D75-101A-8C3D-00AA001A1652", ids, objs, psarr)
	//fmt.Println(xmlobj)
	//fmt.Println(obj)
	//fmt.Println()
	//xmlele := XMLEleStrByTag(xmlobj, "StudentList")
	//fmt.Println(XMLFindAttributes(xmlobj))
	//}

}
