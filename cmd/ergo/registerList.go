package main

import (
	"github.com/AuruTus/Ergo/pkg/handler/cqhttp"
	sp "github.com/AuruTus/Ergo/pkg/servePoint"
	services "github.com/AuruTus/Ergo/pkg/services"
)

// TODO: add a code generator to get code from config files
var registerList = [](func() error){
	services.RegisterNamedService("cqhttp-ws-client", sp.NewWSClient, &cqhttp.WSClientHandler{}, "cqhttp for feifei"),
}
