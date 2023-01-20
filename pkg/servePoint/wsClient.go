package servepoint

import (
	"fmt"
	"time"

	"github.com/AuruTus/Ergo/pkg/handler"
	wsservice "github.com/AuruTus/Ergo/pkg/utils/websocketService"
	"github.com/AuruTus/Ergo/tools"
	"github.com/sirupsen/logrus"
)

type WsClient struct {
	ctx *wsservice.WsClientContext
	h   handler.Handler

	WSConfig *wsservice.WsClientConfig
}

var _ ServePoint = (*WsClient)(nil)

func (s *WsClient) Serve() (err error) {
	// retry at most 3 times
	for i := range [3]struct{}{} {
		tools.Log.Infof("try to connect the websocket server the %d time\n", i+1)
		if err = wsservice.TryConnect(s.ctx, s.WSConfig); err != nil {
			time.Sleep(1 * time.Second)
			tools.Log.Warnf("client failed to create websocket connection")
			continue
		}
		// succeed to connect with the ws server
		break
	}
	if err != nil {
		tools.Log.WithFields(logrus.Fields{"error": err}).Errorf("failed to create websocket connection\n")
		return fmt.Errorf("connect the host: %w", err)
	}

	wsservice.ServeWSClientConnection(s.ctx, s.h)

	return
}

func (s *WsClient) IsAlive() bool {
	return s != nil && s.ctx.IsActive()
}

/* TODO: add close function details */
func (s *WsClient) Close() error {
	s.ctx.Cancel()

	return nil
}

func (s *WsClient) initWsClient(configKey string) (err error) {
	// TODO load configuration from config file
	if s.WSConfig, err = wsservice.NewWSClientConfig(); err != nil {
		return fmt.Errorf("new ws client config: %w", err)
	}

	s.ctx, err = wsservice.NewWsClientContext(s.WSConfig)
	if err != nil {
		tools.Log.WithFields(logrus.Fields{"config": *s.WSConfig}).
			Errorf("error when init websocket context\n")
		return fmt.Errorf("init webscoket context: %w", err)
	}

	return
}

func NewWsClient(configKey string, h handler.Handler) (ServePoint, error) {
	s := &WsClient{}
	if err := s.initWsClient(configKey); err != nil {
		return nil, fmt.Errorf("init ws client: %w", err)
	}
	s.h = h
	return s, nil
}

var _ ServerPointGenerator = NewWsClient
