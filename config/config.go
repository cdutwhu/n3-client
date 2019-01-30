package config

import (
	"os"

	"github.com/burntsushi/toml"
)

type global struct {
	ErrLog string
}

type filewatcher struct {
	DirSif  string
	DirXapi string
}

type rest struct {
	Port     int
	SifPath  string
	XapiPath string
}

type grpc struct {
	Namespace string
	CtxSif    string
	CtxXapi   string
	Server    string
	Port      int
}

type temp struct {
	VerSif  int64
	VerXapi int64
}

// Config is toml
type Config struct {
	Path        string
	Global      global
	Filewatcher filewatcher
	Rest        rest
	Grpc        grpc
	Temp        temp
}

// GetConfig :
func GetConfig(cfgfiles ...string) *Config {
	for _, f := range cfgfiles {
		if _, e := os.Stat(f); e == nil {
			cfg := &Config{Path: f}
			return cfg.set()
		}
	}
	panic("config file error")
}

// set is
func (cfg *Config) set() *Config {
	defer func() { uPH(recover(), "./log.txt", true) }()
	path := cfg.Path /* make a copy of original path for restoring */
	_, e := toml.DecodeFile(cfg.Path, cfg)
	cfg.Path = path
	uPE(e)
	return cfg
}

// Save is
func (cfg *Config) Save() {
	defer func() { uPH(recover(), cfg.Global.ErrLog, true) }()
	f, e := os.OpenFile(cfg.Path, os.O_WRONLY|os.O_TRUNC, 0666)
	uPE(e)
	defer f.Close()
	uPE(toml.NewEncoder(f).Encode(cfg))
}
