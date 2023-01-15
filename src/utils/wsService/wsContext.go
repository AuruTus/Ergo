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

/* NewWsClientContext will try to create context with relevant websocket connection */
func NewWsClientContext(config *WsClientConfig) (*WsClientContext, error) {
	ctx := new(WsClientContext)

	// init context and connection
	ctx.Context, ctx.Cancel = context.WithCancel(context.Background())

	ctx.dialer = ws.DefaultDialer

	ctx.Logger = tools.NewConfiguredLogger(config.LogConfigs)

	return ctx, nil
}

func (ctx *WsClientContext) TryConnect(config *WsClientConfig) error {
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
		return fmt.Errorf("faild websocket handshake: %w", err)
	}
	ctx.conn = conn
	return nil
}

func ServeWSClientConnection(ctx *WsClientContext, handlers ...func(context.Context, any, any)) {

}
