package send

import (
	c "../config"
	"../xjy"
	u "github.com/cdutwhu/go-util"
	"github.com/nsip/n3-messages/messages"
	"github.com/nsip/n3-messages/n3grpc"
)

// Junk :
func Junk(n int) {
	if c.Cfg == nil || c.N3pub == nil {
		panic("Missing Init, do 'Init(&config) before sending'")
	}

	for i := 0; i < n; i++ {
		tuple, _ := messages.NewTuple("sub", "pre", "obj")
		tuple.Version = verSIF
		verSIF++
		PE(c.N3pub.Publish(tuple, c.Cfg.Grpc.Namespace, c.Cfg.Grpc.Ctxsif))
	}
}

/************************************************************/

// Init :
func Init(cfg *c.Config) {
	c.Cfg = cfg
	if c.N3pub == nil {
		c.N3pub, e = n3grpc.NewPublisher(c.Cfg.Grpc.Server, c.Cfg.Grpc.Port)
		PE(e)
	}
}

// SIF is
func SIF(str string) (cntV, cntS int) {
	if c.Cfg == nil || c.N3pub == nil {
		panic("Missing Init, do 'Init(&config) before sending'")
	}

	content := u.Str(str)
	PC(content.L() == 0 || !content.IsXMLSegSimple(), fEf("Incoming string is invalid xml segment"))

	doneV := make(chan int)
	go xjy.YAMLScanAsync(xjy.Xstr2Y(content.V()), "RefId", xjy.SIF, true,
		func(p, v, id string) {
			tuple, _ := messages.NewTuple(id, p, v)
			tuple.Version = verSIF
			verSIF++
			PE(c.N3pub.Publish(tuple, c.Cfg.Grpc.Namespace, c.Cfg.Grpc.Ctxsif))
			// fPln("---", *tuple)
			cntV++
		},
		doneV)
	<-doneV

	doneS := make(chan int)
	go xjy.XMLStructAsync(content.V(), "RefId", true,
		func(p, v string) {
			tuple, _ := messages.NewTuple(p, "::", v)
			tuple.Version = verSIF
			verSIF++
			PE(c.N3pub.Publish(tuple, c.Cfg.Grpc.Namespace, c.Cfg.Grpc.Ctxsif))
			cntS++
		},
		func(p, objid string, arrcnt int) {
			tuple, _ := messages.NewTuple(p, objid, fSpf("%d", arrcnt))
			tuple.Version = verSIF
			verSIF++
			PE(c.N3pub.Publish(tuple, c.Cfg.Grpc.Namespace, c.Cfg.Grpc.Ctxsif))
			fPln("---", *tuple)
		},
		doneS)
	<-doneS

	lPln(fSpf("<%06d> data tuples decoded, <%06d> struct tuples decoded\n", cntV, cntS))
	return
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
