package tool

import (
	"bytes"
	crand "crypto/rand" //加密安全的随机库
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"math/rand" //伪随机库
	"runtime"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/robfig/cron/v3"
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

func RandInt64ForRange(min, max int64) int64 {
	if min >= max {
		return min
	}
	maxBigInt := big.NewInt(max + 1 - min)
	i, err := crand.Int(crand.Reader, maxBigInt)
	if err != nil {
		return min
	}
	i64 := i.Int64()
	return i64 + min
}

//权重随机
func RandWeight(values []int64) int64 {
	var total int64
	for _, v := range values {
		total += v
	}
	ranNum := RandInt64ForRange(0, total)
	for i, v := range values {
		ranNum -= v
		if ranNum <= 0 {
			return int64(i)
		}
	}
	return 0
}

//对比时间范围 startStr<  difftime < endStr
func DiffCronStrNowTime(difftime time.Time, startStr, endStr string) bool {
	parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	startTime, serr := parser.Parse(startStr)
	endTime, eerr := parser.Parse(endStr)
	if serr != nil || eerr != nil {
		fmt.Println("不合法的 cron 格式 ", startStr, endStr)
		return false
	}
	//log.Debug("活动开放时间：", startTime.Next(time.Now()), endTime.Next(time.Now()))

	return startTime.Next(difftime).Unix() > endTime.Next(difftime).Unix()
}

//每日凌晨时间
func GetToDayStartUnix() int64 {
	timeStr := time.Now().Format(DateFormat)
	t2, _ := time.ParseInLocation(DateFormat, timeStr, time.Local)
	return t2.Unix()
}
