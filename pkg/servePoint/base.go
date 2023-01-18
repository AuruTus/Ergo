package servepoint

/*
 ServePoint is an entrance which acts as a service integration point.
 It's implemented as the client or the server depending on protocals and scinarios;
*/
type ServePoint interface {
	Serve() error
	Close() error
}

type ServerPointGenerator (func() (ServePoint, error))
