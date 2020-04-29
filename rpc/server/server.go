package server

import (
	"errors"
	"log"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multipy(args *Args, reply *int) error {
	log.Fatal("接入方调用Multipy RPC方法 ")
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	log.Fatal("接入方调用Divide RPC方法 ")
	if args.B == 0 {
		return errors.New("Divide is zero ")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
