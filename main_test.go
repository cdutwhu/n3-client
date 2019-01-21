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
	defer func() { uPH(recover(), "./log.txt", true) }()

	cfg := &c.Config{}
	cfg.Load("./config/config.toml")
	s.Init(cfg)
	q.Init(cfg)

	done := make(chan string)
	go r.HostHTTPForPubAsync()
	go w.StartFileWatcherAsync()
	<-done
}
