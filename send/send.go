package send

import (
	c "../config"
	"../xjy"
	u "github.com/cdutwhu/go-util"
	"github.com/nsip/n3-messages/messages"
	"github.com/nsip/n3-messages/n3grpc"
)

// Init :
func Init(cfg *c.Config) {
	c.Cfg = cfg
	if c.N3pub == nil {
		c.N3pub, e = n3grpc.NewPublisher(c.Cfg.Grpc.Server, c.Cfg.Grpc.Port)
		PE(e)
	}
}

// SIF is
func SIF(str string) (cnt int) {
	if c.Cfg == nil || c.N3pub == nil {
		panic("Missing Init, do 'Init(&config) before sending'")
	}

	content := u.Str(str)
	PC(content.L() == 0 || !content.IsXMLSegSimple(), fEf("Incoming string is invalid xml segment"))

	done := make(chan int, 2)
	go xjy.YAMLScanAsync(xjy.Xstr2Y(content.V()), "RefId", xjy.SIF, true, func(p, v, id string) {
		tuple, _ := messages.NewTuple(id, p, v)
		tuple.Version = verSIF1
		verSIF1++
		PE(c.N3pub.Publish(tuple, c.Cfg.Grpc.Namespace, c.Cfg.Grpc.Ctxsif))
		// fPln("---", *tuple)
		cnt++
	}, done)
	cnt1 := 0
	go xjy.XMLStructAsync(content.V(), "RefId", true, func(p, v string) {
		tuple, _ := messages.NewTuple(p, "::", v)
		tuple.Version = verSIF2
		verSIF2++
		PE(c.N3pub.Publish(tuple, c.Cfg.Grpc.Namespace, c.Cfg.Grpc.Ctxsif))
		cnt1++
	}, done)
	fPf("sif decode 1: %d\n", <-done)
	fPf("sif decode 2: %d\n", <-done)

	lPln(fSpf("<%06d> data tuples decoded, <%06d> struct tuples decoded\n", cnt, cnt1))
	return cnt
}

// XAPI is
func XAPI(str string) (cnt int) {
	if c.Cfg == nil {
		panic("Missing Send Init, do 'Init(&config) before sending'")
	}

	content := u.Str(str)
	PC(content.L() == 0 || !content.IsJSON(), fEf("Incoming string is invalid json"))

	done := make(chan int)
	go xjy.YAMLScanAsync(xjy.Jstr2Y(content.V()), "id", xjy.XAPI, true, func(p, v, id string) {
		tuple, _ := messages.NewTuple(id, p, v)
		tuple.Version = verXAPI
		verXAPI++
		PE(c.N3pub.Publish(tuple, c.Cfg.Grpc.Namespace, c.Cfg.Grpc.Ctxxapi))
		// fPln("---", *tuple)
		cnt++
	}, done)
	fPf("xapi decoded : %d\n", <-done)

	lPln(fSpf("<%06d> tuples decoded\n", cnt))
	return cnt
}
