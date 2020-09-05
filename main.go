package main

import (
	"net"
	"os"

	"google.golang.org/grpc/reflection"

	"github.com/hashicorp/go-hclog"
	proto "github.com/leenorshn/currency/proto"
	"github.com/leenorshn/currency/server"

	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewCurrency(log)

	proto.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)
	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	gs.Serve(l)
}
