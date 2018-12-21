package xjy

import (
	"fmt"
	"testing"
)

func TestUpperLevelLine(t *testing.T) {
	defer func() { PH(recover(), "./log.txt", true) }()

	yamlbytes, done := Xfile2Y("./files/staffpersonal.xml"), make(chan int)
	go YAMLAllValuesAsync(string(yamlbytes), "RefId", true, func(path, value, id string) {
		fmt.Printf("%s -- %s -- %s\n", path, value, id)
	}, done)
	fmt.Printf("finish: %d\n", <-done)

	//fbytes, err := ioutil.ReadFile("./files/nswdig.yaml")
	//PE(err)
}
