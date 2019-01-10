package send

import (
	"../xjy"
	u "github.com/cdutwhu/util"
	"github.com/nsip/n3-messages/messages"
	"github.com/nsip/n3-messages/n3grpc"
)

// SIF is
func SIF(str string) (cnt int) {
	content := u.Str(str)
	PC(content.L() == 0 || !content.IsXMLSegSimple(), fEf("Incoming string is invalid xml segment"))

	if n3pub == nil {
		n3pub, e = n3grpc.NewPublisher(sendTo, sendToPort)
		PE(e)
	}

	done := make(chan int, 2)

	go xjy.YAMLAllValuesAsync(xjy.Xstr2Y(content.V()), "RefId", true, true, func(p, v, id string) {
		tuple, _ := messages.NewTuple(id, p, v)
		tuple.Version = verSIF1
		verSIF1++
		PE(n3pub.Publish(tuple, nameSpace, ctxNameSIF))
		// fPln("---", *tuple)
		cnt++
	}, done)

	cnt1 := 0
	go xjy.XMLStructAsync(content.V(), "RefId", true, func(p, v string) {
		tuple, _ := messages.NewTuple(p, "::", v)
		tuple.Version = verSIF2
		verSIF2++
		PE(n3pub.Publish(tuple, nameSpace, ctxNameSIF))
		cnt1++
	}, done)

	fPf("sif sent 1: %d\n", <-done)
	fPf("sif sent 2: %d\n", <-done)

	lPln(fSpf("%06d data tuples sent, %06d struct tuples sent\n", cnt, cnt1))
	return cnt
}
