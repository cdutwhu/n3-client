package query

import (
	c "../config"
	g "../global"
	u "github.com/cdutwhu/go-util"
	"github.com/nsip/n3-messages/messages/pb"
	"github.com/nsip/n3-messages/n3grpc"
)

// Init :
func Init(config *c.Config) {
	uPC(config == nil, fEf("Init Config"))
	Cfg = config
	if g.N3pub == nil {
		g.N3pub, e = n3grpc.NewPublisher(Cfg.Grpc.Server, Cfg.Grpc.Port)
		uPE(e)
	}
}

func query(t qType, sp []string) (s, p, o []string, v []int64) {
	// uPC(Cfg == nil || g.N3pub == nil, fEf("Missing Init, do 'Init(&config) before querying'\n"))

	if Cfg == nil || g.N3pub == nil {
		Cfg = c.GetConfig("./config.toml", "../config/config.toml")
		Init(Cfg)
	}

	ctx := u.CaseAssign(t, SIF, XAPI, Cfg.Grpc.CtxSif, Cfg.Grpc.CtxXapi).(string)
	qTuple := &pb.SPOTuple{
		Subject:   sp[0],
		Predicate: sp[1],
		Object:    "",
	}
	for _, t := range g.N3pub.Query(qTuple, Cfg.Grpc.Namespace, ctx) {
		s, p, o, v = append(s, t.Subject), append(p, t.Predicate), append(o, t.Object), append(v, t.Version)
	}
	return
}

// Sif :
func Sif(sp ...string) (s, p, o []string, v []int64) {
	return query(SIF, sp)
}

// Xapi :
func Xapi(sp ...string) (s, p, o []string, v []int64) {
	return query(XAPI, sp)
}
