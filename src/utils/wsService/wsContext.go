package wsservice

import (
	"context"
	"fmt"

	servepoint "github.com/AuruTus/Ergo/src/servePoint"
	ws "github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WsContext struct {
	context.Context
	Cancel context.CancelFunc

	Logger *logrus.Logger

	dialer *ws.Dialer
	conn   *ws.Conn
}

func NewWsContext(config WsClientConfig) (*WsContext, error) {
	ctx := new(WsContext)

	/* init context and connection */
	ctx.Context, ctx.Cancel = context.WithCancel(context.Background())
	ctx.dialer = &ws.Dialer{}

	conn, resp, err := ctx.dialer.DialContext(
		ctx,
		config.HostAddr.Network(),
		config.RequestHeader,
	)
	ctx.conn = conn
	if err != nil {
		servepoint.GlobalLogger.WithFields(
			map[string]any{
				"response": resp,
				"wsconfig": config,
			},
		).Errorf("websocket connection to %s failed\n", config.HostAddr.Network())
		return nil, fmt.Errorf("faild websocket handshake: %w", err)
	}

	ctx.Logger = logrus.New()

	return ctx, nil
}
