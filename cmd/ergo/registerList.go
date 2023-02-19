package main

import (
	"github.com/AuruTus/Ergo/pkg/handler"
	cqhttpHandler "github.com/AuruTus/Ergo/pkg/handler/cqhttp"
	engines "github.com/AuruTus/Ergo/pkg/serveEngines"
	sp "github.com/AuruTus/Ergo/pkg/servePoint"
)

var registerList = [](func() error){
	engines.RegisterNamedEngine[handler.WSClientHandler]("cqhttp_ws_client", sp.NewWSClient, cqhttpHandler.NewWSClientHandler(), "qq robot"),
}
