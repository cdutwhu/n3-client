package query

import (
	c "../config"
	"github.com/nsip/n3-messages/messages/pb"
	"github.com/nsip/n3-messages/n3grpc"
)

type qType int

const (
	qtSif  qType = 0
	qtXapi qType = 1
)

// Init :
func Init(cfg *c.Config) {
	c.Cfg = cfg
	if c.N3pub == nil {
		c.N3pub, e = n3grpc.NewPublisher(c.Cfg.Grpc.Server, c.Cfg.Grpc.Port)
		PE(e)
	}
}

func query(t qType, sp []string) (s, p, o []string, v []int64) {
	ctx := ""
	switch t {
	case qtSif:
		ctx = c.Cfg.Grpc.Ctxsif
	case qtXapi:
		ctx = c.Cfg.Grpc.Ctxxapi
	}

	if c.Cfg == nil || c.N3pub == nil {
		panic("Missing Init, do 'Init(&config) before querying'")
	}
	qTuple := &pb.SPOTuple{
		Subject:   sp[0],
		Predicate: sp[1],
		Object:    "",
	}
	for _, t := range c.N3pub.Query(qTuple, c.Cfg.Grpc.Namespace, ctx) {
		s, p, o, v = append(s, t.Subject), append(p, t.Predicate), append(o, t.Object), append(v, t.Version)
	}
	return
}

// SIF :
func SIF(sp ...string) (s, p, o []string, v []int64) {
	return query(qtSif, sp)
}

// XAPI :
func XAPI(sp ...string) (s, p, o []string, v []int64) {
	return query(qtXapi, sp)
}
