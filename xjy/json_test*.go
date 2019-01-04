package xjy

import (
	"io/ioutil"
	"testing"
)

func TestJSONScanObjects(t *testing.T) {
	defer func() { PH(recover(), "./log.txt", true) }()
	jsonbytes, err := ioutil.ReadFile("./files/xapifile.json")
	PE(err)

	// ids, objs, pos := JSONScanObjects(string(jsonbytes), "id")
	// for i, id := range ids {
	// 	pln(id)
	// 	pln(pos[i])
	// 	pln(objs[i])
	// }

	objstr := JSONObjStrByID(string(jsonbytes), "id", "6690e6c9-3ef0-4ed3-8b37-7f3964730bef")
	// pln(objstr)

	elestr, _ := JSONEleStrByTag(objstr, "testArray")
	pln(elestr)

	//elestr1 := JSONEleStrByTag(elestr, "definition")

	// children, list := JSONFindChildren(elestr)
	// pln(children)
	// pln(list)

}
