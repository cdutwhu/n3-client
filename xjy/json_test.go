package xjy

import (
	"io/ioutil"
	"testing"

	c "../config"
)

func TestJSONScanObjects(t *testing.T) {
	cfg := c.GetConfig("../config/config.toml")
	defer func() { uPH(recover(), cfg.Global.ErrLog, true) }()
	jsonbytes, err := ioutil.ReadFile("./files/xapifile.json")
	uPE(err)

	// ids, objs, pos := JSONScanObjects(string(jsonbytes), "id")
	// for i, id := range ids {
	// 	fPln(id)
	// 	fPln(pos[i])
	// 	fPln(objs[i])
	// }

	objstr := JSONObjStrByID(string(jsonbytes), "id", "6690e6c9-3ef0-4ed3-8b37-7f3964730bee")
	fPln(objstr)

	//elestr := JSONEleStrByTag(objstr, "actor")
	//fPln(elestr)

	//elestr1 := JSONEleStrByTag(elestr, "definition")

	// children, list := JSONFindChildren(elestr)
	// fPln(children)
	// fPln(list)

}
