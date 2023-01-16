package websocketService

import (
	"fmt"
	"net"

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
		return nil, fmt.Errorf("resolve TCP addr failed: %w", err)
	}
	addr.Api = api
	return
}
