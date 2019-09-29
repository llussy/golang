package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	cli ,err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed,err:",err)
		return
	}
	fmt.Println("connect succ")
	defer cli.Close()

	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	_ ,err  = cli.Put(ctx,"/deploy/anfeng","test")
	cancel()
	if err != nil {
		fmt.Println("put failed,err:",err)
		return
	}

	ctx ,cancel = context.WithTimeout(context.Background(),time.Second)
	resp,err := cli.Get(ctx,"/deploy/anfeng")
	cancel()
	if err != nil {
		fmt.Println("get failed,err:",err)
		return
	}
	for _,ev := range resp.Kvs {
		fmt.Printf("%s: %s\n",ev.Key,ev.Value)
	}
}