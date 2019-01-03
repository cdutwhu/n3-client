package xjy

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestJSONScanObjects(t *testing.T) {
	defer func() { PH(recover(), "./log.txt", true) }()
	jsonbytes, err := ioutil.ReadFile("./files/xapifile.json")
	PE(err)

	// ids, objs, pos := JSONScanObjects(string(jsonbytes), "id")
	// for i, id := range ids {
	// 	fmt.Println(id)
	// 	fmt.Println(pos[i])
	// 	fmt.Println(objs[i])
	// }

	objstr := JSONObjStrByID(string(jsonbytes), "id", "6690e6c9-3ef0-4ed3-8b37-7f3964730bef")
	// fmt.Println(objstr)

	fmt.Println(JSONEleStrByTag(objstr, "test5"))
}
