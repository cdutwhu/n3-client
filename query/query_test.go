package query

import (
	"testing"

	c "../config"
)

func TestN3LoadConfig(t *testing.T) {
	cfg := c.GetConfig("../config/config.toml")
	// fPln(cfg.Grpc)
	// fPln(cfg.Filewatcher)
	Init(cfg)
}

func TestQuerySIF(t *testing.T) {
	defer func() { uPH(recover(), Cfg.Global.ErrLog, true) }()
	TestN3LoadConfig(t)
	s, p, o, v := SIF("D3E34F41-9D75-101A-8C3D-00AA001A1652", "StaffPersonal") //SIF("0E11C01D-54A2-4E9F-8C67-4FE2540BA6C8", "StaffPersonal")
	// s, p, o, v := SIF("D0E7421A-38AE-48D0-985F-D5525D32B56D", "TeachingGroup") /* context must end with '-sif' */
	fPln(len(s))
	for i := range s {
		fPln("----------------------------------------------------")
		fPf("%d # %d: Reply: %s\n%s\n%s \n", i, v[i], s[i], p[i], o[i])
	}
	fPln()
}

func TestQueryXAPI(t *testing.T) {
	defer func() { uPH(recover(), Cfg.Global.ErrLog, true) }()
	TestN3LoadConfig(t)
	s, p, o, v := XAPI("D3E34F41-9D75-101A-8C3D-00AA001A1652", "result") /* context must end with '-xapi' */
	fPln(len(s))
	for i := range s {
		fPln("----------------------------------------------------")
		fPf("%d # %d: Reply: %s\n%s\n%s \n", i, v[i], s[i], p[i], o[i])
	}
	fPln()
}
