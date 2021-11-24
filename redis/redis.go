package redis

import (
	"errors"
	"fmt"
	"gopkg.in/redis.v5"
	ottoutils "ottodigital.id/library/utils"
	"time"
)

var (
	// ClientRed ..
	ClientRed       *redis.ClusterClient
	address         string
	addresscluster1 string
	addresscluster2 string
	addresscluster3 string
	queuename       string

	// ClientRed2 ..
	ClientRed2 *redis.Client
	// ClientRed3 ..
	ClientRed3 *redis.Client

	address1 string
	address2 string
	address3 string
	port     string
	dbtype   int
	port2    string
	port3    string

	redismaster1	string
	redismaster2	string
	redismaster3	string
	redisslave1		string
	redisslave2		string
	redisslave3		string
)

func init() {
	//
	//addresscluster1 = utils.GetEnv("REDIS_HOST_CLUSTER1", "13.228.23.160:8079")
	//addresscluster2 = utils.GetEnv("REDIS_HOST_CLUSTER2", "13.228.23.160:8078")
	//addresscluster3 = utils.GetEnv("REDIS_HOST_CLUSTER3", "13.228.23.160:8077")
	//
	////addresscluster1 = beego.AppConfig.DefaultString("redis.host.cluster1", "54.254.222.187:6379")
	////addresscluster2 = beego.AppConfig.DefaultString("redis.host.cluster2", "54.254.222.187:6380")
	////addresscluster3 = beego.AppConfig.DefaultString("redis.host.cluster3", "54.254.222.187:6381")
	//
	//ClientRed = redis.NewClusterClient(&redis.ClusterOptions{
	//	Addrs: []string{addresscluster1, addresscluster2, addresscluster3},
	//})
	//pong, err := ClientRed.Ping().Result()
	//fmt.Println("Redis Ping ", pong)
	//fmt.Println("Redis Ping ", err)
	//
	//address = utils.GetEnv("REDIS_HOST", "13.228.23.160")
	//port = utils.GetEnv("REDIS_PORT", "8377")
	//// dbtype = 0
	//queuename = utils.GetEnv("REDIS_QUEUE", "ottomart")

	//ClientRed = redis.NewClusterClient(&redis.ClusterOptions{
	//	Addrs: []string{"34.101.208.23:8177", "34.101.208.23:8174",
	//		"34.101.208.23:8178", "34.101.208.23:8175",
	//		"34.101.208.23:8179", "34.101.208.23:8176"},
	//	RouteByLatency: true ,
	//	},
	//)

	redismaster1 = ottoutils.GetEnv("REDIS_MASTER_1", "34.101.208.23:8177")
	redismaster2 = ottoutils.GetEnv("REDIS_MASTER_2", "34.101.208.23:8178")
	redismaster3 = ottoutils.GetEnv("REDIS_MASTER_3", "34.101.208.23:8179")

	redisslave1 = ottoutils.GetEnv("REDIS_SLAVE_1", "34.101.208.23:8174")
	redisslave2 = ottoutils.GetEnv("REDIS_SLAVE_2", "34.101.208.23:8175")
	redisslave3 = ottoutils.GetEnv("REDIS_SLAVE_3", "34.101.208.23:8176")


	ClientRed = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{redismaster1, redisslave1,
			redismaster2, redisslave2,
			redismaster3, redisslave3},
		RouteByLatency: true ,
	},
	)
	resping := ClientRed.Ping()
	fmt.Println(resping.String())
}

// GetRedisConnection ...
func GetRedisConnection() *redis.ClusterClient {
	return ClientRed
}

// GetRedisUri ...
func GetRedisUri() string {
	return "redis://" + address + ":" + port + "/"
}

// GetQueueName ...
func GetQueueName() string {
	return queuename
}

/*func init() {
	//address = beego.AppConfig.DefaultString("redis.address", "10.10.0.119")
	//address2 = beego.AppConfig.DefaultString("redis.address", "10.10.0.123")
	//address3 = beego.AppConfig.DefaultString("redis.address", "10.10.0.124")
	//port = beego.AppConfig.DefaultString("redis.port", ":27379")
	dbtype = beego.AppConfig.DefaultInt("redis.databasetype", 1)

	//172.31.32.88
	// address = beego.AppConfig.DefaultString("redis.address.1", "127.0.0.1")
	address1 = beego.AppConfig.DefaultString("redis.address.1", "127.0.0.1")
	//address2 = beego.AppConfig.DefaultString("redis.address.2", "127.0.0.1")
	//address3 = beego.AppConfig.DefaultString("redis.address.3", "127.0.0.1")

	port = beego.AppConfig.DefaultString("redis.port.1", "6379") // 6077
	//port2 = beego.AppConfig.DefaultString("redis.port.2", "6077")
	//port3 = beego.AppConfig.DefaultString("redis.port.3", "6077")

	// port = beego.AppConfig.DefaultString("redis.port.1", "6077")
	// port2 = beego.AppConfig.DefaultString("redis.port.2", "6379")
	// port3 = beego.AppConfig.DefaultString("redis.port.3", "6379")

	//ClientRed = redis.NewFailoverClient(&redis.FailoverOptions{
	//	MasterName:    "master",
	//	SentinelAddrs: []string{address + port},
	//})
	ClientRed = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", address1, port),
		Password: "",     // no password set
		DB:       dbtype, // use default DB
	})

	redisdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	redisdb.Ping().Result()

	pong, err := ClientRed.Ping().Result()
	fmt.Println("Redis Ping ", pong)
	fmt.Println("Redis Ping ", err)

	//fmt.Println("Redis 1 Ping , err", pong, err)

	if err != nil {
		fmt.Errorf("Error Connection Redis ", err)
	}

	//ClientRed2 = redis.NewClient(&redis.Options{
	//	Addr:     fmt.Sprintf("%s:%s", address2, port2),
	//	Password: "",     // no password set
	//	DB:       dbtype, // use default DB
	//})
	//pong2, err2 := ClientRed2.Ping().Result()
	//fmt.Println("Redis 2 Ping , err", pong2, err2)
	//
	//ClientRed3 = redis.NewClient(&redis.Options{
	//	Addr:     fmt.Sprintf("%s:%s", address3, port3),
	//	Password: "",     // no password set
	//	DB:       dbtype, // use default DB
	//})
	//pong3, err3 := ClientRed3.Ping().Result()
	//fmt.Println("Redis 3 Ping  , err", pong3, err3)
}*/

// SaveDataRedis ...
func SaveDataRedis(key string, val interface{}, hourExp time.Duration) error {
	var err error
	for i := 0; i < 3; i++ {
		err = ClientRed.Set(key, val, hourExp).Err()
		if err == nil {
			break
		}
	}
	return err
}

// GetRedisKey ...
//func GetRedisKey(Key string) (string, error) {
//	val, err := ClientRed.Get(Key).Result()
//	if err != nil {
//		fmt.Println("Error : ", err.Error())
//		logs.Info("Redis 1 Not Found , key :", Key)
//		val, err = ClientRed2.Get(Key).Result()
//		if err != nil {
//			logs.Info("Redis 2 Not Found , key :", Key)
//			val, err = ClientRed3.Get(Key).Result()
//			if err != nil {
//				logs.Info("Redis 3 Not Found , key :", Key)
//			}
//		}
//		fmt.Println("Error : ", err)
//	}
//	return val, err
//}

// GetRedisKey ..
func GetRedisKey(Key string) (string, error) {
	val2, err := ClientRed.Get(Key).Result()
	if err == redis.Nil {
		err = errors.New("Key Does Not Exists")
		fmt.Println("keystruct does not exists")
	} else if err != nil {
		fmt.Println("Error : ", err.Error())
	} //else {
	//fmt.Println("keystruct", val2)
	//}
	return val2, err
}

// GetDataRedis ...
func GetDataRedis(key string, delayto, max int) (bool, string) {
	for i := 0; i < max; i++ {
		data, err := GetRedisKey(key)
		fmt.Println(" Err : ", err)
		fmt.Println(" data : ", data)
		if err == nil {
			return true, data
		}
		time.Sleep(time.Duration(delayto) * time.Second)
	}
	return false, ""
}

// GetRedisURI ...
//func GetRedisURI() string {
//	return "redis://" + address1 + ":" + port + "/"
//}

// GetRedisCounterIncr ...
func GetRedisCounterIncr(key string) (int64, error) {
	decr := ClientRed.Incr(key)
	return decr.Val(), decr.Err()

}

// SaveRedisCounter ..
func SaveRedisCounter(key string) (int64, error) {
	incr := ClientRed.Incr(key)
	return incr.Val(), incr.Err()
}

// SaveRedisCounterAuto ..
func SaveRedisCounterAuto(key string, autonom int64) (int64, error) {
	incr := ClientRed.IncrBy(key, autonom)
	return incr.Val(), incr.Err()
}

// GetRedisCounter ..
func GetRedisCounter(key string) (int64, error) {
	decr := ClientRed.Decr(key)
	return decr.Val(), decr.Err()

}

/*
 Redis Standard Set
*/
//func SaveRedis(key string, val interface{}) error {
//	var err error
//	for i := 0; i < 3; i++ {
//		err1 := ClientRed.Set(key, val, 0).Err()
//		logs.Warning("Set Redis 1 ", err1)
//		err2 := ClientRed2.Set(key, val, 0).Err()
//		logs.Warning("Set Redis 2 ", err2)
//		err3 := ClientRed3.Set(key, val, 0).Err()
//		logs.Warning("Set Redis 3 ", err3)
//		if err1 == nil || err2 == nil || err3 == nil {
//			err = nil
//			break
//		}
//	}
//	return err
//}

// SaveRedis ..
func SaveRedis(key string, val interface{}) error {
	var err error
	for i := 0; i < 3; i++ {
		err = ClientRed.Set(key, val, 0).Err()
		if err == nil {
			break
		}
	}
	return err
}

/*
 Redis Standard Set Expired
*/

// SaveRedisExp ..
func SaveRedisExp(key string, menit string, val interface{}) error {
	var err error
	for i := 0; i < 3; i++ {
		duration, _ := time.ParseDuration(menit)
		err = ClientRed.Set(key, val, duration).Err()
		if err == nil {
			break
		}
		fmt.Println("Error : ", err)
	}
	return err
}

// UpdateRedis ..
func UpdateRedis(key string) error {
	var err error
	for i := 0; i < 3; i++ {
		err = ClientRed.Del(key).Err()
		if err == nil {
			break
		}
	}
	return err
}
