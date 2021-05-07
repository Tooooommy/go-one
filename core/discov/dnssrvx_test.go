package discov

import (
	"fmt"
	"github.com/Tooooommy/go-one/core/logx"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/dnssrv"
	"testing"
	"time"
)

func TestDnsx(t *testing.T) {
	instancer := dnssrv.NewInstancer("baidu.com", time.Second, logx.KitL())
	defer instancer.Stop()
	chevent := make(chan sd.Event, 10)
	instancer.Register(chevent)
	for i := 0; i < 100; i++ {
		e := <-chevent
		fmt.Printf("e: %+v\n", e.Err)
		fmt.Printf("e: %+v\n", e.Instances)

	}

	time.Sleep(5 * time.Second)
}
