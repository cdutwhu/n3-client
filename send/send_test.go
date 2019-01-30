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
	cfg := c.GetConfig("./config.toml", "../config/config.toml")
	// fPln(cfg.Grpc)
	// fPln(cfg.Filewatcher)
	// fPln(cfg.Path)
	Init(cfg)
}

func TestSendSIF(t *testing.T) {
	defer func() { uPH(recover(), Cfg.Global.ErrLog, true) }()
	TestN3LoadConfig(t)

	// xmlfile := "../inbound/sif/staffpersonal.xml"
	xmlfile := "../inbound/sif/nswdig.xml"
	bytes, e := ioutil.ReadFile(xmlfile)
	uPE(e)
	SIF(string(bytes))
}

func TestSendXAPI(t *testing.T) {
	defer func() { uPH(recover(), Cfg.Global.ErrLog, true) }()
	TestN3LoadConfig(t)

	jsonfile := "../inbound/xapi/xapifile.json"
	bytes, e := ioutil.ReadFile(jsonfile)
	uPE(e)
	XAPI(string(bytes))
}
