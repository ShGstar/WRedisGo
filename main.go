package main

import (
	test "./testuntil"
	app "./wRedisPackage"
	"time"
)

func main() {
	app.ConfigInit()
	app.RedisInit(app.Conf.Redis.Address, app.Conf.Redis.Password, app.Conf.Redis.ThreadsNum, 0)
	app.GetInstance().RedisDial()
	defer app.GetInstance().Close()

	app.RedisLuaInit()

	test.TestStart()

	//fmt.Println("redis test")

	time.Sleep(time.Second * 10)
}
