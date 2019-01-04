package xjy

import (
	"testing"
)

func TestYAMLAllValuesAsync(t *testing.T) {
	defer func() { PH(recover(), "./log.txt", true) }()

	//yamlbytes, done := Xfile2Y("./files/staffpersonal.xml"), make(chan int)
	yamlbytes, done := Jfile2Y(`./files/xapifile.json`), make(chan int)

	go YAMLAllValuesAsync(string(yamlbytes), "id", false, true, func(path, value, id string) {
		pf("%s -- %s -- %s\n", path, value, id)
	}, done)
	pf("finish: %d\n", <-done)

	//fbytes, err := ioutil.ReadFile("./files/nswdig.yaml")
	//PE(err)
}

func TestYAMLTag(t *testing.T) {
	pln(YAMLTag(`- name: Andrew Downes`))
}

func TestYAMLValue(t *testing.T) {
	pln(YAMLValue(`- a`))
}
