package tool

import (
	"bytes"
	"encoding/binary"
	"math"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const (
	//DateTimeFormat 日期时间格式化
	DateTimeFormat = "2006-01-02 15:04:05"
	//DateFormat 日期式化
	DateFormat = "2006-01-02"
	//TimeFormat 时间格式化
	TimeFormat = "15:04:05"
)

//IsLittleEndian 判断大小端
func IsLittleEndian() bool {
	var i int32 = 0x01020304
	u := unsafe.Pointer(&i)
	pb := (*byte)(u)
	b := *pb
	return (b == 0x04)
}

//GoID go 协程id
func GoID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

//BytesToInt []byte to int
func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.LittleEndian, &data)
	return int(data)
}

// string转成int：
// int, err := strconv.Atoi(string)
// string转成int64：
// int64, err := strconv.ParseInt(string, 10, 64)
// int转成string：
// string := strconv.Itoa(int)
// int64转成string：
// string := strconv.FormatInt(int64,10)

//StringReplace 去除空格和换行
func StringReplace(str string) string {
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	return str
}

//Round 四舍五入
func Round(x float64) int {
	return int(math.Floor(x + 0/5))
}

//DelList 删除
func DelList(list []int32, key int32) []int32 {
	for index, v := range list {
		if v == key {
			return append(list[:index], list[index+1:]...)
		}
	}
	return list
}

//Random 100 随机
func Random(randkey float64) bool {
	rand.Seed(time.Now().Unix())
	num := rand.Intn(100)
	return num < int(randkey*100)

}
