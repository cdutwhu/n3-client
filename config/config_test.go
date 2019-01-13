package config

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	cfg := &Config{}
	cfg.Load("./config.toml")
	fmt.Println(cfg.Grpc)
	fmt.Println(cfg.Filewatcher)
}
