package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	engines "github.com/AuruTus/Ergo/pkg/serveEngines"
	"github.com/AuruTus/Ergo/pkg/utils/logger"
)

func RunEngines() {
	/* NOTE: os signal channel's buffer is needed */
	os_signal := make(chan os.Signal, 1)
	signal.Notify(
		os_signal,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT,
	)
	flag.Parse()

	logger.Infof("Good day! Ergo is at your service.\n")

	// TODO add entrance from service manager
	done := make(chan struct{}, 1)

	engines.RegisterEnginesAll(registerList)
	engines.StartEnginesAll()

	select {
	case <-os_signal:
		engines.CloseEnginesAll()
	case <-done:
	}

	time.Sleep(3 * time.Second)
}

func main() {
	RunEngines()
}
