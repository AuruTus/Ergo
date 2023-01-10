package main

import (
	"os"
	"os/signal"
	"syscall"

	servepoint "github.com/AuruTus/Ergo/src/servePoint"
)

// TODO
func ServePointDone() <-chan struct{} {
	return nil
}

func main() {
	/* NOTE: os signal channel's buffer is needed */
	os_signal := make(chan os.Signal, 1)
	signal.Notify(os_signal, syscall.SIGINT, syscall.SIGTERM)

	servepoint.GlobalLogger.Infof("Good day! Ergo is at your service.\n")

	select {
	case <-os_signal:
	case <-ServePointDone():
	}
}
