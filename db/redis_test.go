package db

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	StartRedis("127.0.0.1:6379", 0)
}

// func TestGetAutoID(t *testing.T) {
// 	key := GetAutoID("log")
// 	assert.Equal(t, key, int32(1))
// 	key = GetAutoID("log")
// 	assert.Equal(t, key, int32(2))
// 	key = GetAutoID("log")
// 	assert.Equal(t, key, int32(3))
// 	v, e := RedisExec("del", "log")
// 	assert.Equal(t, v, int64(1))
// 	assert.Equal(t, e, nil)
// }

//go test -bench=SaveStruct -run=XXX -benchtime=10s
// func BenchmarkSaveStruct(t *testing.B) {

// 	data := Testdata{Name: "wq", Age: 18}
// 	SetStruct("t1", data)
// 	readdata := &Testdata{}
// 	GetStruct("t1", readdata)

// 	assert.Equal(t, readdata.Age, int32(18))
// 	assert.Equal(t, readdata.Name, "wq")
// 	RedisExec("del", "t1")
// }

func TestHMGET(t *testing.T) {
	HMSET("field", "name", "天王盖地", 123, 18, "show", "23434")

	data, _ := HMGET("field", "name")
	assert.Equal(t, data["name"], "天王盖地")

	data, _ = HMGET("field", 123, "show", "s")
	assert.Equal(t, data[123], "18")
	assert.Equal(t, data["show"], "23434")
	assert.Equal(t, data["s"], "")

	log.Info(data)
	RedisExec("del", "field")

	log.Info("RedisGetInt:", RedisGetInt("test11"))
	log.Info("INCRBY:", RedisINCRBY("test11", 1))
	log.Info("RedisGetInt:", RedisGetInt("test11"))
	log.Info("INCRBY:", RedisINCRBY("test11", 1))
	log.Info("RedisGetInt:", RedisGetInt("test11"))
	log.Info("INCRBY:", RedisINCRBY("test11", -1))
	log.Info("RedisGetInt:", RedisGetInt("test11"))
	RedisExec("del", "test11")
}

// func TestSyncMap(t *testing.T) {
// 	var smp sync.Map
// 	for i := 123; i < 130; i++ {
// 		areas := bigmapmanage.AreasInfo{AreasIndex: int32(i)}
// 		smp.Store(areas.AreasIndex, areas)
// 	}

// 	smp.Range(func(key, value interface{}) bool {
// 		areas := value.(bigmapmanage.AreasInfo)
// 		b, err := json.Marshal(areas)
// 		if err != nil {
// 			return true
// 		}
// 		HMSET("areasSMap", areas.AreasIndex, b)
// 		return true
// 	})

// 	value, _ := HVALS("areasSMap")
// 	for _, v := range value {
// 		areas := &bigmapmanage.AreasInfo{}
// 		json.Unmarshal(v, areas)
// 		log.Infof("  %v", areas)
// 	}

// }
