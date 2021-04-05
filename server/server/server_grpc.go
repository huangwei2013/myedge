package server

import (
	"context"
	"fmt"
	"log"
	"myedge/server/model"
	"myedge/service/pb"
	"sync"
	"time"
)



func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("SayHello: %s", in.Id)

	// TODO: check term in storage

	if _, found := s.Terms[in.Id]; !found {
		s.Mutex.Lock()
		defer s.Mutex.Unlock()
		curTerm := model.NewTerm(in.Id)

		s.Terms[in.Id] = curTerm
		s.Terms[in.Id].Mid = 0

		s.Terms[in.Id].ServerKeeper.SdId = s.SdId
		s.Terms[in.Id].ServerKeeper.Age = 1
		s.Terms[in.Id].ServerKeeper.StartedAt = time.Now().UnixNano()
		s.Terms[in.Id].ServerKeeper.ExpiredAt = time.Now().Add(time.Duration(5 * time.Minute)).UnixNano()
	}

	tokenSeed := fmt.Sprintf("%s/%s/%s", time.Now(), in.Id, "token")

	s.Terms[in.Id].SessionToken = s.Terms[in.Id].GenSessionToken(tokenSeed)
	s.Terms[in.Id].LastConnectedTimestamp = time.Now()
	s.Terms[in.Id].LastActivedTimestamp = time.Now()
	s.Terms[in.Id].Status = model.CONST_TERM_STATUS_ACTIVED
	s.Terms[in.Id].Connected = true

	return &pb.HelloReply{Code: 0, NextMid: s.Terms[in.Id].Mid + 1}, nil
}

func (s *Server) HeartBeat(ctx context.Context, in *pb.HBRequest) (*pb.HBReply, error) {
	log.Printf("HeartBeat: %s", in.Id)

	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	if _, found := s.Terms[in.Id]; !found {
		return &pb.HBReply{Code: -1}, nil
	}

	s.Terms[in.Id].LastHeartbeatTimestamp = time.Now()
	s.Terms[in.Id].Status = model.CONST_TERM_STATUS_ACTIVED

	return &pb.HBReply{Code: 0}, nil
}

func (s *Server) DataRecv(ctx context.Context, in *pb.DataRequest) (*pb.DataReply, error) {
	log.Printf("s.Terms len: %d, DataRecv: %s", len(s.Terms), in.Id)

	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	if _, found := s.Terms[in.Id]; !found {
		return &pb.DataReply{Code: -1}, nil
	}

	records := in.GetRecords()
	for _,value := range records {
		s.Terms[in.Id].Mid = value.GetMid()
		log.Printf("key=%d, value= %s", value.GetMid(), value.GetMsg())
	}
	s.Terms[in.Id].Status = model.CONST_TERM_STATUS_ACTIVED

	return &pb.DataReply{Code: 0,  NextMid: s.Terms[in.Id].Mid + 1}, nil
}

func (s *Server) DataC2SStream(in *pb.StreamDataRequest, stream pb.Data_DataC2SStreamServer) error {
	return nil
}

func (s *Server) DataS2CStream(stream pb.Data_DataS2CStreamServer) error {
	return nil
}

func (s *Server) DataBiStreaming(stream pb.Data_DataBiStreamingServer) error{

	wg := sync.WaitGroup{}
	wg.Add(2)

	// read
	go func(){
		for s.IsRunning{
			inData, recvErr := stream.Recv()
			if nil != recvErr {
				log.Printf("接收出错：%v", recvErr)
				return
			}
			if nil != inData {
				log.Printf("接收到数据：%v", inData)
			}
		}
	}()

	//write
	go func(){
		for s.IsRunning{
			outData := "abcd"
			toSend := pb.StreamDataReply{Code:0, NextMid: 222}
			sendErr := stream.Send(&toSend)
			if nil != sendErr{
				log.Printf("发送出错：%v", sendErr)
				return
			}
			log.Printf("发送数据：%v", outData)
			time.Sleep(time.Second * 2)
		}
	}()

	wg.Wait()
	return nil
}

func (s *Server) SayBye(ctx context.Context, in *pb.ByeRequest) (*pb.ByeReply, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if _, found := s.Terms[in.Id]; found {
		s.Terms[in.Id].Status = model.CONST_TERM_STATUS_INACTIVED
		s.Terms[in.Id].Connected = false
	}

	return &pb.ByeReply{Code: 0}, nil
}
