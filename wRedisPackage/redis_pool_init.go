package wRedisPackage

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

type RedisPoolProcessor struct {
	redisPool      *redis.Pool
	MaxActive      int
	MaxIdle        int
	IdleTimeout    time.Duration
	mServerAddress string
	mPassword      string
	mDBNum         int
}

var (
	instancePool *RedisPoolProcessor
)

func RedisPoolInit(active int, idle int, idletimeout time.Duration) {
	if instancePool == nil {
		instancePool = &RedisPoolProcessor{}
	}
	instancePool.MaxActive = active
	instancePool.MaxIdle = idle
	instancePool.IdleTimeout = idletimeout
}

func (m *RedisPoolProcessor) StartRedisPool() {
	redisPool := &redis.Pool{
		MaxActive:   m.MaxActive,   //100 最大闲置连接
		MaxIdle:     m.MaxIdle,     //10 最大活动连接数 0等于无限
		IdleTimeout: m.IdleTimeout, //10 闲置连接的超时时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", m.mServerAddress, redis.DialPassword(m.mPassword), redis.DialDatabase(m.mDBNum))
		}}
	m.redisPool = redisPool
}

func (m *RedisPoolProcessor) StopRedisPool() {
	err := m.redisPool.Close()
	if err != nil {
		log.Panic(err)
	}
}

/*获得Redis连接*/
//通过redis连接池
/*func GetRedisConn() redis.Conn {
	if &RedisProcessor. == nil{
		str := config.Conf.Redis.Host+":"+config.Conf.Redis.Port
		redisPool = &redis.Pool{
			MaxActive:   100,
			MaxIdle:     10,
			IdleTimeout: 10,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", str)
			}}
	}
	return redisPool.Get()
}*/
