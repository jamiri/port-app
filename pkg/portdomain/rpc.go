package portdomain

import (
	"context"
	protocol_message_store "github.com/jamiri/port-app/protocol/message/store"
	"google.golang.org/grpc"
)

type storage struct {
	proc processor
}

func (s *storage) Add(ctx context.Context, in *protocol_message_store.AddRequest, opts ...grpc.CallOption) (*protocol_message_store.GetPortResponse, error) {
	return s.proc.Add(ctx, in)
}

func (s *storage) Get(ctx context.Context, in *protocol_message_store.AddRequest, opts ...grpc.CallOption) (*protocol_message_store.GetPortResponse, error) {
	return s.proc.Get(ctx, in)
}
