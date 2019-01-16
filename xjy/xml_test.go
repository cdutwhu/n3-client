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
		fPf("%-60s:: %s\n", p, v)
	}, done)

	fPf("finish: %d\n", <-done)

	ids, objtags, psarr := XMLScanObjects(string(xmlbytes), "RefId")
	fPln(len(objtags))
	for _, objtag := range objtags {
		fPln(objtag)
	}
	for i := range ids {
		fPf("%s -- %s -- %d\n", objtags[i], ids[i], psarr[i])
	}

	//fmt.Print(string(xmlbytes[psarr[1]:psarr[2]]))

	// xmlobj := XMLObjStrByID(string(xmlbytes), "RefId", "D3E34F41-9D75-101A-8C3D-00AA001A1652")
	// fPln(xmlobj)
	// fPln()
	// fPln(XMLFindAttributes(xmlobj))
}

func TestXMLEleStrByTag(t *testing.T) {
	fPln(XMLEleStrByTag(`		<OtherNames>
	<Name Type="AKA">
		<FamilyName>Anderson</FamilyName>
		<GivenName>Samuel</GivenName>
		<FullName>Samuel Anderson</FullName>
	</Name>
	<Name Type="PRF">
		<FamilyName>Rowinski</FamilyName>
		<GivenName>Sam</GivenName>
		<FullName>Sam Rowinski </FullName>
	</Name>
</OtherNames>`, "Name", 1))
}
