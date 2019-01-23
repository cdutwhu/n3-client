package main

import (
	"testing"

	c "./config"
	q "./query"
	s "./send"
	w "./send/filewatcher"
	r "./send/rest"
)

func TestMain(t *testing.T) {
	cfg := c.GetConfig("./config/config.toml")
	defer func() { uPH(recover(), cfg.Global.ErrLog, true) }()
	s.Init(cfg)
	q.Init(cfg)

	done := make(chan string)
	go r.HostHTTPForPubAsync()
	go w.StartFileWatcherAsync()
	<-done
}
