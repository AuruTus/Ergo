package servepoint

import "github.com/AuruTus/Ergo/pkg/handler"

/*
 ServePoint is an entrance which acts as a service integration point.
 It's implemented as the client or the server depending on protocals and scinarios;
*/
type ServePoint interface {
	Serve() error
	IsAlive() bool
	Close() error
}

type ServerPointGenerator (func(configKey string, h handler.Handler) (ServePoint, error))
