package n3ctest

import (
	"io/ioutil"
	"testing"
	"time"

	c "../config"
	"../send"
)

func TestN3LoadConfig(t *testing.T) {
	cfg := &c.Config{}
	cfg.Load("../config/config.toml")
	fPln(cfg.Grpc)
	fPln(cfg.Filewatcher)
	send.Init(cfg)
}

func TestN3PubSIF(t *testing.T) {
	defer func() { PH(recover(), "./log.txt", true) }()
	TestN3LoadConfig(t)

	xmlfile := "../xjy/files/staffpersonal.xml"
	bytes, e := ioutil.ReadFile(xmlfile)
	PE(e)
	n := send.SIF(string(bytes))
	fPln(n)
	time.Sleep(2 * time.Second)
}

func TestN3PubXAPI(t *testing.T) {
	defer func() { PH(recover(), "./log.txt", true) }()
	TestN3LoadConfig(t)

	jsonfile := "../xjy/files/xapifile.json"
	bytes, e := ioutil.ReadFile(jsonfile)
	PE(e)
	n := send.XAPI(string(bytes))
	fPln(n)
	time.Sleep(2 * time.Second)
}
