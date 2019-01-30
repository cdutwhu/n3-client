package query

import (
	"testing"

	c "../config"
)

func TestN3LoadConfig(t *testing.T) {
	cfg := c.GetConfig("./config.toml", "../config/config.toml")
	// fPln(cfg.Grpc)
	// fPln(cfg.Filewatcher)
	Init(cfg)
}

func TestQuerySIF(t *testing.T) {
	defer func() { uPH(recover(), Cfg.Global.ErrLog, true) }()
	TestN3LoadConfig(t)
	s, p, o, v := SIF("0E11C01D-54A2-4E9F-8C67-4FE2540BA6C8", "StaffPersonal") // SIF("D3E34F41-9D75-101A-8C3D-00AA001A1652", "StaffPersonal") //
	// s, p, o, v := SIF("9269671A-BB89-4281-B20D-668C1D7FFD05", "TeachingGroup") /* context must end with '-sif' */
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
