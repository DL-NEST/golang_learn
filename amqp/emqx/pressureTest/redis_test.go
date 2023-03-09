package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
	"testing"
	"time"
)

func TestConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	//var lock sync.Mutex

	wg.Add(2)
	Rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	go func() {
		for i := 0; i < 50000; i++ {
			Rdb.LPush(context.Background(), "test", i)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 50000; i++ {
			Rdb.RPop(context.Background(), "test")
		}
		wg.Done()
	}()
	wg.Wait()
}

func BenchmarkName(b *testing.B) {
	Rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; i < 50000; i++ {
			Rdb.LPush(context.Background(), "test", i)
		}
	})
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; i < 50000; i++ {
			Rdb.RPop(context.Background(), "test")
		}
	})
}

func TestLink(t *testing.T) {
	Rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	for {
		time.Sleep(1 * time.Second)
		Rdb.LLen(context.Background(), "logCache").Val()
	}
}

func TestSend(t *testing.T) {
	Rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	for i := 0; i < 1000; i++ {
		Rdb.RPush(context.Background(), "logCache", "res")
	}
}

func TestDV(t *testing.T) {
	aa := time.UnixMilli(1678368596628).Format("2006-01-02 15:04:05.000")

	t.Logf("%s", aa)
	t.Logf("%d", time.Now().UnixMilli())
}
