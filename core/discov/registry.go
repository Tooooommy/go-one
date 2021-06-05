package discov

import (
	"github.com/Tooooommy/go-one/core/zapx"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"strconv"
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
	for index, host := range c.cfg.Hosts {
		key := c.cfg.Name + "-" + strconv.Itoa(index)
		err := cli.Register(etcdv3.Service{
			Key:   key,
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
	for index := range c.cfg.Hosts {
		key := c.cfg.Name + "-" + strconv.Itoa(index)
		err = cli.Deregister(etcdv3.Service{Key: key})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) NewInstancer(prefix string) (sd.Instancer, error) {
	if c.cfg.HaveEtcd() {
		cli, err := c.getClient()
		if err != nil {
			return nil, err
		}
		return etcdv3.NewInstancer(cli, prefix, zapx.KitL())
	} else {
		return sd.FixedInstancer([]string{prefix}), nil
	}
}
