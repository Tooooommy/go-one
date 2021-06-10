package breaker

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestNewBreaker(t *testing.T) {
	breaker := NewBreaker("test")
	for i := 0; i <= 1000000; i++ {
		go func(val int) {
			req, err := breaker.Execute(func() (interface{}, error) {
				// if math.Mod(float64(val), 3) == 0 {
				if val < 1000 {
					return val, errors.New("val mod")
				}
				// }
				return val, nil
			})
			fmt.Printf("val: %+v, time: %+v, state: %+v, request: %+v, error: %+v\n", val, time.Now().Format("04:05"), breaker.State(), req, err)
		}(i)
	}

	time.Sleep(1 * time.Hour)
}
