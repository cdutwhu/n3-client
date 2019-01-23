package send

import (
	"io/ioutil"
	"testing"
	"time"

	c "../config"
)

func TestJunk(t *testing.T) {
	defer func() { uPH(recover(), Cfg.Global.ErrLog, true) }()
	TestN3LoadConfig(t)
	Junk(10000)
	time.Sleep(2 * time.Second)
}

/************************************************************/

func TestN3LoadConfig(t *testing.T) {
	cfg := c.GetConfig("../config/config.toml")
	fPln(cfg.Grpc)
	fPln(cfg.Filewatcher)
	Init(cfg)
	fPln(cfg.Path)
}

func TestSendSIF(t *testing.T) {
	defer func() { uPH(recover(), Cfg.Global.ErrLog, true) }()
	TestN3LoadConfig(t)

	xmlfile := "../inbound/sif/staffpersonal.xml"
	// xmlfile := "../inbound/sif/nswdig.xml"
	bytes, e := ioutil.ReadFile(xmlfile)
	uPE(e)
	nV, nS, nA := SIF(string(bytes))
	fPln(nV, nS, nA)
	time.Sleep(2 * time.Second)
}

func TestSendXAPI(t *testing.T) {
	defer func() { uPH(recover(), Cfg.Global.ErrLog, true) }()
	TestN3LoadConfig(t)

	jsonfile := "../inbound/xapi/xapifile.json"
	bytes, e := ioutil.ReadFile(jsonfile)
	uPE(e)
	n := XAPI(string(bytes))
	fPln(n)
	time.Sleep(2 * time.Second)
}
