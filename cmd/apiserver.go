package main

import (
	"my-kubesphere/pkg/config"
)

type ServerRunOptions struct {
	*config.Config
	DebugMode bool
}

func main() {
	s := &ServerRunOptions{
		Config:    config.New(),
		DebugMode: false,
	}
}
