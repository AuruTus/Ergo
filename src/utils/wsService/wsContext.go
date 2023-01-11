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

	/* init context and connection */
	ctx.Context, ctx.Cancel = context.WithCancel(context.Background())

	ctx.dialer = ws.DefaultDialer
	conn, resp, err := ctx.dialer.DialContext(
		ctx,
		config.HostAddr.Network(),
		config.RequestHeader,
	)
	ctx.conn = conn
	if err != nil {
		tools.Log.WithFields(
			map[string]any{
				"response": resp,
				"wsconfig": config,
			},
		).Errorf("websocket connection to %s failed\n", config.HostAddr.Network())
		return nil, fmt.Errorf("faild websocket handshake: %w", err)
	}

	// TODO using tools to wrap it
	ctx.Logger = logrus.New()

	return ctx, nil
}
