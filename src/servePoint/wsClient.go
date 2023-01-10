package servepoint

import (
	"context"
)

type WsClient struct {
	ctx context.Context
}

var _ ServerPoint = (*WsClient)(nil)

func (s *WsClient) Serve() {

}

func (s *WsClient) initNewWsClient() {
	s.ctx = context.Background()
}

func NewWsClient() *WsClient {
	return &WsClient{}
}
