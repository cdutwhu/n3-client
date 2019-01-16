package query

import (
	"testing"
	"time"

	c "../config"
)

func TestN3LoadConfig(t *testing.T) {
	cfg := &c.Config{}
	cfg.Load("../config/config.toml")
	//fPln(cfg.Grpc)
	//fPln(cfg.Filewatcher)
	Init(cfg)
}

func TestQuerySIF(t *testing.T) {
	defer func() { PH(recover(), "./log.txt", true) }()
	TestN3LoadConfig(t)
	s, p, o, v := SIF("D3E34F41-9D75-101A-8C3D-00AA001A1652", "StaffPersonal")
	fPln(len(s))
	for i := range s {
		fPln("----------------------------------------------------")
		fPf("%d # %d: Reply: %s\n%s\n%s \n", i, v[i], s[i], p[i], o[i])
	}
	fPln()
	time.Sleep(2 * time.Second)
}

func TestQueryXAPI(t *testing.T) {
	defer func() { PH(recover(), "./log.txt", true) }()
	TestN3LoadConfig(t)
	s, p, o, v := XAPI("", "")
	fPln(len(s))
	for i := range s {
		fPln("----------------------------------------------------")
		fPf("%d # %d: Reply: %s\n%s\n%s \n", i, v[i], s[i], p[i], o[i])
	}
	fPln()
	time.Sleep(2 * time.Second)
}