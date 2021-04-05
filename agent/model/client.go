package model

import (
	"context"
	"google.golang.org/grpc"
	"myedge/service/pb"
)

type Client struct {
	Conn *grpc.ClientConn

	HelloClient *pb.HelloClient
	HeartBeatClient *pb.HeartBeatClient
	DataClient *pb.DataClient
	ByeClient *pb.ByeClient

	Ctx context.Context

	Term *Term
}


func NewClient(Id string) *Client{

	term := NewTerm(Id)
	return &Client{
		Term: term,
	}
}