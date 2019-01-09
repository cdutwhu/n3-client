package n3ctest

import (
	"log"
	"testing"

	"github.com/nsip/n3-messages/messages"
	"github.com/nsip/n3-messages/n3grpc"

	"../xjy"
)

func TestN3Pub(t *testing.T) {
	defer func() { PH(recover(), "./log.txt", true) }()

	to := "192.168.76.10" //"localhost"
	n3pub, err := n3grpc.NewPublisher(to, 5777)
	PE(err)
	defer n3pub.Close()

	/*******************************************/
	nameSpace := "Aa5fKf2UmyfCufY6JFmQpX12j1jjDFSUfbFUEE92t2nx"
	contextName := "abc-xapi"

	done := make(chan int)
	jsonfile := "../xjy/files/xapifile.json"
	yamlStr := xjy.Jfile2Y(jsonfile)

	ver1 := int64(1)
	go xjy.YAMLAllValuesAsync(yamlStr, "id", false, true, func(p, v, id string) {
		tuple, err := messages.NewTuple(id, p, v)
		PE(err)
		tuple.Version = ver1
		ver1++
		err = n3pub.Publish(tuple, nameSpace, contextName)
		PE(err)
		pln("---", *tuple)
	}, done)

	pf("finish: %d\n", <-done)
	log.Println("messages sent")

	/*******************************************/

	// nameSpace := "Aa5fKf2UmyfCufY6JFmQpX12j1jjDFSUfbFUEE92t2nx"
	// contextName := "abc-sif"

	// done := make(chan int, 2)

	// xmlfile := "../xjy/files/staffpersonal.xml"
	// yamlStr := xjy.Xfile2Y(xmlfile)
	// ioutil.WriteFile("../xjy/files/staffpersonal.yaml", []byte(yamlStr), 0666)

	// ver1 := int64(1)
	// go xjy.YAMLAllValuesAsync(yamlStr, "RefId", true, true, func(p, v, id string) {
	// 	tuple, err := messages.NewTuple(id, p, v)
	// 	PE(err)
	// 	tuple.Version = ver1
	// 	ver1++
	// 	err = n3pub.Publish(tuple, nameSpace, contextName)
	// 	PE(err)
	// 	fmt.Println("---", *tuple)
	// }, done)

	// xmlbytes, err := ioutil.ReadFile(xmlfile)
	// PE(err)
	// ver2 := int64(1)
	// go xjy.XMLStructAsync(string(xmlbytes), "RefId", true, func(p, v string) {
	// 	tuple, err := messages.NewTuple(p, "::", v)
	// 	PE(err)
	// 	tuple.Version = ver2
	// 	ver2++
	// 	err = n3pub.Publish(tuple, nameSpace, contextName)
	// 	PE(err)
	// }, done)

	// fmt.Printf("finish1: %d\n", <-done)
	// fmt.Printf("finish2: %d\n", <-done)
	// log.Println("messages sent")
}
