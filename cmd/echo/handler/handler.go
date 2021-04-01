package handler

import (
	"context"
	"fmt"

	"github.com/jdxj/study_kubernetes/cmd/echo/proto"
)

type EchoService struct {
}

func (es *EchoService) Hello(ctx context.Context, req *proto.Ping, resp *proto.Pong) error {
	resp.Data = []byte(fmt.Sprintf("hello: %s", req.Data))
	return nil
}
