package discov

import (
	"github.com/go-kit/kit/sd/etcdv3"
	"time"
)

type (

	// Registry
	Registry interface {
		Register() error
		Deregister() error
	}
)

func NewRegistry(cfg *Config) Registry {
	cli := NewClient(cfg)
	return cli
}

func (c *Client) Register() error {
	if c.cfg.HaveEtcd() {
		return nil
	}
	// 注册服务
	cli, err := c.getClient()
	if err != nil {
		return err
	}
	heartbeat := time.Duration(c.cfg.Heartbeat) * time.Second
	ttl := time.Duration(c.cfg.Ttl) * time.Second
	for _, host := range c.cfg.Hosts {
		err := cli.Register(etcdv3.Service{
			Key:   c.cfg.Key + "/" + host,
			Value: host,
			TTL:   etcdv3.NewTTLOption(heartbeat, ttl),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) Deregister() error {
	if c.cfg.HaveEtcd() {
		return nil
	}
	cli, err := c.getClient()
	if err != nil {
		return err
	}
	for _, host := range c.cfg.Hosts {
		err = cli.Deregister(etcdv3.Service{
			Key: c.cfg.Key + "/" + host,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
