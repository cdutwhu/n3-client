package xjy

import (
	"io/ioutil"
	"testing"
)

func TestYAMLAllValuesAsync(t *testing.T) {
	defer func() { PH(recover(), "./log.txt", true) }()

	//yamlstr, done := Xfile2Y("./files/nswdig.xml"), make(chan int)
	//ioutil.WriteFile(`./files/nswdig.yaml`, []byte(yamlstr), 0666)
	//yamlstr, done := Xfile2Y("./files/staffpersonal.xml"), make(chan int)
	//ioutil.WriteFile(`./files/staffpersonal.yaml`, []byte(yamlstr), 0666)
	yamlstr, done := Jfile2Y(`./files/xapifile.json`), make(chan int)
	ioutil.WriteFile(`./files/xapifile.yaml`, []byte(yamlstr), 0666)

	idx := 0
	go YAMLAllValuesAsync(yamlstr, "id", !true, true, func(path, value, id string) {
		idx++
		pf("%06d : %s\n", idx, path)
		pf("%s\n", value)
		pf("%s\n", id)
		pln("-----------------------------------------")
	}, done)
	pf("finish: %d\n", <-done)

	//fbytes, err := ioutil.ReadFile("./files/nswdig.yaml")
	//PE(err)
}

func TestYAMLTag(t *testing.T) {
	pln(YAMLTag(`- name: Andrew Downes`))
	pln(YAMLTag(`actor:`))
	pln(YAMLTag(`  mbox: mailto:teampb@example.com`))
	pln(YAMLTag(`      homePage: http://www.example.com`))
	pln(YAMLTag(`  - mbox_sha1sum: ebd31e95054c018b10727ccffd2ef2ec3a016ee9`))
	pln(YAMLTag(`version: 1.0.0`))
	pln(YAMLTag(`      - "9"`))
	pln(YAMLTag(`- a`))
	pln(YAMLTag(`-RefId: D3E34F41-9D75-101A-8C3D-00AA001A1652`))
}

func TestYAMLValue(t *testing.T) {
	pln(YAMLValue(`- name: Andrew Downes`))
	pln(YAMLValue(`actor:`))
	pln(YAMLValue(`  mbox: mailto:teampb@example.com`))
	pln(YAMLValue(`      homePage: http://www.example.com`))
	pln(YAMLValue(`  - mbox_sha1sum: ebd31e95054c018b10727ccffd2ef2ec3a016ee9`))
	pln(YAMLValue(`version: 1.0.0`))
	pln(YAMLValue(`      - "9"`))
	pln(YAMLValue(`- a`))
	pln(YAMLValue(`-RefId: D3E34F41-9D75-101A-8C3D-00AA001A1652`))
}
