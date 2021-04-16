package testuntil

import (
	"math/rand"
	"sync"
	"time"
)

var (
	//姓氏
	familyNames = []string{"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "楚", "卫", "蒋", "沈", "韩", "杨", "张", "欧阳", "东门", "西门", "上官", "诸葛", "司徒", "司空", "夏侯"}
	//辈分（宗的永其光...）
	middleNamesMap = map[string][]string{}
	//名字
	lastNames = []string{"春", "夏", "秋", "冬", "风", "霜", "雨", "雪", "木", "禾", "米", "竹", "山", "石", "田", "土", "福", "禄", "寿", "喜", "文", "武", "才", "华"}
)

var (
	//随机数互斥锁（确保GetRandomInt不能被并发访问）
	randomMutex sync.Mutex
)

var namesMap = make(map[int]string)

func init() {
	for _, x := range familyNames {
		if x != "欧阳" {
			middleNamesMap[x] = []string{"德", "惟", "守", "世", "令", "子", "伯", "师", "希", "与", "孟", "由", "宜", "顺", "元", "允", "宗", "仲", "士", "不", "善", "汝", "崇", "必", "良", "友", "季", "同"}
		} else {
			middleNamesMap[x] = []string{"宗", "的", "永", "其", "光"}
		}
	}
}

//[start,end]
func GetRandomInt(start, end int) int {
	randomMutex.Lock()
	<-time.After(1 * time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := start + r.Intn(end-start+1)
	randomMutex.Unlock()
	return n
}

//产生随机数据名字
func GetRandomName() (name string) {
	familyName := familyNames[GetRandomInt(0, len(familyNames)-1)]
	middleName := middleNamesMap[familyName][GetRandomInt(0, len(middleNamesMap[familyName])-1)]
	lastName := lastNames[GetRandomInt(0, len(lastNames)-1)]
	return familyName + middleName + lastName
}

//产生n个名字
func ProductNames(n int) {
	for i := 0; i < n; i++ {
		name := GetRandomName()
		namesMap[i] = name
	}
}
