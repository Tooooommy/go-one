package task

import (
	"github.com/Tooooommy/go-one/core/zapx"
	"time"
)

func TickHandler(duration int64, fn func() error) {
	defer func() {
		if result := recover(); result != nil {
			zapx.Error().Any("Recover Result", result).
				Msg("tick_handler recover")
		}
		go TickHandler(duration, fn)
	}()
	ticker := time.NewTicker(time.Duration(duration))
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := fn()
			if err != nil {
				zapx.Error().Error(err).Msg("tick_handler logic error")
			}
		}
	}
}
