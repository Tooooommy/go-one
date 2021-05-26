package syncx

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

type (
	Pool interface {
		Get() (interface{}, error)
		Put(interface{}) error
		Close(interface{}) error
		Release()
		Len() int
	}

	PoolOption func(p *pool)

	pool struct {
		mux      sync.RWMutex
		conn     chan *conn
		factory  func() (interface{}, error)
		close    func(interface{}) error
		ping     func(interface{}) error
		idleTime time.Duration
		maxConn  int32 // 最大连接数
		openConn int32 // 连接数
	}

	conn struct {
		conn interface{}
		idle time.Time
	}
)

var (
	defaultPoolValue = 5
	ErrClosed        = errors.New("pool is closed")
)

func NewPool(options ...PoolOption) (Pool, error) {

	p := &pool{
		idleTime: time.Duration(defaultPoolValue) * time.Second, // 空闲时间
		maxConn:  int32(defaultPoolValue),
		openConn: int32(defaultPoolValue),
	}
	for _, option := range options {
		option(p)
	}
	if p.factory == nil {
		return nil, errors.New("factory func invalid")
	}
	if p.close == nil {
		return nil, errors.New("close func invalid")
	}
	if p.openConn > p.maxConn {
		return nil, errors.New("max_count must greater than or equal to open_count")
	}
	p.conn = make(chan *conn, p.maxConn)

	for i := 0; i < int(p.openConn); i++ {
		c, err := p.factory()
		if err != nil {
			p.Release()
			return nil, err
		}
		p.conn <- &conn{conn: c, idle: time.Now()}
	}

	return p, nil
}

func SetFactory(factory func() (interface{}, error)) PoolOption {
	return func(p *pool) {
		p.factory = factory
	}
}

func SetClose(close func(interface{}) error) PoolOption {
	return func(p *pool) {
		p.close = close
	}
}

func SetPing(ping func(interface{}) error) PoolOption {
	return func(p *pool) {
		p.ping = ping
	}
}

func SetIdleTime(idleTime time.Duration) PoolOption {
	return func(p *pool) {
		if idleTime > 0 {
			p.idleTime = idleTime
		}
	}
}

func SetMaxCoon(max int32) PoolOption {
	return func(p *pool) {
		if max > 0 {
			p.maxConn = max
		}
	}
}

func SetOpenConn(open int32) PoolOption {
	return func(p *pool) {
		if open > 0 {
			p.openConn = open
		}
	}
}

// getConn 获取所有连接
func (p *pool) getConn() chan *conn {
	p.mux.RLock()
	conn := p.conn
	p.mux.RUnlock()
	return conn
}

func (p *pool) Get() (interface{}, error) {
	conn := p.getConn()
	if conn == nil {
		return nil, ErrClosed
	}
	for {
		select {
		case c, ok := <-conn:
			if !ok {
				return nil, ErrClosed
			}
			if p.idleTime > 0 && c.idle.Add(p.idleTime).Before(time.Now()) {
				_ = p.Close(c.conn)
				continue
			}
			if err := p.Ping(c.conn); err != nil {
				_ = p.Close(c.conn)
				continue
			}
			return c.conn, nil
		default:
			if p.openConn > p.maxConn {
				return nil, errors.New("connection reach limit")
			}
			c, err := p.factory()
			if err != nil {
				return nil, err
			}
			atomic.AddInt32(&p.openConn, 1)
			p.openConn++
			return c, nil
		}
	}
}

func (p *pool) Put(c interface{}) error {
	if c == nil {
		return errors.New("connection is nil. rejecting")
	}
	if p.conn == nil {
		return ErrClosed
	}

	select {
	case p.conn <- &conn{conn: c, idle: time.Now()}:
		atomic.AddInt32(&p.openConn, 1)
		return nil
	default:
		return p.Close(c)
	}
}

func (p *pool) Close(c interface{}) error {
	if p.conn == nil {
		return ErrClosed
	}
	if c == nil {
		return errors.New("connection is nil, rejecting")
	}
	atomic.AddInt32(&p.openConn, -1)

	if p.close == nil {
		return nil
	}
	return p.close(c)
}

func (p *pool) Ping(c interface{}) error {
	if p.conn == nil {
		return ErrClosed
	}
	if c == nil {
		return errors.New("connection is nil, rejecting")
	}
	if p.ping == nil {
		return nil
	}
	return p.ping(c)
}

func (p *pool) Release() {
	p.mux.Lock()
	conn := p.conn
	clos := p.close
	p.conn = nil
	p.factory = nil
	p.ping = nil
	p.close = nil
	p.mux.Unlock()
	if conn == nil {
		return
	}
	close(conn)
	if clos != nil {
		for c := range conn {
			_ = clos(c)
		}

	}
	return
}

func (p *pool) Len() int {
	return len(p.getConn())
}
