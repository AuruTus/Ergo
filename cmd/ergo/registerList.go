package main

import (
	cqhttpHandler "github.com/AuruTus/Ergo/pkg/handler/cqhttp"
	sp "github.com/AuruTus/Ergo/pkg/servePoint"
	services "github.com/AuruTus/Ergo/pkg/services"
)

// TODO: add a code generator to get code from config files
var registerList = [](func() error){
	services.RegisterNamedService("cqhttp-ws-client", sp.NewWSClient, &cqhttpHandler.CQHTTPHandler{}, "cqhttp for feifei"),
}
