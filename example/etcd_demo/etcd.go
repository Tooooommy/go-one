package main

import (
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
)

func main() {

	nodes := NewNodeManager()
	dis, _ := NewDiscovery(&NodeInfo{
		Name: "server name/aaaa",
		Addr: "127.0.0.1:8888",
	}, clientv3.Config{
		Endpoints:   []string{"0.0.0.0:2479"},
		DialTimeout: 5 * time.Second,
	}, nodes)

	reg, _ := NewRegister(&NodeInfo{
		Name:     "testsvr",
		Addr:     "127.0.0.1:8888",
		UniqueId: "discovery/testsvr/instance_id/aaabbbccc",
	}, clientv3.Config{
		Endpoints:   []string{"0.0.0.0:2479"},
		DialTimeout: 5 * time.Second,
	})

	reg2, _ := NewRegister(&NodeInfo{
		Name:     "testsvr",
		Addr:     "127.0.0.1:8888",
		UniqueId: "discovery/testsvr/instance_id/testqqqqq",
	}, clientv3.Config{
		Endpoints:   []string{"0.0.0.0:2479"},
		DialTimeout: 5 * time.Second,
	})

	go reg.Run()
	time.Sleep(time.Second * 2)
	dis.pull()
	go dis.watch()
	time.Sleep(time.Second * 1)
	go reg2.Run()
	time.Sleep(time.Second * 1)
	nodes.Dump()
	log.Printf("[Main] nodes pick:%+v", nodes.Pick("testsvr"))
	time.Sleep(time.Second * 5)
}
