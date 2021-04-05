package main

import (
	"context"
	"fmt"
	"log"
	"myedge/agent/model"
	"myedge/service/pb"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:8080"
	IdPrefix	= "ID123456"
)


var client *model.Client
var mid int64

func sayhello(client *model.Client){
	// 调用protobuf的函数创建客户端连接句柄
	var c pb.HelloClient
	if nil == client.HelloClient {
		c = pb.NewHelloClient(client.Conn)
		client.HelloClient = &c
	}else{
		c = *client.HelloClient
	}

	// 调用protobuf的SayHello函数
	r, err := c.SayHello(client.Ctx, &pb.HelloRequest{Id: client.Term.Id, Mid:mid})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// 打印结果
	log.Printf("Greeting: code=%v, mid=%d", r.GetCode(), r.GetNextMid())
	mid = r.GetNextMid()
}

func sayheartbeat(client *model.Client){
	// 调用protobuf的函数创建客户端连接句柄
	var c pb.HeartBeatClient
	if nil == client.HeartBeatClient {
		c = pb.NewHeartBeatClient(client.Conn)
		client.HeartBeatClient = &c
	}else{
		c = *client.HeartBeatClient
	}

	// 调用protobuf的Say函数
	r, err := c.HeartBeat(client.Ctx, &pb.HBRequest{Id: client.Term.Id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// 打印结果
	log.Printf("HeartBeating: code=%v", r.GetCode())
}

func saydatasend(client *model.Client){
	// 调用protobuf的函数创建客户端连接句柄
	var c pb.DataClient
	if nil == client.DataClient {
		c = pb.NewDataClient(client.Conn)
		client.DataClient = &c
	}else{
		c = *client.DataClient
	}

	records := make([]*pb.DataRecord, 0)
	records = append(records, &pb.DataRecord{
		Mid: mid + 1,
		Msg: []byte("aaaaa"),
	})
	records = append(records, &pb.DataRecord{
		Mid: mid + 2,
		Msg: []byte("fffffffffffffff"),
	})

	// 调用protobuf的函数
	r, err := c.DataRecv(client.Ctx, &pb.DataRequest{Id: client.Term.Id, Records: records})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// 打印结果
	log.Printf("DataSending: code=%v, mid=%d", r.GetCode(), r.GetNextMid())

	if 0 == r.GetCode(){
		mid = r.GetNextMid() - 1
	}
}

func saybye(client *model.Client){
	// 调用protobuf的函数创建客户端连接句柄
	var c pb.ByeClient
	if nil == client.ByeClient {
		c = pb.NewByeClient(client.Conn)
		client.ByeClient = &c
	}else{
		c = *client.ByeClient
	}

	// 调用protobuf的Say函数
	r, err := c.SayBye(client.Ctx, &pb.ByeRequest{Id: client.Term.Id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// 打印结果
	log.Printf("Bye: code=%v", r.GetCode())
}

func run() {

	var err error

	rId := fmt.Sprintf("%s_%d", IdPrefix, time.Now().UnixNano())
	client = model.NewClient(rId)
	mid = 1

	// 建立与服务器的连接
	client.Conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 延迟关闭连接
	defer client.Conn.Close()

	// context的超时设置
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client.Ctx = ctx

	sayhello(client)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(){
		i := 0
		for i < 10 {
			i++
			sayheartbeat(client)
			time.Sleep(1)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func(){
		i := 0
		for i < 10 {
			i++
			saydatasend(client)
			time.Sleep(1)
		}
		wg.Done()
	}()

	wg.Wait()

	saybye(client)

}

func main(){
	run()
}