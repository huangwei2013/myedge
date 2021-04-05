package server

import (
	"context"
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"

	"github.com/gogf/gf/net/ghttp"

	"myedge/server/model"
	"myedge/service/store"
)

const (
	IdPrefix	= "ID123456"
)

type Server struct {

	SdId string // Server动态ID，需要唯一性，每次启动时生成，并在该生命周期中保持不变

	HttpPort int
	HttpServer *ghttp.Server

	GrpcServer *grpc.Server

	IsRunning bool
	ServerCtx context.Context
	ServerCancel context.CancelFunc

	GrpcHost string
	GrpcPort int
	EtcdEndPoints []string

	logger  *log.Logger
	Mutex   sync.RWMutex
	Store   *store.Store
	Terms   map[string]*model.Term
	Regions map[string]*model.Region
}


func NewServer(etcdEndPoints []string) *Server {

	SdId := fmt.Sprintf("SdId_%d", time.Now().UnixNano())

	terms := make(map[string]*model.Term)
	store, _ := store.NewStore(etcdEndPoints)

	return &Server{
		SdId: SdId,
		IsRunning: true,
		EtcdEndPoints: etcdEndPoints,
		Terms: terms,
		Regions: make(map[string]*model.Region),
		Store: store,
	}
}

/**
 * reload term from store
 */
func (s *Server) ReloadTermsFromStore() {
	termsRecord:= s.Store.GetByPrefix(IdPrefix)

	terms := make(map[string]interface{})
	for termId, termInfoEncoded := range termsRecord{
		terms[termId],_ = gjson.Decode(termInfoEncoded)
		j, err := gjson.DecodeToJson(termInfoEncoded)
		if err != nil {
			log.Fatalf("reload %s failed where decoded : %s", termId, err)
			continue
		}
		term := model.Term{}
		if err := j.Struct(&term); err != nil {
			log.Fatalf("reload %s failed where trans struct :%s", termId, err)
			continue
		}
		s.Terms[termId] = &term
	}
}

func sendHttpRspAndExit(r *ghttp.Request, code int, msg string, data ...interface{}){
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(g.Map{
		"code":  code,
		"message":  msg,
		"data": responseData,
	})
	r.Exit()
}


func sendHttpRsp(r *ghttp.Request, code int, msg string, data ...interface{}){
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(g.Map{
		"code":  code,
		"message":  msg,
		"data": responseData,
	})
}