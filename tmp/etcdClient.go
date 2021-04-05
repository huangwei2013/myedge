package main

import (
	"context"
	"fmt"
	"time"

	ec "go.etcd.io/etcd/clientv3"
)
const (
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
)
func main() {
	var (
		config ec.Config
		client *ec.Client
		err    error
		kv     ec.KV
		putRes *ec.PutResponse
	)

	config = ec.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 30 * time.Second,
	}

	if client, err = ec.New(config); err != nil {
		fmt.Println(err)
		return
	}

	kv = ec.NewKV(client)

	if putRes, err = kv.Put(context.TODO(), "/cron/jobs/job3", "111111"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("revision: ", putRes.Header.Revision)
		if putRes.PrevKv != nil{
			fmt.Println(string(putRes.PrevKv.Value))
		}
	}


	fmt.Println("-------------")
	op:= ec.OpGet("/cron/jobs/job3")
	fmt.Println(string(op.KeyBytes()))
	fmt.Println(string(op.ValueBytes()))

	fmt.Println("-------------")
	var getResp *ec.GetResponse
	getResp, err = kv.Get(context.TODO(),"/cron/jobs/job3")
	fmt.Println(getResp.Kvs)

	ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := client.Get(ctx, "ID123456", ec.WithPrefix(), ec.WithSort(ec.SortByKey, ec.SortDescend))
	if err != nil {
		fmt.Errorf("err:%s",err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
