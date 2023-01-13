package wsservice

import (
	"context"
	"fmt"

	"github.com/AuruTus/Ergo/src/tools"
	ws "github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WsClientContext struct {
	context.Context
	Cancel context.CancelFunc

	Logger *logrus.Logger

	dialer *ws.Dialer
	conn   *ws.Conn
}

func NewWsClientContext(config WsClientConfig) (*WsClientContext, error) {
	ctx := new(WsClientContext)

	// init context and connection
	ctx.Context, ctx.Cancel = context.WithCancel(context.Background())

	ctx.dialer = ws.DefaultDialer
	conn, resp, err := ctx.dialer.DialContext(
		ctx,
		config.HostAddr.Network(),
		config.RequestHeader,
	)
	if err != nil {
		tools.Log.WithFields(
			map[string]any{
				"response": resp,
				"wsconfig": config,
			},
		).Errorf("websocket connection to %s failed\n", config.HostAddr.Network())
		return nil, fmt.Errorf("faild websocket handshake: %w", err)
	}
	ctx.conn = conn

	ctx.Logger = tools.NewConfiguredLogger(config.LogConfigs)

	return ctx, nil
}
