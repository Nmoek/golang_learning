package context

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestContextValue(t *testing.T) {
	ctx := context.WithValue(context.Background(), "test_key", "test_val")

	val, ok := ctx.Value("test_key").(string)
	assert.Equal(t, true, ok)

	t.Log(val, ok)
}

func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {

		time.Sleep(10 * time.Second)
		cancel()
	}()

	<-ctx.Done()
	t.Log("已经cancel")

}

func TestContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	<-ctx.Done()

	t.Log("超时取消了")
}
