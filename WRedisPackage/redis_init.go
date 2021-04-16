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
	mThreadsNum    int //开启N个conn N个协程并发
	connMap        map[*redis.Conn]bool
	Done           chan interface{}
	close          chan interface{}
	taksFIFO       chan *connTask
}
type connTask struct {
	cmd  string
	args []interface{}
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
		instance.connMap = make(map[*redis.Conn]bool)
	}

	instance.mServerAddress = address
	instance.mDBNum = dbNum
	instance.mPassword = password
	instance.mThreadsNum = threadsNum
	instance.Done = make(chan interface{})
	instance.close = make(chan interface{}, threadsNum)
}

func (m *RedisProcessor) RedisDial() {
	for i := 0; i < m.mThreadsNum; i++ {
		conn, err := redis.Dial("tcp", m.mServerAddress,
			redis.DialPassword(m.mPassword), redis.DialDatabase(m.mDBNum))
		if err != nil {
			log.Panicf("Redis Connect Error : %v", err)
		}
		m.connMap[&conn] = true

		go m.ConnRecover(&conn)
	}
	fmt.Printf("RedisDial %d success ", m.mThreadsNum)
}

func (m *RedisProcessor) ConnRecover(conn *redis.Conn) {
	for {
		select {

		case <-m.close:
			m.Done <- "close conn Goroutinue"
			return
		}
	}
}

func (m *RedisProcessor) Close() {
	close(m.close)
	closed := 0
	for conn := range m.connMap {
		err := (*conn).Close()
		if err != nil {
			log.Panic(err)
			continue
		}
		closed++
	}

	//wait conn Goroutinue done
	for i := 0; i < m.mThreadsNum; i++ {
		<-m.Done
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
