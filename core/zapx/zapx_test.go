package zapx

import (
	"context"
	"testing"
)

func TestZapx(t *testing.T) {
	// Debug().String("hello", "world").Msg("say")
	S(context.Background()).Error("test")
}
