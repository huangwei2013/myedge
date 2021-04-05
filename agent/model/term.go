package model

import (
	"time"
)



type Term struct {
	Id string	//全局唯一
	Name string	//可读性

	Status int
	Connected bool
	LastConnectedTimestamp time.Time	//最近一次建立连接的时间

	// [临时]会话相关
	SessionToken string //会话token，临时性(为空时，表示未建立有效连接)
	LastHeartbeatTimestamp time.Time
}

func NewTerm(Id string) *Term{
	return &Term{
		Id: Id,
	}
}
