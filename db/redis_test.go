package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	StartRedis("127.0.0.1:6379")
}

func TestGetAutoID(t *testing.T) {
	key := GetAutoID("log")
	assert.Equal(t, key, int32(1))
	key = GetAutoID("log")
	assert.Equal(t, key, int32(2))
	key = GetAutoID("log")
	assert.Equal(t, key, int32(3))
	v, e := RedisExec("del", "log")
	assert.Equal(t, v, int64(1))
	assert.Equal(t, e, nil)
}

//go test -bench=SaveStruct -run=XXX -benchtime=10s
func BenchmarkSaveStruct(t *testing.B) {

	data := Testdata{Name: "wq", Age: 18}
	SetStruct("t1", data)
	readdata := &Testdata{}
	GetStruct("t1", readdata)

	assert.Equal(t, readdata.Age, int32(18))
	assert.Equal(t, readdata.Name, "wq")
}
