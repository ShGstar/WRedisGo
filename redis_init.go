package WRedisPackage

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

type RedisProcessor struct {
	mServerAddress string
	mPassword      string
	mDBNum         int
	mThreadsNum    int
	connMap        map[*redis.Conn]bool

	redisPool   *redis.Pool
	MaxActive   int
	MaxIdle     int
	IdleTimeout time.Duration
}

var (
	instance *RedisProcessor
)

func GetInstance() *RedisProcessor {
	if instance == nil {
		log.Panic("uninitialized RedisProcessor")
	}

	return instance
}

func RedisInit(address string, password string, threadsNum int, dbNum int) {
	if instance == nil {
		instance = &RedisProcessor{}
	}

	instance.mServerAddress = address
	instance.mDBNum = dbNum
	instance.mPassword = password
	instance.mThreadsNum = threadsNum
}

func RedisPoolInit(active int, idle int, idletimeout time.Duration) {
	if instance == nil {
		instance = &RedisProcessor{}
	}
	instance.MaxActive = active
	instance.MaxIdle = idle
	instance.IdleTimeout = idletimeout
}

func (m *RedisProcessor) StartRedisPool() {
	redisPool := &redis.Pool{
		MaxActive:   m.MaxActive,   //100 最大闲置连接
		MaxIdle:     m.MaxIdle,     //10 最大活动连接数 0等于无限
		IdleTimeout: m.IdleTimeout, //10 闲置连接的超时时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", m.mServerAddress, redis.DialPassword(m.mPassword), redis.DialDatabase(m.mDBNum))
		}}
	m.redisPool = redisPool
}

func (m *RedisProcessor) StopRedisPool() {
	err := m.redisPool.Close()
	if err != nil {
		log.Panic(err)
	}
}

func (m *RedisProcessor) RedisDial() redis.Conn {
	conn, err := redis.Dial("tcp", m.mServerAddress,
		redis.DialPassword(m.mPassword), redis.DialDatabase(m.mDBNum))
	if err != nil {
		log.Panicf("Redis Connect Error : %v", err)
	}
	m.connMap[&conn] = true
	return conn
}

func (m *RedisProcessor) Close() {
	for conn := range m.connMap {
		err := (*conn).Close()
		if err != nil {
			log.Panic(err)
		}
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
