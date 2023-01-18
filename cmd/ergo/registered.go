package main

import (
	sp "github.com/AuruTus/Ergo/pkg/servePoint"
	services "github.com/AuruTus/Ergo/pkg/services"
	"github.com/AuruTus/Ergo/tools"
	"github.com/sirupsen/logrus"
)

// TODO: add a code generator to get code from config files
func registerServices() {
	registerFuncs := [](func() error){
		services.RegisterNamedService("cqhttp-ws-client", sp.NewWsClient, "cqhttp for feifei"),
	}

	for i, f := range registerFuncs {
		if err := f(); err != nil {
			tools.Log.WithFields(logrus.Fields{"error": err}).Errorf("start service %v failed\n", i)
		}
	}
}
