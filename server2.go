package main

import (
    "flag"

    "github.com/smallnest/rpcx/server"
    "github.com/smallnest/rpcx/protocol"
    "context"
    "errors"
)

var (
    addr2 = flag.String("addr", "0.0.0.0:8972", "server address")
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
    s.RegisterName("Arith", new(Arith), "")
    // s.Register(new(Arith), "")
    // s.Serve("tcp", *addr)
    s.AuthFunc = auth
    s.Serve("reuseport",*addr2)
}

func auth(ctx context.Context, req *protocol.Message, token string) error {
    
        if token == "bearer tGzv3JOkF0XG5Qx2TlKWIA" {
            return nil
        }
    
        return errors.New("invalid token")
    }