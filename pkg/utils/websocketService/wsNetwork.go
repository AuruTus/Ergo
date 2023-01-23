package websocketService

import (
	"errors"
	"fmt"
	"net"

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

func TryConnect(ctx *WSClientContext, config *WSClientConfig) error {
	if ctx.conn != nil {
		return fmt.Errorf("connection for context %s is already established", ctx.CID)
	}
	conn, resp, err := ctx.dialer.DialContext(
		ctx,
		config.HostAddr.String(),
		config.RequestHeader,
	)
	if err != nil {
		ctx.Logger.WithFields(
			map[string]any{
				"response":      resp,
				"requestHeader": config.RequestHeader,
			},
		).Errorf("failed to dial the address %s\n", config.HostAddr.String())
		return fmt.Errorf("dial websocket server: %w", err)
	}
	ctx.conn = conn
	ctx.activate()
	ctx.Logger.Infof("succeed connecting to websocket serverend %s\n", ctx.conn.RemoteAddr().String())

	return nil
}

func ServeWSClientConnection(ctx *WSClientContext, h handler.WSClientHandler) {
	// sync between reader and writer
	msgBuffer := make(chan []byte, ctx.writerBufferSize)
	wDone := make(chan struct{})

	defer func() {
		close(msgBuffer)
		<-wDone
		ctx.closeConn()
		ctx.deactivate()
		ctx.Logger.Infof("connection with host %s closed\n", ctx.conn.RemoteAddr().String())
		ctx.Logger.Infof("context %s is deactive, ws client is down", ctx.CID)
	}()

	// writer goroutine
	go func() {
		// ensure the last ws writer is closed
		defer func() { close(wDone) }()

		for msg := range msgBuffer {
			select {
			case <-ctx.Done():
				return
			default:
				tools.SafeRun(func() {
					// ctx.Logger.Infof("writer gets %d info: %s\n", len(msg), msg)
					h.HandleWrite(ctx.conn, msg)
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
			msg, err := h.HandleRead(ctx.conn)
			// todo complete error check
			switch {
			case errors.Is(err, handler.ErrWSControlMsg):
			case errors.Is(err, handler.ErrWSResponseMsg):
				ctx.Logger.Infof("handler get response: %s\n", msg)
			case err != nil:
				ctx.Logger.Infof("websocket client gets err: %v\n", err)
				return
			default:
				ctx.Logger.Infof("reader sends %d info: %s\n", len(msg), msg)
				msgBuffer <- msg
			}
		}
	}
}

func TrySendCloseClosure(ctx *WSClientContext) error {
	return ctx.conn.WriteMessage(
		ws.CloseMessage,
		ws.FormatCloseMessage(ws.CloseNormalClosure, ""),
	)
}
