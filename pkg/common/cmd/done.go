package cmd_helper

import (
	"os"
	"os/signal"
	"syscall"
)

func Done() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return c
}
