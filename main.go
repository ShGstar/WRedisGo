package main

import (
	app "./WRedisPackage"
	test "./testuntil"
	"fmt"
)

func main() {
	app.ConfigInit()
	app.RedisInit(app.Conf.Redis.Address, app.Conf.Redis.Password, 0, 0)
	app.GetInstance().RedisDial()
	defer app.GetInstance().Close()

	test.TestStart()

	fmt.Println("redis test")
}
