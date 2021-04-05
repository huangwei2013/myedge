package server

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"log"
	"myedge/server/model"
)

func (s *Server) ServerStop (r *ghttp.Request){
	s.IsRunning = false

	defer s.ServerCancel()

	sendHttpRsp(r,0, "OK")
}

func (s *Server) QueryTerms(r *ghttp.Request){
	if 0 < len(s.Terms) {
		s.Mutex.RLock()
		defer s.Mutex.RUnlock()
		sendHttpRspAndExit(r,0, "OK",  s.Terms)
	}

	sendHttpRspAndExit(r,0, "OK")
}

func (s *Server) QueryTermsFromStore(r *ghttp.Request){
	termsRecord:= s.Store.GetByPrefix(IdPrefix)

	terms := make(map[string]interface{})
	for termId, termInfoEncoded := range termsRecord{
		j, err := gjson.DecodeToJson(termInfoEncoded)
		if err != nil {
			log.Fatalf("DecodeToJson :%s",err)
			continue
		}
		term := model.Term{}
		if err := j.Struct(&term); err != nil {
			log.Fatalf("Struct:%s",err)
			continue
		}
		s.Terms[termId] = &term
	}

	sendHttpRspAndExit(r,0, "OK",  terms)
}

func (s *Server) TermSetType (r *ghttp.Request){


	sendHttpRspAndExit(r,0, "OK")
}

func (s *Server) GetTerm (r *ghttp.Request){
	termId := r.GetString("termid")
	termRecord := s.Store.Get(termId)
	if "" == termRecord {
		sendHttpRspAndExit(r,0, "OK")
	}

	j, err := gjson.DecodeToJson(termRecord)
	if err != nil {
		msg := fmt.Sprintf("DecodeToJson :%s",err)
		sendHttpRspAndExit(r,-1, msg)
	}
	term := model.Term{}
	if err := j.Struct(&term); err != nil {
		msg := fmt.Sprintf("Struct :%s",err)
		sendHttpRspAndExit(r,-1, msg)
	}

	sendHttpRspAndExit(r,0, "OK" ,term)
}