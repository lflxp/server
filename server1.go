package main

import (
    "flag"

    "github.com/smallnest/rpcx/server"
    "context"
)

var (
    addr = flag.String("addr", "0.0.0.0:8972", "server address")
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
    flag.Parse()

    s := server.NewServer()
    //s.RegisterName("Arith", new(example.Arith), "")
    s.Register(new(Arith), "")
    s.Serve("tcp", *addr)
}