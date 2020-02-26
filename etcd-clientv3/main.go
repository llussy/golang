package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc/grpclog"

	"github.com/coreos/etcd/clientv3"
)

func main() {

	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))

	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"http://10.21.8.13:2379"},
		Username:  "etcduser",
		Password:  "password",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	_, err = cli.Put(context.TODO(), "/foo1", "test111")
	if err != nil {
		log.Fatal(err)
	}
}
