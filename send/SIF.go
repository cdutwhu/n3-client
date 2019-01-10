package send

import (
	"../xjy"
	u "github.com/cdutwhu/util"
	"github.com/nsip/n3-messages/messages"
	"github.com/nsip/n3-messages/n3grpc"
)

// SIF is
func SIF(str string) int {
	content := u.Str(str)
	PC(content.L() == 0 || !content.IsXMLSegSimple(), fEf("Incoming string is invalid xml segment"))

	if n3pub == nil {
		n3pub, e = n3grpc.NewPublisher(sendTo, sendToPort)
		PE(e)
	}
	done, cnt := make(chan int, 2), 0

	go xjy.YAMLAllValuesAsync(xjy.Xstr2Y(content.V()), "RefId", true, true, func(p, v, id string) {
		tuple, _ := messages.NewTuple(id, p, v)
		tuple.Version = verSIF1
		verSIF1++
		n3pub.Publish(tuple, nameSpace, ctxNameSIF)
		// fPln("---", *tuple)
		cnt++
	}, done)

	go xjy.XMLStructAsync(content.V(), "RefId", true, func(p, v string) {
		tuple, _ := messages.NewTuple(p, "::", v)
		tuple.Version = verSIF2
		verSIF2++
		n3pub.Publish(tuple, nameSpace, ctxNameSIF)
	}, done)

	fPf("sif sent 1: %d\n", <-done)
	fPf("sif sent 2: %d\n", <-done)
	lPln(fSpf("%d tuples sent\n", cnt))
	return cnt
}
