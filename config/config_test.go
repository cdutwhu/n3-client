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
	cfg.Temp.Versif = 10
	cfg.Save()

	cfg1 := GetConfig("./config.toml")
	fPln(cfg1.Grpc)
	fPln(cfg1.Filewatcher)
	fPln(cfg1.Temp.Versif, cfg1.Temp.Verxapi)
}
