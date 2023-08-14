package redis

import (
	"os"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

const (
	EnvMaxIdleConns   = "REDIS_MAX_IDLE_CONNS"
	EnvMaxActiveConns = "REDIS_MAX_ACTIVE_CONNS"
	EnvRedisPool      = "REDIS_POOL"
	EnvReidsAuth      = "REDIS_AUTH"
	EnvRedisMaster    = "REDIS_MASTER"
	EnvRedisDB        = "REDIS_DB"
)

var rdb redis.UniversalClient

func Redis() redis.UniversalClient {
	return rdb
}

func Init() {
	addrs := strings.Split(os.Getenv(EnvRedisPool), ",")
	if len(addrs) == 0 {
		return
	}

	auth := os.Getenv(EnvReidsAuth)

	var db int64
	if redisDB := os.Getenv(EnvRedisDB); len(redisDB) > 0 {
		var err error
		db, err = strconv.ParseInt(redisDB, 10, 64)
		if err != nil {
			panic(err)
		}
	}
	opts := redis.UniversalOptions{
		Addrs:          addrs,
		Password:       auth,
		DB:             int(db),
		RouteByLatency: true,
	}
	master := os.Getenv(EnvRedisMaster)
	if len(addrs) > 1 {
		if len(master) == 0 {
			master = "mymaster"
		}
		opts.MasterName = master
	}

	rdb = redis.NewUniversalClient(&opts)
}
