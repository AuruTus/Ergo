package websocketService

import (
	"context"
	"sync"
	"sync/atomic"

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
	cancel context.CancelFunc
	// websocket connection flieds
	dialer        *ws.Dialer
	Conn          *ws.Conn
	closeConnOnce sync.Once
	// used as atomic
	active uint32

	// context level logger
	Logger *logrus.Logger

	writerBufferSize int
}

func (ctx *WsClientContext) Cancel() {
	if ctx.Conn != nil {
		// Just send close handshake control message
		// The close of ws connection will be really completed in the reader main goroutine
		ctx.closeConnOnce.Do(func() {
			TrySendCloseClosure(ctx)
		})
	}
	if ctx.cancel != nil {
		ctx.cancel()
	}
}

func (ctx *WsClientContext) IsActive() bool {
	return ctx != nil && atomic.LoadUint32(&ctx.active) != 0
}

/* NewWsClientContext will try to create context with relevant websocket connection */
func NewWsClientContext(config *WsClientConfig) (*WsClientContext, error) {
	ctx := new(WsClientContext)

	ctx.CID = uuid.New().String()

	ctx.Context, ctx.cancel = context.WithCancel(context.Background())
	ctx.dialer = ws.DefaultDialer
	ctx.Logger = tools.NewConfiguredLogger(config.LogConfigs)

	ctx.writerBufferSize = DEFAULT_WRITER_BUFFER_SIZE
	if config.WriterBufferSize > 0 {
		ctx.writerBufferSize = config.WriterBufferSize
	}

	return ctx, nil
}
