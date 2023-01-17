package websocketService

import (
	"context"
	"sync"
	"time"

	"github.com/AuruTus/Ergo/src/tools"
	ws "github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WsClientContext struct {
	context.Context
	Cancel context.CancelFunc

	Logger *logrus.Logger

	dialer         *ws.Dialer
	Conn           *ws.Conn
	cancelConnOnce sync.Once

	heartBeatInterval time.Duration
}

/* NewWsClientContext will try to create context with relevant websocket connection */
func NewWsClientContext(config *WsClientConfig) (*WsClientContext, error) {
	ctx := new(WsClientContext)

	ctx.Context, ctx.Cancel = context.WithCancel(context.Background())
	ctx.Cancel = func() {
		if ctx.Conn != nil {
			ctx.cancelConnOnce.Do(func() { ctx.Conn.Close() })
		}
		ctx.Cancel()
	}
	ctx.dialer = ws.DefaultDialer
	ctx.Logger = tools.NewConfiguredLogger(config.LogConfigs)

	ctx.heartBeatInterval = 1 * time.Second

	return ctx, nil
}
