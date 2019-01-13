package n3ctest

import (
	"io/ioutil"
	"testing"
	"time"

	c "../config"
	"../send"
)

func TestN3PubSIF(t *testing.T) {
	defer func() { PH(recover(), "./log.txt", true) }()

	cfg := &c.Config{}
	cfg.Load("../config/config.toml")
	fPln(cfg.Grpc)
	fPln(cfg.Filewatcher)
	send.Init(cfg)

	xmlfile := "../xjy/files/staffpersonal.xml"
	bytes, e := ioutil.ReadFile(xmlfile)
	PE(e)
	n := send.SIF(string(bytes))
	fPln(n)
	time.Sleep(2 * time.Second)
}

func TestN3PubXAPI(t *testing.T) {
	defer func() { PH(recover(), "./log.txt", true) }()

	cfg := &c.Config{}
	cfg.Load("../config/config.toml")
	fPln(cfg.Grpc)
	fPln(cfg.Filewatcher)
	send.Init(cfg)

	jsonfile := "../xjy/files/xapifile.json"
	bytes, e := ioutil.ReadFile(jsonfile)
	PE(e)
	n := send.XAPI(string(bytes))
	fPln(n)
	time.Sleep(2 * time.Second)
}
