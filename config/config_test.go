package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	cfg := GetConfig("./config.toml")
	fPln(cfg.Grpc)
	fPln(cfg.Filewatcher)
	fPln(cfg.Global.ErrLog)
}

func TestSave(t *testing.T) {
	cfg := GetConfig("./config.toml")
	cfg.Temp.VerSif = 10
	cfg.Save()

	cfg1 := GetConfig("./config.toml")
	fPln(cfg1.Grpc)
	fPln(cfg1.Filewatcher)
	fPln(cfg1.Temp.VerSif, cfg1.Temp.VerXapi)
}
