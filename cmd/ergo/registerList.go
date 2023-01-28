package main

import (
	"github.com/AuruTus/Ergo/pkg/handler"
	cqhttpHandler "github.com/AuruTus/Ergo/pkg/handler/cqhttp"
	sp "github.com/AuruTus/Ergo/pkg/servePoint"
	services "github.com/AuruTus/Ergo/pkg/services"
)

var registerList = [](func() error){
	services.RegisterNamedService[handler.WSClientHandler]("cqhttp_ws_client", sp.NewWSClient, cqhttpHandler.NewWSClientHandler(), "qq robot"),
}
