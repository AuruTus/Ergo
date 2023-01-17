package websocketService

import (
	"context"
	"sync"

	"github.com/AuruTus/Ergo/tools"
	"github.com/google/uuid"
	ws "github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WsClientContext struct {
	// CID is the uuid for each wscontext
	CID string

	/*
	 websocket connection related fields
	*/
	// embeded context
	context.Context
	Cancel context.CancelFunc
	// websocket connection flieds
	dialer        *ws.Dialer
	Conn          *ws.Conn
	closeConnOnce sync.Once

	// context level logger
	Logger *logrus.Logger
}

/* NewWsClientContext will try to create context with relevant websocket connection */
func NewWsClientContext(config *WsClientConfig) (*WsClientContext, error) {
	ctx := new(WsClientContext)

	ctx.CID = uuid.New().String()

	_ctx, cancel := context.WithCancel(context.Background())
	ctx.Context = _ctx
	ctx.Cancel = func() {
		if ctx.Conn != nil {
			ctx.closeConnOnce.Do(func() { ctx.Conn.Close() })
		}
		cancel()
	}
	ctx.dialer = ws.DefaultDialer
	ctx.Logger = tools.NewConfiguredLogger(config.LogConfigs)

	return ctx, nil
}
