package redis_api

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type RedisAPISuite struct {
	suite.Suite
	client redis.Cmdable
}

func initRedis() redis.Cmdable {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func (r *RedisAPISuite) SetupSuite() {
	r.client = initRedis()

}

func (r *RedisAPISuite) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	r.client.FlushAll(ctx)
}

func (r *RedisAPISuite) TestString() {
	t := r.T()

	testCases := []struct {
		name string

		before func(t *testing.T, ctx context.Context)
		after  func(t *testing.T, ctx context.Context)
	}{
		// set/get操作
		{
			name: "set and get key",
			before: func(t *testing.T, ctx context.Context) {
				err := r.client.Set(ctx, "test", "test", -1).Err()
				assert.NoError(t, err)
			},

			after: func(t *testing.T, ctx context.Context) {
				val := r.client.Get(ctx, "test").Val()
				assert.Equal(t, "test", val)
			},
		},
		// del操作
		{
			name: "del key",
			before: func(t *testing.T, ctx context.Context) {
				// 设置两个key
				err := r.client.Set(ctx, "test1", "test1", -1).Err()
				assert.NoError(t, err)
				err = r.client.Set(ctx, "test2", "test2", -1).Err()
				assert.NoError(t, err)

				err = r.client.Del(ctx, "test1").Err()
				assert.NoError(t, err)
			},

			after: func(t *testing.T, ctx context.Context) {
				err := r.client.Get(ctx, "test1").Err()
				assert.Equal(t, redis.Nil, err)
				val := r.client.Get(ctx, "test2").Val()
				assert.Equal(t, "test2", val)
			},
		},
		// setnx操作
		{
			name: "setnx key",
			before: func(t *testing.T, ctx context.Context) {
				// 设置两个key
				err := r.client.SetNX(ctx, "test1", "test1", -1).Err()
				assert.NoError(t, err)
				err = r.client.SetNX(ctx, "test1", "test2", -1).Err()
				assert.NoError(t, err)

			},

			after: func(t *testing.T, ctx context.Context) {

				val := r.client.Get(ctx, "test1").Val()
				assert.Equal(t, "test1", val)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
			defer cancel()
			tc.before(t, ctx)
			tc.after(t, ctx)
		})
	}

}

func (r *RedisAPISuite) TestList() {
	t := r.T()

	testCases := []struct {
		name string

		before func(t *testing.T, ctx context.Context)
		after  func(t *testing.T, ctx context.Context)
	}{
		// LPUSH/RPUSH/LRANGE操作
		{
			name: "LPUSH/RPUSH/LRANGE key",
			before: func(t *testing.T, ctx context.Context) {
				err := r.client.LPush(ctx, "test", []any{1, 2, 3, "test", "ljk"}).Err()
				assert.NoError(t, err)
				err = r.client.RPush(ctx, "test", []any{"test", "ljk"}).Err()
				assert.NoError(t, err)
			},

			after: func(t *testing.T, ctx context.Context) {
				val := r.client.LRange(ctx, "test", 0, -1).Val()
				assert.Equal(t, []string{
					"ljk", "test", "3", "2", "1", "test", "ljk",
				}, val)
			},
		},
		// LPOP/RPOP操作
		{
			name: "LPOP/RPOP key",
			before: func(t *testing.T, ctx context.Context) {
				err := r.client.RPush(ctx, "test", []any{1, 2, 3, 4, 5}).Err()
				assert.NoError(t, err)
			},

			after: func(t *testing.T, ctx context.Context) {
				val := r.client.LPop(ctx, "test").Val()
				assert.Equal(t, "1", val)

				val = r.client.RPop(ctx, "test").Val()
				assert.Equal(t, "5", val)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
			defer cancel()
			tc.before(t, ctx)
			tc.after(t, ctx)
		})
	}

}

func (r *RedisAPISuite) TestSet() {
	t := r.T()

	testCases := []struct {
		name string

		before func(t *testing.T, ctx context.Context)
		after  func(t *testing.T, ctx context.Context)
	}{
		// LPUSH/RPUSH/LRANGE操作
		{
			name: "LPUSH/RPUSH/LRANGE key",
			before: func(t *testing.T, ctx context.Context) {
				err := r.client.LPush(ctx, "test", []any{1, 2, 3, "test", "ljk"}).Err()
				assert.NoError(t, err)
				err = r.client.RPush(ctx, "test", []any{"test", "ljk"}).Err()
				assert.NoError(t, err)
			},

			after: func(t *testing.T, ctx context.Context) {
				val := r.client.LRange(ctx, "test", 0, -1).Val()
				assert.Equal(t, []string{
					"ljk", "test", "3", "2", "1", "test", "ljk",
				}, val)
			},
		},
		// LPOP/RPOP操作
		{
			name: "LPOP/RPOP key",
			before: func(t *testing.T, ctx context.Context) {
				err := r.client.RPush(ctx, "test", []any{1, 2, 3, 4, 5}).Err()
				assert.NoError(t, err)
			},

			after: func(t *testing.T, ctx context.Context) {
				val := r.client.LPop(ctx, "test").Val()
				assert.Equal(t, "1", val)

				val = r.client.RPop(ctx, "test").Val()
				assert.Equal(t, "5", val)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
			defer cancel()
			tc.before(t, ctx)
			tc.after(t, ctx)
		})
	}

}

func TestRedisAPI(t *testing.T) {
	suite.Run(t, &RedisAPISuite{})
}
