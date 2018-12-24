package n3pub

import (
	"fmt"
	"testing"

	"github.com/nsip/n3-messages/messages/pb"
	"github.com/nsip/n3-messages/n3grpc"
)

func TestN3Pub(t *testing.T) {
	defer func() { PH(recover(), "./log.txt", true) }()

	nameSpace := "Aa5fKf2UmyfCufY6JFmQpX12j1jjDFSUfbFUEE92t2nx"
	contextName := "abc"

	n3pub, err := n3grpc.NewPublisher("localhost", 5777)
	PE(err)
	defer n3pub.Close()

	/*******************************************/

	qTuple := &pb.SPOTuple{
		Subject:   "D3E34F41-9D75-101A-8C3D-00AA001A1652", //"0CD1D251-C873-4222-B01D-EC375FC8A1CC", //
		Predicate: "StaffPersonal",                        //"sif.TeachingGroup",                        //
		Object:    "",
	}

	fmt.Printf("Query: %v\n", *qTuple)

	ts := n3pub.Query(qTuple, nameSpace, contextName)
	fmt.Println(len(ts))
	for i, t := range ts {
		fmt.Println("----------------------------------------------------")
		fmt.Printf("%d # %d: Reply: %s\n%s\n%s \n", i, t.Version, t.Subject, t.Predicate, t.Object)
	}

	/*******************************************/
	/*
		done := make(chan int, 2)

		xmlfile := "../xjy/files/staffpersonal.xml"
		yamlStr := xjy.Xfile2Y(xmlfile)

		ioutil.WriteFile("../xjy/files/staffpersonal.yaml", []byte(yamlStr), 0666)

		ver1 := int64(1)
		go xjy.YAMLAllValuesAsync(yamlStr, "RefId", true, func(p, v, id string) {
			tuple, err := messages.NewTuple(id, p, v)
			PE(err)
			tuple.Version = ver1
			ver1++
			err = n3pub.Publish(tuple, nameSpace, contextName)
			PE(err)
			fmt.Println("---", *tuple)
		}, done)

		xmlbytes, err := ioutil.ReadFile(xmlfile)
		PE(err)
		ver2 := int64(1)
		go xjy.XMLStructAsync(string(xmlbytes), "RefId", true, func(p, v string) {
			tuple, err := messages.NewTuple(p, "::", v)
			PE(err)
			tuple.Version = ver2
			ver2++
			err = n3pub.Publish(tuple, nameSpace, contextName)
			PE(err)
		}, done)

		fmt.Printf("finish1: %d\n", <-done)
		fmt.Printf("finish2: %d\n", <-done)
		log.Println("messages sent")
	*/
}
