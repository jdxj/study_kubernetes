package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rc  *redis.Client
	rcc *redis.ClusterClient
)

func InitClient(host, port, pass string, dbNum int) error {
	rc = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: pass,
		DB:       dbNum,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return rc.Ping(ctx).Err()
}

func InitCluster(addrs []string) error {
	rcc = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addrs,
		ReadOnly: true,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return rcc.Ping(ctx).Err()
}
