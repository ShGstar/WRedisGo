package WRedisPackage

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

type RedisProcessor struct {
	mServerAddress string
	mPassword      string
	mDBNum         int
	mThreadsNum    int
	connMap        map[*redis.Conn]bool
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

func (m *RedisProcessor) RedisDial(dialnum int) {
	for i := 0; i < dialnum; i++ {
		conn, err := redis.Dial("tcp", m.mServerAddress,
			redis.DialPassword(m.mPassword), redis.DialDatabase(m.mDBNum))
		if err != nil {
			log.Panicf("Redis Connect Error : %v", err)
		}
		m.connMap[&conn] = true
	}
	fmt.Printf("RedisDial %d success ", dialnum)
}

func (m *RedisProcessor) Close() {
	closed := 0
	for conn := range m.connMap {
		err := (*conn).Close()
		if err != nil {
			log.Panic(err)
		}
		closed++
	}
	fmt.Printf("Redis connMap %d close successful ", closed)
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
