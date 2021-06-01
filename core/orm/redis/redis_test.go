package redis

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	cli, err := NewClient(context.Background())
	if err != nil {
		panic(err)
	}
	orm := cli.ORM().Pipeline()
	orm.Set(context.Background(), "key2", "value2", 24*time.Hour)
	orm.Set(context.Background(), "key3", "value3", 24*time.Hour)
	result, err := orm.Exec(context.Background())
	fmt.Println(err)
	fmt.Println(result)
}
