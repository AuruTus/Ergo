package wsservice

import (
	"context"
	"fmt"
	"time"

	"github.com/AuruTus/Ergo/src/tools"
	ws "github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WsClientContext struct {
	context.Context
	Cancel context.CancelFunc

	Logger *logrus.Logger

	dialer *ws.Dialer
	Conn   *ws.Conn

	heartBeatInterval time.Duration
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
	ctx.Conn = conn
	return nil
}

func ServeWSClientConnection(ctx *WsClientContext, handlers ...func(context.Context, any, any)) {
	defer ctx.Conn.Close()

	heartBeatTicker := time.NewTicker(ctx.heartBeatInterval)
	defer heartBeatTicker.Stop()

	// todo
	receiveHandler := func() {
		for {
			_, msg, err := ctx.Conn.ReadMessage()
			if err != nil {
				ctx.Logger.WithFields(logrus.Fields{"err": err}).Errorf("error when received message\n")
				return
			}
			ctx.Logger.Infof("msg: %s\n", msg)
		}
	}
	go receiveHandler()

	for {
		select {
		// todo read cqhttp api
		case <-heartBeatTicker.C:
			ctx.Conn.WriteMessage(ws.TextMessage, []byte("Aloha"))
		case <-ctx.Done():
			handlers[0](ctx, struct{}{}, struct{}{})
		}
	}
}
