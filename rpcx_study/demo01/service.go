package main

import "context"

// -------------------------------------------
// @file          : service.go
// @author        : binshow
// @time          : 2022/6/12 4:55 PM
// @description   :
// -------------------------------------------

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}


// 实现一个 Arith Service
type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

