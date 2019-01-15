package config

import "github.com/nsip/n3-messages/n3grpc"

var (
	// Cfg : global config file
	Cfg = &Config{}

	// N3pub :
	N3pub *n3grpc.Publisher
)
