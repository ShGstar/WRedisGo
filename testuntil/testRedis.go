package testuntil

import (
	wRedis "../wRedisPackage"
	"fmt"
	"sync"
)

const (
	MAX_GORUNTINUE       = 100
	MAX_GORUNTINUE_NAMES = 10
)

var (
	successN = 0
	failN    = 0
)

func testgoruntinueSet(group *sync.WaitGroup) {
	for i := 0; i < MAX_GORUNTINUE_NAMES; i++ {
		instance := wRedis.GetInstance()
		name := GetRandomName()
		value := GetRandomInt(0, 100)
		taskDo := instance.PushTaskDo("set", name+"20210419", value)
		res := <-taskDo.TaskResult
		if valueRes, ok := res.(string); ok {
			//fmt.Println(valueRes)
			successN++
		} else {
			fmt.Println("faile valueRes:", valueRes)
			failN++
		}
	}

	group.Done()
}

func TestRedisLua(group *sync.WaitGroup) {
	for i := 0; i < MAX_GORUNTINUE_NAMES; i++ {
		name := GetRandomName()
		value := GetRandomInt(0, 100)
		taskDolua := wRedis.GetInstance().PushTaskLuaScrpit("testSetName", name+"20210420", value)
		res := <-taskDolua.TaskResult
		if valueRes, ok := res.(string); ok && valueRes == "OK" {
			successN++
			//fmt.Println(valueRes)
		} else {
			fmt.Println("faile valueRes:", valueRes)
			failN++
		}
	}

	group.Done()
}

func TestStart() {

	group := sync.WaitGroup{}

	for i := 0; i < MAX_GORUNTINUE; i++ {
		group.Add(1)
		//testgoruntinueSet(&group)
		TestRedisLua(&group)
	}

	group.Wait()
	fmt.Println("TestStart over ,successN/failN :", successN, " ", failN)
}
