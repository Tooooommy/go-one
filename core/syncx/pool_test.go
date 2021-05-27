package syncx

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var i int32 = 0
	pool, err := NewPool(
		SetPoolFactory(func() (interface{}, error) {
			atomic.AddInt32(&i, 1)
			return i, nil
		}),
		SetPoolClose(func(i interface{}) error {
			fmt.Println("close", i)
			return nil
		}),
		SetPoolPing(func(i interface{}) error {
			fmt.Println("ping", i)
			return nil
		}),
		SetPoolMaxCoon(1000),
	)
	if err != nil {
		t.Error(err)
		return
	}
	for j := 0; j < 10000; j++ {
		var c interface{}
		go func() {
			c, err = pool.Get()
			if err != nil {
				fmt.Println("get from pool error", err)
				return
			}

			err = pool.Put(c)
			if err != nil {
				fmt.Println("put from poll error", err)
				return
			}
		}()
	}
	time.Sleep(10 * time.Second)
	pool.Release()
}
