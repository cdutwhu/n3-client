package config

import "github.com/burntsushi/toml"

type filewatcher struct {
	Dirsif  string
	Dirxapi string
}

type grpc struct {
	Namespace string
	Ctxsif    string
	Ctxxapi   string
	Server    string
	Port      int
}

// Config is toml
type Config struct {
	Filewatcher filewatcher
	Grpc        grpc
}

// Load is
func (cfg *Config) Load(cfgfile string) {
	defer func() { uPH(recover(), "./log.txt", true) }()
	_, e := toml.DecodeFile(cfgfile, cfg)
	uPE(e)
}
