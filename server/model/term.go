package model

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

const (
	CONST_TERM_STATUS_ACTIVED   = 1
	CONST_TERM_STATUS_INACTIVED = 2
	CONST_TERM_STATUS_DEAD      = 9
)

/**
	管理维度：
		- 作为终端的影子
 */
type Term struct {
	Id string	//全局唯一
	Name string	//可读性

	Type string //类型，排他
	Tags []string //标签，非排他

	Status int
	Connected bool
	LastConnectedTimestamp time.Time	//最近一次建立连接的时间
	LastActivedTimestamp time.Time 		//最近一次活跃时间( = max(LastConnectedTimestamp, LastHeartbeatTimestamp) )

	// [临时]会话相关
	SessionToken string //会话token，临时性(为空时，表示未建立有效连接)
	LastHeartbeatTimestamp time.Time

	// [动态-关键数据]
	Mid int64

	ServerKeeper	// 持有该 term 的 server 的身份信息
}

func NewTerm(Id string) *Term{

	return &Term{
		Id: Id,
	}
}


func (t *Term) GenSessionToken(seed string) string{
	hasher := md5.New()
	hasher.Write([]byte(seed))
	return hex.EncodeToString(hasher.Sum(nil)[0:])
}
