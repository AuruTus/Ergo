package servepoint

import (
	"context"
	"fmt"
	"time"

	"github.com/AuruTus/Ergo/src/tools"
	wsservice "github.com/AuruTus/Ergo/src/utils/wsService"
	"github.com/sirupsen/logrus"
)

type WsClient struct {
	ctx *wsservice.WsClientContext

	WSConfig *wsservice.WsClientConfig
	Handlers []func(context.Context, any, any)
}

var _ ServerPoint = (*WsClient)(nil)

func (s *WsClient) Serve() (err error) {
	s.initWsClient()

	s.ctx, err = wsservice.NewWsClientContext(s.WSConfig)
	if err != nil {
		tools.Log.WithFields(logrus.Fields{"config": *s.WSConfig}).
			Errorf("error when init websocket context\n")
		return fmt.Errorf("webscoket context init failed: %w", err)
	}

	// retry at most 3 times
	for i := range [3]struct{}{} {
		tools.Log.Infof("try to connect the websocket server the %d time\n", i)
		if err = s.ctx.TryConnect(s.WSConfig); err != nil {
			time.Sleep(1 * time.Second)
			tools.Log.WithFields(logrus.Fields{"error": err}).Warnf("try websocket connection failed\n")
			continue
		}
		// succeed to connect with the ws server
		break
	}
	if err != nil {
		tools.Log.WithFields(logrus.Fields{"error": err}).Errorf("failed to create websocket connection\n")
		return fmt.Errorf("failed to connect with the host: %w", err)
	}

	wsservice.ServeWSClientConnection(s.ctx, s.Handlers...)

	return
}

func (s *WsClient) initWsClient() {
	// TODO
	s.WSConfig = nil
}

func NewWsClient() *WsClient {
	return &WsClient{}
}
