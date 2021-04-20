package main

import (
	test "./testuntil"
	app "./wRedisPackage"
	"time"
)

func main() {
	app.ConfigInit()
	/*
		app.RedisInit(app.Conf.Redis.Address, app.Conf.Redis.Password, app.Conf.Redis.ThreadsNum, 0)
		app.GetInstance().RedisDial()
		defer app.GetInstance().Close()*/

	app.RedisPoolInit(app.Conf.Redis.Address, app.Conf.Redis.Password, app.Conf.Redis.DBNum,
		app.Conf.Redis.MaxActive, app.Conf.Redis.MaxIdle, app.Conf.Redis.Idletimeout)

	defer app.GetPoolInstance().StopRedisPool()

	//app.RedisLuaInit()

	//test.TestStart()
	app.GetPoolInstance().RedisPoolStart()
	test.TestRedisLool()
	//fmt.Println("redis test")

	time.Sleep(time.Second * 1000)
}
