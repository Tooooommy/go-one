package zapx

import (
	"testing"
)

func TestZapx(t *testing.T) {
	Debug().String("hello", "world").Msg("say")
	S().Error("test")
}
