package n3ctest

import (
	"fmt"
	"testing"

	"github.com/nsip/n3-messages/messages/pb"
	"github.com/nsip/n3-messages/n3grpc"
)

func TestN3Query(t *testing.T) {
	defer func() { PH(recover(), "./log.txt", true) }()

	nameSpace := "Aa5fKf2UmyfCufY6JFmQpX12j1jjDFSUfbFUEE92t2nx"
	//contextName := "abc-xapi"
	contextName := "abc-sif"

	to := "192.168.76.10"
	// to := "localhost"
	n3pub, err := n3grpc.NewPublisher(to, 5777)
	PE(err)
	defer n3pub.Close()

	/*******************************************/

	qTuple := &pb.SPOTuple{
		Subject:   "D3E34F41-9D75-101A-8C3D-00AA001A1652", //"0CD1D251-C873-4222-B01D-EC375FC8A1CC", //
		Predicate: "StaffPersonal.Title",                  //"sif.TeachingGroup",                        //
		Object:    "",
	}

	fmt.Printf("Query: %v\n", *qTuple)

	ts := n3pub.Query(qTuple, nameSpace, contextName)
	fmt.Println(len(ts))
	for i, t := range ts {
		fmt.Println("----------------------------------------------------")
		fmt.Printf("%d # %d: Reply: %s\n%s\n%s \n", i, t.Version, t.Subject, t.Predicate, t.Object)
	}
}
