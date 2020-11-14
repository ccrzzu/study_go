package main

import (
	"MyProject/config"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"sync"
	"time"
)

func redisClient() (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:        config.Config.RedisAddr,
		Password:    config.Config.RedisPassword,
		DB:          0,
		PoolSize:    100,
		IdleTimeout: time.Duration(30) * time.Second,
		MaxRetries:  5,
	})
	return client
}

var client = redisClient()

func doSomething(index int) error {
	startRedis := time.Now().UnixNano() / 1000000
	//fmt.Println(index, "start to redis:", startRedis)
	SyncRedis := true
	if SyncRedis {
		data := map[string]interface{}{
			"method": "set",
			"key":    "cui_" + strconv.Itoa(index),
			"param":  map[string]string{"value": "", "ts": strconv.FormatInt(time.Now().UnixNano()/1000000, 10)},
		}
		jsonData, err := json.Marshal(data)
		if err == nil {
			err = client.LPush("db.data.global.sync", string(jsonData)).Err()
		}
		endRedis := time.Now().UnixNano() / 1000000
		fmt.Println(index, "end to redis, cost time:", endRedis-startRedis)
		return err
	} else {
		return nil
	}
}

func todoSyncData() {
	count := 10
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			doSomething(index)
		}(i)
	}
	wg.Wait()
}

func main() {
	todoSyncData()
}
