package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"log"
)

type Discovery struct {
	cli   *clientv3.Client
	info  *NodeInfo
	nodes *NodesManager
}

func NewDiscovery(info *NodeInfo, conf clientv3.Config, mgr *NodesManager) (dis *Discovery, err error) {
	d := &Discovery{}
	d.info = info
	if mgr == nil {
		return nil, fmt.Errorf("[Discovery] mgr == nil")
	}
	d.nodes = mgr
	d.cli, err = clientv3.New(conf)
	return d, err
}

func (d *Discovery) pull() {
	kv := clientv3.NewKV(d.cli)
	resp, err := kv.Get(context.TODO(), "discovery/", clientv3.WithPrefix())
	if err != nil {
		log.Fatalf("[Discovery] kv.Get err:%+v", err)
		return
	}
	for _, v := range resp.Kvs {
		node := &NodeInfo{}
		err = json.Unmarshal(v.Value, node)
		if err != nil {
			log.Fatalf("[Discovery] json.Unmarshal err:%+v", err)
			continue
		}
		d.nodes.AddNode(node)
		log.Printf("[Discovery] pull node:%+v", node)
	}
}

func (d *Discovery) watch() {
	watcher := clientv3.NewWatcher(d.cli)
	watchChan := watcher.Watch(context.TODO(), "discovery", clientv3.WithPrefix())
	for {
		select {
		case resp := <-watchChan:
			d.watchEvent(resp.Events)
		}
	}
}

func (d *Discovery) watchEvent(evs []*clientv3.Event) {
	for _, ev := range evs {
		switch ev.Type {
		case clientv3.EventTypePut:
			node := &NodeInfo{}
			err := json.Unmarshal(ev.Kv.Value, node)
			if err != nil {
				log.Fatalf("[Discovery] json.Unmarshal err:%+v", err)
				continue
			}
			d.nodes.AddNode(node)
			log.Printf(fmt.Sprintf("[Discovery] new node:%s", string(ev.Kv.Value)))
		case clientv3.EventTypeDelete:
			d.nodes.DelNode(string(ev.Kv.Key))
			log.Printf(fmt.Sprintf("[Discovery] del node:%s data:%s", string(ev.Kv.Key), string(ev.Kv.Value)))
		}
	}
}
