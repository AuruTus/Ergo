package services

import (
	"fmt"
	"sync"

	"github.com/AuruTus/Ergo/pkg/handler"
	sp "github.com/AuruTus/Ergo/pkg/servePoint"
	"github.com/AuruTus/Ergo/pkg/utils/logger"
	"github.com/AuruTus/Ergo/tools"
	"github.com/sirupsen/logrus"
)

func init() {
	// init services singleton
	initGlobalServices()
}

type service struct {
	servePoint sp.ServePoint

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

func RegisterNamedService[H handler.Handler](name string, g sp.ServerPointGenerator[H], h H, desc string) func() error {
	// lazy calling
	return func() error {
		servePoint, err := g(name, h)
		if err != nil {
			return err
		}

		if _, ok := services.get(name); ok {
			return fmt.Errorf("service %s already registered", name)
		}

		sv := &service{
			servePoint: servePoint,
			Desc:       desc,
		}
		if err := services.set(name, sv); err != nil {
			return err
		}
		return nil
	}
}

func RegisterServicesAll(registerList [](func() error)) {
	for i, f := range registerList {
		if err := f(); err != nil {
			logger.WithFields(logrus.Fields{"error": err}).Errorf("start service %v failed\n", i)
		}
	}
}

func StartServicesAll() {
	for k, s := range services.kv {
		go func(key string, s *service) {
			if err := tools.WithRecover(func() {
				if err := s.servePoint.Serve(); err != nil {
					panic(fmt.Errorf("service %s with %T serving: %w", key, s.servePoint, err))
				}
			}); err != nil {
				logger.Infof("%v", err)
			}
		}(k, s)
	}
}

func CloseServicesAll() {
	for k, s := range services.kv {
		if err := tools.WithRecover(func() {
			s.servePoint.Close()
		}); err != nil {
			logger.Infof("unable to close service %v: %v", k, err)
		}
	}
}
