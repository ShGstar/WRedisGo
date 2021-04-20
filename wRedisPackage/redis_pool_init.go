package wRedisPackage

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

type RedisPoolProcessor struct {
	redisPool      *redis.Pool
	MaxActive      int
	MaxIdle        int
	IdleTimeout    int
	mServerAddress string
	mPassword      string
	mDBNum         int

	taksFIFO chan *ConnTask
	close    chan interface{}
}

var (
	instancePool *RedisPoolProcessor
)

func GetPoolInstance() *RedisPoolProcessor {
	if instancePool == nil {
		log.Panic("uninitialized RedisPoolProcessor")
	}
	return instancePool
}

func RedisPoolInit(address string, password string, dbNum int, active int, idle int, idletimeout int) {
	if instancePool == nil {
		instancePool = &RedisPoolProcessor{}
	}
	instancePool.mServerAddress = address
	instancePool.mPassword = password
	instancePool.mDBNum = dbNum
	instancePool.MaxActive = active
	instancePool.MaxIdle = idle
	instancePool.IdleTimeout = idletimeout
	instancePool.taksFIFO = make(chan *ConnTask, active)
	instancePool.close = make(chan interface{})
}

func (m *RedisPoolProcessor) ConnPoolRecover() {
	for {
		select {
		case task := <-m.taksFIFO:
			conn := m.GetRedisConn()
			value, err := conn.Do(task.cmd, task.args...)
			if err != nil {
				fmt.Println("ConnPoolRecover is err :", err)
			}
			task.TaskResult <- value
			conn.Close()
			fmt.Println("ConnPoolRecover is processing ,", task.cmd, " ", task.args)
		case <-m.close:
			fmt.Println("ConnPoolRecover is close ")
			return
		}
	}
}

func (m *RedisPoolProcessor) RedisPoolStart() {
	redisPool := &redis.Pool{
		MaxActive:   m.MaxActive,                  //100 最大闲置连接
		MaxIdle:     m.MaxIdle,                    //10 最大活动连接数 0等于无限
		IdleTimeout: time.Duration(m.IdleTimeout), //10 闲置连接的超时时间
		Dial: func() (redis.Conn, error) {
			dial, err := redis.Dial("tcp", m.mServerAddress, redis.DialPassword(m.mPassword), redis.DialDatabase(m.mDBNum))
			if err != nil {
				fmt.Println("StartRedisPool is err :", err)
			}
			/*			if _, err := dial.Do("PING");err != nil{
						fmt.Println("StartRedisPool PING is err :",err)
					}*/
			return dial, err
		}}
	m.redisPool = redisPool

	go m.ConnPoolRecover()
	fmt.Printf("StartRedisPool is success,address:%v,MaxActive:%v,MaxIdle:%v,IdleTimeout:%v\n ",
		m.mServerAddress, m.MaxActive, m.MaxIdle, m.IdleTimeout)
}

func (m *RedisPoolProcessor) StopRedisPool() {
	if m.redisPool == nil {
		panic("StopRedisPool fail,redisPool == nil, please StartRedisPool !")
	}
	err := m.redisPool.Close()
	m.close <- "close"
	if err != nil {
		log.Panic(err)
	}
}

func (m *RedisPoolProcessor) GetRedisConn() redis.Conn {
	if m.redisPool == nil {
		panic("GetRedisConn fail,redisPool == nil, please StartRedisPool !")
	}
	return m.redisPool.Get()
}

func (m *RedisPoolProcessor) PushPoolTaskDo(cmd string, args ...interface{}) *ConnTask {
	if m.redisPool == nil {
		panic("PushPoolTaskDo fail,redisPool == nil, please StartRedisPool !")
	}
	task := &ConnTask{}
	task.cmd = cmd
	task.args = args
	task.TaskResult = make(chan interface{}, 1)
	m.taksFIFO <- task
	return task
}
