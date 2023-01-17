package websocketService

import (
	"context"
	"fmt"
	"net"

	"github.com/AuruTus/Ergo/tools"
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
	if ctx.Conn != nil {
		return fmt.Errorf("connection for context %s is already established", ctx.CID)
	}
	conn, resp, err := ctx.dialer.DialContext(
		ctx,
		config.HostAddr.String(),
		config.RequestHeader,
	)
	if err != nil {
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
	// goroutine for send request
	go func() {
		for {
			select {
			case <-ctx.Done():
				// todo read cqhttp api
			}
		}
	}()

	// awaiting
	for {
		select {
		case <-ctx.Done():
		default:
			// TODO add support for converting []byte to json
			_, msg, err := ctx.Conn.ReadMessage()
			if err != nil {
				ctx.Logger.WithFields(logrus.Fields{"err": err}).Errorf("error when received message\n")
				return
			}
			ctx.Logger.Infof("msg: %s\n", msg)
		}
	}
}
