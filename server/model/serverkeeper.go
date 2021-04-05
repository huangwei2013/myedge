package model


type ServerKeeper struct{
	SdId string
	StartedAt int64	//持有的起始时间
	ExpiredAt int64	//持有的过期时间，当过期后不再有效持有
	Age int // 每次持有更迭时 + 1
}


func NewServerKeeper(SdId string) *ServerKeeper{

	return &ServerKeeper{
		SdId: SdId,
		Age: 0,
	}
}
