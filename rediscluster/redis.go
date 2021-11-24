package rediscluster

import (
	"context"
	"fmt"
	"strings"
	"time"
	redis "github.com/go-redis/redis/v8"
	//"gopkg.in/redis.v5"
	ottoutils "ottodigital.id/library/utils"
)

var (
	//ctx1 = context.Background()
	redisclient *redis.ClusterClient
	ctx1        = context.Background()
	masters     []string
	slaves      []string
)

func init() {

	master := ottoutils.GetEnv("REDIS_MASTER", "13.228.23.160:8079;13.228.23.160:8078;13.228.23.160:8077")
	slave := ottoutils.GetEnv("REDIS_SLAVE", "13.228.23.160:6479;13.228.23.160:6478;13.228.23.160:6477")

	masters = strings.Split(master, ";")
	slaves = strings.Split(slave, ";")

	redisclient = openRedis()
}

func openRedis() *redis.ClusterClient {

	clusterSlots := func(ctx1 context.Context) ([]redis.ClusterSlot, error) {
		slots := []redis.ClusterSlot{
			// First node with 1 master and 1 slave.
			{
				Start: 0,
				End:   5461,
				Nodes: []redis.ClusterNode{{
					Addr: masters[0], // master
				}, {
					Addr: slaves[0], // 1st slave
				}},
			},
			// Second node with 1 master and 1 slave.
			{
				Start: 5462,
				End:   10921,
				Nodes: []redis.ClusterNode{{
					Addr: masters[1], // master
				}, {
					Addr: slaves[1], // 1st slave
				}},
			},
			// Second node with 1 master and 1 slave.
			{
				Start: 10922,
				End:   16383,
				Nodes: []redis.ClusterNode{{
					Addr: masters[2], // master
				}, {
					Addr: slaves[2], // 1st slave
				}},
			},
		}
		return slots, nil
	}

	rdb1 := redis.NewClusterClient(&redis.ClusterOptions{
		ClusterSlots:  clusterSlots,
		RouteRandomly: true,
	})

	rdb1.Ping(ctx1)

	// ReloadState reloads cluster state. It calls ClusterSlots func
	// to get cluster slots information.
	rdb1.ReloadState(ctx1)

	return rdb1
}

// GetRedisClusterClient ...
func GetRedisClusterClient() *redis.ClusterClient {
	return redisclient
}

// GetRedisKey ...
func GetRedisKey(Key string) (string, error) {
	var val string
	var err error

	for i := 0; i < 3; i++ {
		val, err = redisclient.Get(ctx1, Key).Result()
		if err == nil {
			break
		}
	}
	return val, err
}

// SaveRedis ...
func SaveRedis(key string, val interface{}) error {
	var err error
	for i := 0; i < 3; i++ {
		err = redisclient.Set(ctx1, key, val, 0).Err()
		if err == nil {
			break
		}
	}
	return err
}

// SaveRedisExp ...
func SaveRedisExp(key string, menit string, val interface{}) error {
	var err error
	for i := 0; i < 3; i++ {
		duration, _ := time.ParseDuration(menit)
		err = redisclient.Set(ctx1, key, val, duration).Err()
		if err == nil {
			break
		}
		fmt.Println("Error : ", err)
	}
	return err
}
