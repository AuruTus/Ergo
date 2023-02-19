package engines

import (
	"fmt"
	"sync"

	"github.com/AuruTus/Ergo/pkg/handler"
	sp "github.com/AuruTus/Ergo/pkg/servePoint"
	"github.com/AuruTus/Ergo/pkg/utils"
	"github.com/AuruTus/Ergo/pkg/utils/logger"
	"github.com/sirupsen/logrus"
)

func init() {
	// init engines singleton
	initGlobalEngines()
}

type engine struct {
	servePoint sp.ServePoint

	Desc string
}

type Engines struct {
	kv map[string]*engine
}

func (s *Engines) set(key string, val *engine) error {
	if s == nil || s.kv == nil {
		return fmt.Errorf("invlalid Engines %v", s)
	}
	s.kv[key] = val
	return nil
}

func (s *Engines) get(key string) (val *engine, ok bool) {
	if s == nil || s.kv == nil {
		return nil, false
	}
	val, ok = s.kv[key]
	return
}

var (
	engines      *Engines
	initListOnce sync.Once
)

func newEngines() *Engines {
	return &Engines{
		kv: make(map[string]*engine),
	}
}

func initGlobalEngines() {
	if engines == nil {
		initListOnce.Do(func() { engines = newEngines() })
	}
}

func RegisterNamedEngine[H handler.Handler](name string, g sp.ServerPointGenerator[H], h H, desc string) func() error {
	// lazy calling
	return func() error {
		servePoint, err := g(name, h)
		if err != nil {
			return err
		}

		if _, ok := engines.get(name); ok {
			return fmt.Errorf("engine %s already registered", name)
		}

		sv := &engine{
			servePoint: servePoint,
			Desc:       desc,
		}
		if err := engines.set(name, sv); err != nil {
			return err
		}
		return nil
	}
}

func RegisterEnginesAll(registerList [](func() error)) {
	for i, f := range registerList {
		if err := f(); err != nil {
			logger.WithFields(logrus.Fields{"error": err}).Errorf("start engine %v failed\n", i)
		}
	}
}

func StartEnginesAll() {
	for k, s := range engines.kv {
		go func(key string, s *engine) {
			if err := utils.WithRecover(func() {
				if err := s.servePoint.Serve(); err != nil {
					panic(fmt.Errorf("engine %s with %T serving: %w", key, s.servePoint, err))
				}
			}); err != nil {
				logger.Infof("%v", err)
			}
		}(k, s)
	}
}

func CloseEnginesAll() {
	for k, s := range engines.kv {
		if err := utils.WithRecover(func() {
			s.servePoint.Close()
		}); err != nil {
			logger.Infof("unable to close engine %v: %v", k, err)
		}
	}
}
