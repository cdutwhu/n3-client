package send

import (
	"time"

	c "../config"
	g "../global"
	q "../query"
	"../xjy"
	u "github.com/cdutwhu/go-util"
	"github.com/google/uuid"
	"github.com/nsip/n3-messages/messages"
	"github.com/nsip/n3-messages/n3grpc"
)

// Junk :
func Junk(n int) {
	uPC(Cfg == nil || g.N3pub == nil, fEf("Missing Init, do 'Init(&config) before sending'\n"))
	for i := 0; i < n; i++ {
		tuple, e := messages.NewTuple("sub", "pre", "obj")
		uPE(e)
		tuple.Version = Cfg.Temp.VerSif
		Cfg.Temp.VerSif++
		uPE(g.N3pub.Publish(tuple, Cfg.Grpc.Namespace, Cfg.Grpc.CtxSif))
	}
	Cfg.Save()
}

// Terminate :
func Terminate(t sType, n int) string {
	// uPC(Cfg == nil || g.N3pub == nil, fEf("Missing Init, do 'Init(&config) before sending'\n"))
	if Cfg == nil || g.N3pub == nil {
		Cfg = c.GetConfig("./config.toml", "../config/config.toml")
		Init(Cfg)
	}

	ctx := u.CaseAssign(t, SIF, XAPI, Cfg.Grpc.CtxSif, Cfg.Grpc.CtxXapi).(string)
	termID := uuid.New().String()
	tuple, e := messages.NewTuple(termID, TERMMARK, fSpf("%d", n))
	uPE(e)
	tuple.Version = Cfg.Temp.VerSif
	Cfg.Temp.VerSif++
	uPE(g.N3pub.Publish(tuple, Cfg.Grpc.Namespace, ctx))
	Cfg.Save()
	return termID
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

// Sif is
func Sif(str string) (cntV, cntS, cntA int, termID string) {
	uPC(Cfg == nil || g.N3pub == nil, fEf("Missing Send Init, do 'Init(&config) before sending'\n"))

	content := u.Str(str)
	uPC(content.L() == 0 || !content.IsXMLSegSimple(), fEf("Incoming string is invalid xml segment\n"))

	xjy.XMLModelInfo(content.V(), "RefId", true,
		func(p, v string) {
			tuple, e := messages.NewTuple(u.Str(p).RmPrefix(HEADTRIM), "::", v)
			uPE(e)
			tuple.Version = Cfg.Temp.VerSif
			Cfg.Temp.VerSif++
			uPE(g.N3pub.Publish(tuple, Cfg.Grpc.Namespace, Cfg.Grpc.CtxSif))
			cntS++
		},
		func(p, objID string, arrCnt int) {
			tuple, e := messages.NewTuple(u.Str(p).RmPrefix(HEADTRIM), objID, fSpf("%d", arrCnt))
			uPE(e)
			tuple.Version = Cfg.Temp.VerSif
			Cfg.Temp.VerSif++
			uPE(g.N3pub.Publish(tuple, Cfg.Grpc.Namespace, Cfg.Grpc.CtxSif))
			cntA++
		},
	)

	doneV := make(chan int)
	go xjy.YAMLScanAsync(xjy.Xstr2Y(content.V()), "RefId", xjy.SIF, true,
		func(p, v, id string) {
			tuple, e := messages.NewTuple(id, u.Str(p).RmPrefix(HEADTRIM), v)
			uPE(e)
			tuple.Version = Cfg.Temp.VerSif
			Cfg.Temp.VerSif++
			uPE(g.N3pub.Publish(tuple, Cfg.Grpc.Namespace, Cfg.Grpc.CtxSif))
			// fPln("---", *tuple)
			cntV++
		},
		doneV)
	<-doneV

	Cfg.Save()
	lPln(fSpf("<%06d> data tuples sent, <%06d> struct tuples sent, <%06d> array tuples sent\n", cntV, cntS, cntA))

	termID = Terminate(SIF, cntV+cntS+cntA)
CHECK:
	if _, _, _, v := q.Sif(termID, TERMMARK); v == nil || len(v) == 0 {
		time.Sleep(DELAY * time.Millisecond)
		goto CHECK
	}
	return
}

// Xapi is
func Xapi(str string) (cnt int, termID string) {
	uPC(Cfg == nil, fEf("Missing Send Init, do 'Init(&config) before sending'\n"))

	content := u.Str(str)
	uPC(content.L() == 0 || !content.IsJSON(), fEf("Incoming string is invalid json\n"))

	done := make(chan int)
	go xjy.YAMLScanAsync(xjy.Jstr2Y(content.V()), "id", xjy.XAPI, true, func(p, v, id string) {
		tuple, e := messages.NewTuple(id, p, v)
		uPE(e)
		tuple.Version = Cfg.Temp.VerXapi
		Cfg.Temp.VerXapi++
		uPE(g.N3pub.Publish(tuple, Cfg.Grpc.Namespace, Cfg.Grpc.CtxXapi))
		// fPln("---", *tuple)
		cnt++
	}, done)
	<-done

	Cfg.Save()
	lPln(fSpf("<%06d> tuples sent\n", cnt))

	termID = Terminate(XAPI, cnt)
CHECK:
	if _, _, _, v := q.Xapi(termID, TERMMARK); v == nil || len(v) == 0 {
		time.Sleep(DELAY * time.Millisecond)
		goto CHECK
	}
	return
}
