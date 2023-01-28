package servepoint

import (
	"fmt"
	"time"

	"github.com/AuruTus/Ergo/pkg/handler"
	"github.com/AuruTus/Ergo/pkg/utils/logger"
	wsservice "github.com/AuruTus/Ergo/pkg/utils/websocketService"
	"github.com/sirupsen/logrus"
)

type WSClient struct {
	ctx *wsservice.WSClientContext
	h   handler.Handler

	WSConfig *wsservice.WSClientConfig
}

var _ ServePoint = (*WSClient)(nil)

func (s *WSClient) Serve() (err error) {
	// retry at most 3 times
	for i := range [3]struct{}{} {
		logger.Infof("try to connect with the address %s the %d time\n", s.WSConfig.HostAddr.String(), i+1)
		if err = wsservice.TryConnect(s.ctx, s.WSConfig); err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		// succeed to connect with the ws server
		break
	}
	if err != nil {
		logger.WithFields(logrus.Fields{"error": err}).Errorf("failed to create websocket connection\n")
		return fmt.Errorf("connect the host: %w", err)
	}

	wsservice.ServeWSClientConnection(s.ctx, s.h)

	return
}

func (s *WSClient) IsAlive() bool {
	return s != nil && s.ctx.IsActive()
}

func (s *WSClient) Close() (err error) {
	if !s.IsAlive() {
		return fmt.Errorf("dead service")
	}
	// Just send close handshake control message
	// The close of ws connection will be really completed in the reader main goroutine
	for i := range [3]struct{}{} {
		logger.Infof("say goodbye with the websocket server the %d time\n", i+1)
		if err = wsservice.TrySendCloseClosure(s.ctx); err != nil {
			time.Sleep(1 * time.Second)
			logger.Warnf("client failed to send close closure")
			continue
		}
		break
	}
	if err != nil {
		logger.WithFields(logrus.Fields{"error": err}).Errorf("failed to send close closure\n")
		err = fmt.Errorf("send close closure: %w", err)
	}
	s.ctx.Cancel()
	return
}

func (s *WSClient) initWSClient(configKey string) (err error) {
	// TODO load configuration from config file
	if s.WSConfig, err = wsservice.NewWSClientConfig(); err != nil {
		return fmt.Errorf("new ws client config: %w", err)
	}

	s.ctx, err = wsservice.NewWSClientContext(s.WSConfig)
	if err != nil {
		logger.WithFields(logrus.Fields{"config": *s.WSConfig}).
			Errorf("error when init websocket context\n")
		return fmt.Errorf("init webscoket context: %w", err)
	}

	return
}

func NewWSClient(configKey string, h handler.WSClientHandler) (ServePoint, error) {
	s := &WSClient{}
	if err := s.initWSClient(configKey); err != nil {
		return nil, fmt.Errorf("init ws client: %w", err)
	}
	s.h = h
	return s, nil
}
