package portdomain

import (
	"context"
	"encoding/json"
	protocol_message_store "github.com/jamiri/port-app/protocol/message/store"
	"google.golang.org/grpc"
	"log"
)

type processor struct {
	store Store
}

func (p *processor) Add(ctx context.Context, in *protocol_message_store.AddRequest) (*protocol_message_store.GetPortResponse, error) {
	jobj, err := json.Marshal(in.Data)
	if err != nil {
		log.Fatal("Store: Unable to marshal data to json")
		return nil, err
	}
	err = p.store.Add(ctx, in.Key, string(jobj))
	if err != nil {
		log.Fatal("Store: Unable to save")
		return nil, err
	}
	return nil, nil
}

func (p *processor) Get(ctx context.Context, in *protocol_message_store.AddRequest, opts ...grpc.CallOption) (*protocol_message_store.GetPortResponse, error) {
	v, err := p.store.Get(ctx, in.Key)
	if err != nil {
		log.Fatal("Store: Unable to get data")
		return nil, err
	}

	retObj := &protocol_message_store.Port{}
	err = json.Unmarshal([]byte(v), retObj)
	if err != nil {
		log.Fatal("Store: Unable to get data")
		return nil, err
	}

	return &protocol_message_store.GetPortResponse{
		Data: retObj,
	}, nil
}
