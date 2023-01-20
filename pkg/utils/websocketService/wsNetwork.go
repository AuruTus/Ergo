package websocketService

import (
	"fmt"
	"net"
	"sync/atomic"

	"github.com/AuruTus/Ergo/pkg/handler"
	"github.com/AuruTus/Ergo/tools"
	ws "github.com/gorilla/websocket"
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

const (
	DEFAULT_WRITER_BUFFER_SIZE = 10

	DEFAULT_RETRY_TIME = 3
)

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
	atomic.StoreUint32(&ctx.active, 1)
	tools.Log.Infof("succeed connecting to ws server end %s")

	return nil
}

func ServeWSClientConnection(ctx *WsClientContext, handler handler.Handler) {
	// sync between reader and writer
	info := make(chan interface{}, ctx.writerBufferSize)
	wDone := make(chan struct{})

	defer func() {
		close(info)
		<-wDone
		ctx.Conn.Close()
		atomic.StoreUint32(&ctx.active, 0)
	}()

	// writer goroutine
	go func() {
		// ensure the last ws writer is closed
		defer func() { close(wDone) }()

		for b := range info {
			select {
			case <-ctx.Done():
				return
			default:
				tools.SafeRun(func() {
					ctx.Logger.Infof("writer get info %v\n", len(b.([]byte)))
				})
			}
		}
	}()

	// reader main goroutine
	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, msg, err := ctx.Conn.ReadMessage()
			switch {
			case err != nil:
				ctx.Logger.WithFields(logrus.Fields{"err": err}).Errorf("error when received message\n")
			default:
				ctx.Logger.Infof("msg: %s\n", msg)
				info <- msg
			}
		}
	}
}

func TrySendCloseClosure(ctx *WsClientContext) error {
	return ctx.Conn.WriteMessage(
		ws.CloseMessage,
		ws.FormatCloseMessage(ws.CloseNormalClosure, ""),
	)
}
