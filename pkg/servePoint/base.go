package servepoint

import "sync"

/*
 ServePoint is an entrance which acts as a service integration point.
 It's implemented as the client or the server depending on protocals and scinarios;
*/
type ServePoint interface {
	Register() error
	Serve() error
	Close() error
}

type ServiceRecords struct {
	Records map[string]ServePoint
}

var (
	serviceRecords *ServiceRecords
	initListOnce   sync.Once
)

func NewServiceRecords() *ServiceRecords {
	return &ServiceRecords{
		Records: make(map[string]ServePoint),
	}
}

func InitGlobalServiceRecords() {
	if serviceRecords == nil {
		initListOnce.Do(func() { serviceRecords = NewServiceRecords() })
	}
}
