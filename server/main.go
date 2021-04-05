package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"myedge/server/server"
	"myedge/service/pb"
)


func startGRPCServer(server *server.Server, ctx context.Context, protocol string, address string) {
	// 监听本地端口
	lis, err := net.Listen(protocol, address)
	if err != nil {
		log.Printf("监听端口失败: %s", err)
		return
	}
	// 创建gRPC服务器
	server.GrpcServer = grpc.NewServer()

	// 注册服务
	pb.RegisterHelloServer(server.GrpcServer, server)
	pb.RegisterHeartBeatServer(server.GrpcServer, server)
	pb.RegisterDataServer(server.GrpcServer, server)
	pb.RegisterByeServer(server.GrpcServer, server)

	reflection.Register(server.GrpcServer)

	// server Run...
	go func() {
		log.Println("服务启动中")
		err = server.GrpcServer.Serve(lis)
		if err != nil {
			log.Printf("开启服务失败: %s", err)
			return
		}
		log.Println("服务退出中")
	}()

	// ctx.Done() --> server.GracefulStop()
	go func(){
		for {
			select {
			case <-ctx.Done():
				server.GrpcServer.GracefulStop()
				log.Printf("Quiting by GracefulStop")
				return
			}
		}
	}()
}

func startFlush2Store(server *server.Server, ctx context.Context){
	log.Printf("[todo] startFlush2Store")

	ticker := time.NewTicker(10 * time.Second)
	go func(){
		for {
			select {
			case <-ticker.C:
				log.Printf("flushing : %d ", len(server.Terms))

				// TODO
				if 0 < len(server.Terms) {
					server.Mutex.RLock()
					for termId, termInfo := range server.Terms {
						b, e := gjson.Encode(termInfo)
						if e != nil {
							log.Printf("error:%s",e)
							continue
						}
						server.Store.Put(termId, string(b) )
					}
					server.Mutex.RUnlock()
					server.Store.Flush()
				}

			case <-ctx.Done():
				log.Printf("Quiting")
				return
			}
		}
	}()
}

func startHttpServer(server *server.Server, ctx context.Context){
	server.HttpServer = g.Server()

	server.HttpServer.BindHandler("/server/stop", server.ServerStop)
	server.HttpServer.BindHandler("/api/term", server.GetTerm)
	server.HttpServer.BindHandler("/api/term/list", server.QueryTerms)
	server.HttpServer.BindHandler("/api/term/listStore", server.QueryTermsFromStore)

	server.HttpServer.SetPort(server.HttpPort)
	server.HttpServer.Run()
}

func main() {

	var etcdEndPoints []string
	etcdEndPoints = append(etcdEndPoints, "localhost:2379")

	server := server.NewServer(etcdEndPoints)
	server.ReloadTermsFromStore()
	server.HttpPort = 8100

	server.ServerCtx, server.ServerCancel = context.WithCancel(context.Background())

	go startGRPCServer(server, server.ServerCtx, "tcp", ":8080")
	go startFlush2Store(server, server.ServerCtx)
	go startHttpServer(server, server.ServerCtx)

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	for  {
		select {
		case <-server.ServerCtx.Done():
			log.Printf("quit running after 3 seconds")
			time.Sleep(time.Second * 3)
			return
		case <- ticker.C:
			log.Println("心跳一下，继续运行")
		}
	}
}