package send

import (
	"../xjy"
	u "github.com/cdutwhu/util"
	"github.com/nsip/n3-messages/messages"
	"github.com/nsip/n3-messages/n3grpc"
)

// XAPI is
func XAPI(str string) (cnt int) {
	content := u.Str(str)
	PC(content.L() == 0 || !content.IsJSON(), fEf("Incoming string is invalid json"))

	if n3pub == nil {
		n3pub, e = n3grpc.NewPublisher(sendTo, sendToPort)
	}
	done := make(chan int)

	go xjy.YAMLAllValuesAsync(xjy.Jstr2Y(content.V()), "id", false, true, func(p, v, id string) {
		tuple, _ := messages.NewTuple(id, p, v)
		tuple.Version = verXAPI
		verXAPI++
		PE(n3pub.Publish(tuple, nameSpace, ctxNameXAPI))
		// fPln("---", *tuple)
		cnt++
	}, done)

	fPf("xapi sent : %d\n", <-done)

	lPln(fSpf("%d tuples sent\n", cnt))
	return cnt
}
