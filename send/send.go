package send

import (
	c "../config"
	g "../global"
	"../xjy"
	u "github.com/cdutwhu/go-util"
	"github.com/nsip/n3-messages/messages"
	"github.com/nsip/n3-messages/n3grpc"
)

// Junk :
func Junk(n int) {
	uPC(Cfg == nil || g.N3pub == nil, fEf("Missing Init, do 'Init(&config) before sending'\n"))
	for i := 0; i < n; i++ {
		tuple, _ := messages.NewTuple("sub", "pre", "obj")
		tuple.Version = Cfg.Temp.VerSif
		Cfg.Temp.VerSif++
		uPE(g.N3pub.Publish(tuple, Cfg.Grpc.Namespace, Cfg.Grpc.CtxSif))
	}
	Cfg.Save()
}

/************************************************************/

// Init :
func Init(config *c.Config) {
	uPC(config == nil, fEf("Init Config"))
	Cfg = config
	if g.N3pub == nil {
		g.N3pub, e = n3grpc.NewPublisher(Cfg.Grpc.Server, Cfg.Grpc.Port)
		uPE(e)
	}
}

// SIF is
func SIF(str string) (cntV, cntS, cntA int) {
	uPC(Cfg == nil || g.N3pub == nil, fEf("Missing Send Init, do 'Init(&config) before sending'\n"))

	content := u.Str(str)
	uPC(content.L() == 0 || !content.IsXMLSegSimple(), fEf("Incoming string is invalid xml segment\n"))

	doneV := make(chan int)
	go xjy.YAMLScanAsync(xjy.Xstr2Y(content.V()), "RefId", xjy.SIF, true,
		func(p, v, id string) {
			tuple, _ := messages.NewTuple(id, p, v)
			tuple.Version = Cfg.Temp.VerSif
			Cfg.Temp.VerSif++
			uPE(g.N3pub.Publish(tuple, Cfg.Grpc.Namespace, Cfg.Grpc.CtxSif))
			// fPln("---", *tuple)
			cntV++
		},
		doneV)
	<-doneV

	doneS := make(chan int)
	go xjy.XMLStructAsync(content.V(), "RefId", true,
		func(p, v string) {
			tuple, _ := messages.NewTuple(p, "::", v)
			tuple.Version = Cfg.Temp.VerSif
			Cfg.Temp.VerSif++
			uPE(g.N3pub.Publish(tuple, Cfg.Grpc.Namespace, Cfg.Grpc.CtxSif))
			cntS++
		},
		func(p, objid string, arrcnt int) {
			tuple, _ := messages.NewTuple(p, objid, fSpf("%d", arrcnt))
			tuple.Version = Cfg.Temp.VerSif
			Cfg.Temp.VerSif++
			uPE(g.N3pub.Publish(tuple, Cfg.Grpc.Namespace, Cfg.Grpc.CtxSif))
			cntA++
			fPln("---", *tuple)
		},
		doneS)
	<-doneS

	Cfg.Save()
	lPln(fSpf("<%06d> data tuples decoded, <%06d> struct tuples decoded, <%06d> array tuples decoded\n", cntV, cntS, cntA))
	return
}

// XAPI is
func XAPI(str string) (cnt int) {
	uPC(Cfg == nil, fEf("Missing Send Init, do 'Init(&config) before sending'\n"))

	content := u.Str(str)
	uPC(content.L() == 0 || !content.IsJSON(), fEf("Incoming string is invalid json\n"))

	done := make(chan int)
	go xjy.YAMLScanAsync(xjy.Jstr2Y(content.V()), "id", xjy.XAPI, true, func(p, v, id string) {
		tuple, _ := messages.NewTuple(id, p, v)
		tuple.Version = Cfg.Temp.VerXapi
		Cfg.Temp.VerXapi++
		uPE(g.N3pub.Publish(tuple, Cfg.Grpc.Namespace, Cfg.Grpc.CtxXapi))
		// fPln("---", *tuple)
		cnt++
	}, done)
	fPf("xapi decoded : %d\n", <-done)

	Cfg.Save()
	lPln(fSpf("<%06d> tuples decoded\n", cnt))
	return cnt
}
