package servepoint

import (
	"context"
	"fmt"
	"time"

	"github.com/AuruTus/Ergo/src/tools"
	wsservice "github.com/AuruTus/Ergo/src/utils/websocketService"
	"github.com/sirupsen/logrus"
)

type WsClient struct {
	ctx *wsservice.WsClientContext

	WSConfig *wsservice.WsClientConfig
	Handlers []func(context.Context, any, any)
}

var _ ServePoint = (*WsClient)(nil)

/* TODO: add service register logic */
func (s *WsClient) Register() (err error) {
	return nil
}

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

	wsservice.ServeWSClientConnection(s.ctx, s.Handlers...)

	return
}

/* TODO: add close function details */
func (s *WsClient) Close() error {
	s.ctx.Cancel()

	return nil
}

func (s *WsClient) initWsClient() (err error) {
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

func NewWsClient() (s *WsClient, err error) {
	s = &WsClient{}
	if err = s.initWsClient(); err != nil {
		return nil, fmt.Errorf("init ws client: %w", err)
	}
	return
}
