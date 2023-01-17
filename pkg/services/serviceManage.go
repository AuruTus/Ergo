package common

import (
	"os"
	"os/signal"
	"syscall"

	sp "github.com/AuruTus/Ergo/pkg/servePoint"
	"github.com/AuruTus/Ergo/tools"
)

// TODO complete servePointer register
func ServePointDone() <-chan struct{} {
	return nil
}

func RunServices() {
	/* NOTE: os signal channel's buffer is needed */
	os_signal := make(chan os.Signal, 1)
	signal.Notify(os_signal, syscall.SIGINT, syscall.SIGTERM)

	tools.Log.Infof("Good day! Ergo is at your service.\n")

	// TODO add entrance from service manager
	done := make(chan struct{}, 1)

	s, _ := sp.NewWsClient()
	s.Register()
	if err := s.Serve(); err != nil {
		tools.Log.Errorf("%v\n", err)
		done <- struct{}{}
	}

	select {
	case <-os_signal:
		s.Close()
	case <-ServePointDone():
	case <-done:
	}
}
