package tool

import (
	"bytes"
	"encoding/binary"
	"runtime"
	"strconv"
	"unsafe"
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
