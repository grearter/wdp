package food

import (
	"fmt"
	"github.com/grearter/wz_rpc_go"
	"net"
	"os"
)

const (
	ParaError int = iota + 1000
	DupError
	SrvError
)

var (
	FoodDaoPool wz_rpc_go.Pool
)

func init() {
	addr := "127.0.0.1:4031"

	p, err := wz_rpc_go.NewChanPool(10, func() (conn *wz_rpc_go.Conn, e error) {
		c, e := net.Dial("tcp", addr)
		if e != nil {
			return
		}

		conn = &wz_rpc_go.Conn{
			Conn:  c,
			Codec: wz_rpc_go.NewJsonCodec(c),
		}

		return
	})

	if err != nil {
		fmt.Printf("create pool err: %v, addr: %v\n", err, addr)
		os.Exit(-1)
		return
	}

	FoodDaoPool = p
}
