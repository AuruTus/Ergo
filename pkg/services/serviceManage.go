package services

import (
	"fmt"
	"sync"

	sp "github.com/AuruTus/Ergo/pkg/servePoint"
	"github.com/AuruTus/Ergo/tools"
)

func init() {
	// init services singleton
	initGlobalServices()
}

type service struct {
	servePoint sp.ServePoint
	active     uint8

	Desc string
}

type Services struct {
	kv map[string]*service
}

func (s *Services) set(key string, val *service) error {
	if s == nil || s.kv == nil {
		return fmt.Errorf("invlalid Services %v", s)
	}
	s.kv[key] = val
	return nil
}

func (s *Services) get(key string) (val *service, ok bool) {
	if s == nil || s.kv == nil {
		return nil, false
	}
	val, ok = s.kv[key]
	return
}

var (
	services     *Services
	initListOnce sync.Once
)

func newServices() *Services {
	return &Services{
		kv: make(map[string]*service),
	}
}

func initGlobalServices() {
	if services == nil {
		initListOnce.Do(func() { services = newServices() })
	}
}

func RegisterNamedService(name string, g sp.ServerPointGenerator, desc string) func() error {
	// lazy calling
	return func() error {
		servePoint, err := g()
		if err != nil {
			return err
		}

		if _, ok := services.get(name); ok {
			return fmt.Errorf("service %s already registered", name)
		}

		sv := &service{
			servePoint: servePoint,
			active:     1,
			Desc:       desc,
		}
		if err := services.set(name, sv); err != nil {
			return err
		}
		return nil
	}
}

func StartServicesAll() {
	for _, s := range services.kv {
		tools.Go(func() {
			if err := s.servePoint.Serve(); err != nil {
				panic(fmt.Errorf("%T serving: %w", s.servePoint, err))
			}
		})
	}
}

func CloseServicesAll() {
	for _, s := range services.kv {
		s.servePoint.Close()
	}
}
