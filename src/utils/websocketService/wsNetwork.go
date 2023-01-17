package websocketService

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/AuruTus/Ergo/src/tools"
	"github.com/sirupsen/logrus"
)

type WSAddr struct {
	net.Addr
	Api string
}

var _ net.Addr = (*WSAddr)(nil)

func (addr *WSAddr) Network() string { return "ws" }

func (addr *WSAddr) String() string {
	if addr == nil || addr.Addr == nil || len(addr.Api) == 0 {
		return "<nil>"
	}
	return fmt.Sprintf("%s://%s%s", addr.Network(), addr.Addr.String(), addr.Api)
}

func ResolveWSAddrFromSocket(socket, api string) (addr *WSAddr, err error) {
	addr = &WSAddr{}
	addr.Addr, err = net.ResolveTCPAddr("tcp", socket)
	if err != nil {
		tools.Log.WithFields(logrus.Fields{"socket": socket}).Errorf("fail to resolve tcp addr\n")
		return nil, fmt.Errorf("resolve TCP addr: %w", err)
	}
	addr.Api = api
	return
}

/*
	WS Connection
*/
func TryConnect(ctx *WsClientContext, config *WsClientConfig) error {
	conn, resp, err := ctx.dialer.DialContext(
		ctx,
		config.HostAddr.String(),
		config.RequestHeader,
	)
	if err != nil {
		if !errors.Is(err, io.ErrUnexpectedEOF) {
			tools.Log.Errorf("dial err: %v\n", err)
		}
		tools.Log.WithFields(
			map[string]any{
				"response":      resp,
				"requestHeader": config.RequestHeader,
			},
		).Errorf("failed to dial the websocket server %s\n", config.HostAddr.String())
		return fmt.Errorf("dial websocket server: %w", err)
	}
	ctx.Conn = conn
	return nil
}

func ServeWSClientConnection(ctx *WsClientContext, handlers ...func(context.Context, any, any)) {
	go func() {
		for {
			// TODO add support for converting []byte to json
			_, msg, err := ctx.Conn.ReadMessage()
			if err != nil {
				ctx.Logger.WithFields(logrus.Fields{"err": err}).Errorf("error when received message\n")
				return
			}
			ctx.Logger.Infof("msg: %s\n", msg)
		}
	}()

	go func() {
		heartBeatTicker := time.NewTicker(ctx.heartBeatInterval)
		defer heartBeatTicker.Stop()

		for {
			select {
			// todo read cqhttp api
			case <-heartBeatTicker.C:
				// ctx.Conn.WriteMessage(ws.TextMessage, []byte("Aloha"))
			case <-ctx.Done():
				handlers[0](ctx, struct{}{}, struct{}{})
			}
		}
	}()
}
