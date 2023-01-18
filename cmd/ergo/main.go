package main

import (
	"os"
	"os/signal"
	"syscall"

	services "github.com/AuruTus/Ergo/pkg/services"
	"github.com/AuruTus/Ergo/tools"
)

func RunServices() {
	/* NOTE: os signal channel's buffer is needed */
	os_signal := make(chan os.Signal, 1)
	signal.Notify(os_signal, syscall.SIGINT, syscall.SIGTERM)

	tools.Log.Infof("Good day! Ergo is at your service.\n")

	// TODO add entrance from service manager
	done := make(chan struct{}, 1)

	// s, _ := sp.NewWsClient()
	// if err := s.Serve(); err != nil {
	// 	tools.Log.Errorf("%v\n", err)
	// 	done <- struct{}{}
	// }
	registerServices()
	services.StartServicesAll()

	select {
	case <-os_signal:
		services.CloseServicesAll()
	// case <-services.CloseServicesAll():
	case <-done:
	}
}

func main() {
	RunServices()
}
