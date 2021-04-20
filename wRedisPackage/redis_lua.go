package wRedisPackage

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"log"
)

type RedisLuaCache struct {
	loadShaMap   map[string]*redis.Script
	luafilenames map[string]int //file name /  key count
}

var (
	instanceLua *RedisLuaCache
)

func GetRedisLuaCache() *RedisLuaCache {
	if instanceLua == nil {
		log.Panic("uninitialized RedisLuaCache")
	}

	return instanceLua
}

func (m *RedisLuaCache) setLuaAndKeyCount() {
	m.luafilenames["testSetName.lua"] = 1
	m.luafilenames["getName.lua"] = 1
}

func (m *RedisLuaCache) GetLuaScrpit(luaname string) *redis.Script {
	if script, ok := m.loadShaMap[luaname]; ok {
		return script
	}
	fmt.Println("Not Find GetLuaScrpit :", luaname)
	return nil
}

func RedisLuaInit() {
	if instanceLua == nil {
		instanceLua = &RedisLuaCache{}
	}

	instanceLua.loadShaMap = make(map[string]*redis.Script)
	instanceLua.luafilenames = make(map[string]int)
	instanceLua.setLuaAndKeyCount()

	for filename, keycount := range instanceLua.luafilenames {
		bytes, err := ioutil.ReadFile("../wRedisPackage/wReadsLua/" + filename)
		if err != nil {
			fmt.Println("ioutil.ReadFile Lua is fail: ", err, filename)
			return
		}

		luafile := string(bytes)
		script := redis.NewScript(keycount, luafile)
		instanceLua.loadShaMap[filename[:len(filename)-4]] = script
		fmt.Println("loadShaMap lua name: ", filename, "  keycount:", keycount)
	}
}
