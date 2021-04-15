package WRedisPackage

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

func ConnDo(conn *redis.Conn, cmd string, args ...interface{}) interface{} {
	res, err := (*conn).Do(cmd, args...)
	if err != nil {
		log.Panicln(err)
	}
	return res
}
