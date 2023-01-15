package servepoint

/*
	ServerPoint is an entrance which acts as a service integration point.
	It's implemented as the client or the server depending on protocals and scinarios;
*/
type ServerPoint interface {
	Register() error
	Serve() error
}
