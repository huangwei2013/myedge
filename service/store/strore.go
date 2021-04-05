package store

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type Store struct {
	Cli *clientv3.Client
	mutex  sync.RWMutex
}

func NewStore(endpointAddresses []string) (*Store, error){
	cli,err := clientv3.New(clientv3.Config{
		Endpoints: endpointAddresses,
		DialTimeout: 5 * time.Second,
	})

	//下面代码 是因为当etcd ip链接不同 不报错导致超时 阻塞的办法
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	_, err = cli.Status(timeoutCtx, endpointAddresses[0])
	log.Println(err)
	if err != nil {
		err = fmt.Errorf("error checking etcd status: %v", err)
		return nil, err
	}

	return &Store{
		Cli: cli,
	}, err

	//if nil != err{
	//	log.Fatalf("connected err : %s", err)
	//}
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	//resp, err := cli.Put(ctx, "sample_key", "sample_value")
	//cancel()
	//if err != nil {
	//	// handle error!
	//}
	//
	//log.Printf("resp : ", resp)
	//
	//defer cli.Close()
}


func (s *Store) Flush(){
	s.mutex.Lock()
	defer s.mutex.Unlock()

}

func (s *Store) Put( key string, value string){
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.Cli.Put(context.TODO(), key, value)
}

func (s *Store) Get( key string) string{
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	resp, err := s.Cli.Get(context.TODO(), key)
	if nil != err || 0 == resp.Count{
		return ""
	}
	return string(resp.Kvs[0].Value)
}

func (s *Store) GetByPrefix( prefix string) map[string]interface{}{
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	ret := make(map[string]interface{})
	resp, err := s.Cli.Get(context.TODO(), prefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if nil == err && 0 < resp.Count{
		for _, ev := range resp.Kvs {
			ret[string(ev.Key)] = ev.Value
		}
	}
	return ret
}

func (s *Store) Close(){
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.Close()
}