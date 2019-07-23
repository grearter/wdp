package main

import (
	"dao/food"
	"encoding/json"
	"github.com/grearter/wz_rpc_go"
	"io"
	"net"
)

var (
	FoodDao = &food.Food{}
)

func main() {
	addr := "127.0.0.1:4031"
	server := wz_rpc_go.NewServer(addr, func(conn net.Conn) wz_rpc_go.Codec {
		return &wz_rpc_go.JsonCodec{
			Enc: json.NewEncoder(conn.(io.Writer)),
			Dec: json.NewDecoder(conn.(io.Reader)),
		}
	})

	server.Register(FoodDao)

	server.Serve()
}
