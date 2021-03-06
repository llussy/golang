package main

import (
	"context"
	"time"

	"github.com/coreos/etcd/clientv3"
	"fmt"
)

func main() {
	cli ,err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed ,err:",err)
		return

	}
	fmt.Println("connect succ")
	defer cli.Close()

	_, err  =  cli.Put(context.Background(),"/logagent/conf/","8888")
	if err != nil {
		fmt.Println("put failed, err:", err)
	}
	for {
		rch := cli.Watch(context.Background(),"/logagent/conf/")
		for wresp := range rch {
			for _,ev := range wresp.Events {
				fmt.Printf("%s %q: %q\n",ev.Type,ev.Kv.Key,ev.Kv.Value)
			}
		}
	}
}