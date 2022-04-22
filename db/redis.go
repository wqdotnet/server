package db

import (
	//"encoding/json"
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"

	red "github.com/gomodule/redigo/redis"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//Redis redis
type Redis struct {
	pool *red.Pool
}

var redis *Redis

//StartRedis 初始化
func StartRedis(address string, selectdb int) {
	logrus.Infof("StartRedis  create redis.pool:  [%v]", address)
	redis = new(Redis)
	redis.pool = &red.Pool{
		MaxIdle:     256,
		MaxActive:   0,
		IdleTimeout: time.Duration(120),
		Dial: func() (red.Conn, error) {
			return red.Dial(
				"tcp",
				address,
				red.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				red.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				red.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				red.DialDatabase(selectdb),
				//red.DialPassword(""),
			)
		},
	}
}

func RedisConn() (bool, error) {
	con := redis.pool.Get()
	if err := con.Err(); err != nil {
		return false, err
	}
	return true, nil
}

//RedisExec 命令
func RedisExec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	con := redis.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	parmas := make([]interface{}, 0)
	parmas = append(parmas, key)

	if len(args) > 0 {
		parmas = append(parmas, args...)
	}
	return con.Do(cmd, parmas...)
}

//GetAutoID 获取自增id
func GetAutoID(tabname string) int32 {
	autoid, err := red.Int(RedisExec("incr", tabname))
	if err != nil {
		logrus.Error(err)
	}
	return int32(autoid)
}

//RedisINCRBY RedisINCRBY num
func RedisINCRBY(tabname string, num int32) int32 {
	autoid, err := red.Int(RedisExec("INCRBY", tabname, num))
	if err != nil {
		logrus.Error(err)
	}
	return int32(autoid)
}

func RedisDel(key string) {
	_, err := RedisExec("del", key)
	if err != nil {
		logrus.Error(err)
	}
}

//RedisGetInt get redis int
func RedisGetInt(tabname string) int {
	if data, err := red.Int(RedisExec("get", tabname)); err == nil {
		return data
	}
	return 0
}

//RedisSetStruct save struct
func RedisSetStruct(key string, v interface{}) (interface{}, error) {
	conn := redis.pool.Get()
	defer conn.Close()
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return conn.Do("SET", key, string(b))
}

//GetStruct get
// func GetStruct(key string, obj interface{}) error {
// 	conn := redis.pool.Get()
// 	defer conn.Close()
// 	objStr, err := red.String(conn.Do("GET", key))
// 	if err != nil {
// 		return err
// 	}
// 	b := []byte(objStr)
// 	err = json.Unmarshal(b, obj)
// 	return err
// }

//GetStruct get
func GetStruct[T any](key string) (*T, error) {
	conn := redis.pool.Get()
	defer conn.Close()
	objStr, err := red.String(conn.Do("GET", key))

	obj := new(T)
	if err != nil {
		return obj, err
	}
	b := []byte(objStr)

	err = json.Unmarshal(b, obj)
	return obj, err
}

//HMSET HMSET redsi
func HMSET(field string, args ...interface{}) {
	RedisExec("HMSET", field, args...)
}

//HMGET HMGET redsi
func HMGET(field string, keys ...interface{}) (map[interface{}]string, error) {
	value, err := red.Values(RedisExec("HMGET", field, keys...))
	if err != nil {
		return nil, err
	}

	var returnmap = make(map[interface{}]string)
	for k, v := range value {
		if v == nil {
			break
		}
		v := v.([]byte)
		returnmap[keys[k]] = string(v)
	}
	return returnmap, err
}

//HVALS 获取所有值
func HVALS(field string) (map[interface{}][]byte, error) {
	value, err := red.Values(RedisExec("HVALS", field))
	if err != nil {
		return nil, err
	}

	var returnmap = make(map[interface{}][]byte)
	for k, v := range value {

		if v == nil {
			break
		}
		v := v.([]byte)

		returnmap[k] = v
	}

	return returnmap, err
}

//==========================有序集合==================================================
//保存数据  key   score:分值   member 数据
func RedisZADD(key string, score int64, member interface{}) (reply interface{}, e error) {

	switch member.(type) {
	case string:
		return RedisExec("ZADD", key, score, member)
	default:
		b, _ := json.Marshal(member)
		return RedisExec("ZADD", key, score, string(b))
	}

}

//返回有序集中成员的排名。其中有序集成员按分数值递减(从大到小)排序。
func RedisZrevrank(key string, member interface{}) (int, error) {
	switch member.(type) {
	case string:
		return red.Int(RedisExec("ZREVRANK", key, member))
	default:
		b, _ := json.Marshal(member)
		fmt.Println("数据：", string(b))
		return red.Int(RedisExec("ZREVRANK", key, string(b)))
	}

}

//有序集合成员数
func RedisZCARD(key string) (interface{}, error) {
	return red.Int(RedisExec("ZCARD", key))
}

//返回有序集中，指定区间内的成员。 分数由低到高
//key   strart 起始  stop 结束   withscores:是否附带分值
func RedisZrange(key string, start, stop int32, withscores bool) ([]string, error) {
	if withscores {
		return red.Strings(RedisExec("ZRANGE", key, start, stop, "WITHSCORES"))
	}
	return red.Strings(RedisExec("ZRANGE", key, start, stop))
}

//返回有序集中，指定区间内的成员。 分数由高到低
//key   strart 起始  stop 结束   withscores:是否附带分值
func RedisZrevRange(key string, start, stop int32, withscores bool) ([]string, error) {
	if withscores {
		return red.Strings(RedisExec("zrevrange", key, start, stop, "WITHSCORES"))
	}
	return red.Strings(RedisExec("zrevrange", key, start, stop))
}

//按分值 返回有序集合中指定分数区间的成员列表。有序集成员按分数值递增(从小到大)次序排列。
//  min < score <= max
// withscores :是否附带分值
func RedisZrangeByScore(key string, min, max int32, withscores bool) ([]string, error) {
	if withscores {
		return red.Strings(RedisExec("ZRANGEBYSCORE", key, fmt.Sprintf("(%v", min), max, "WITHSCORES"))
	}
	return red.Strings(RedisExec("ZRANGEBYSCORE", key, fmt.Sprintf("(%v", min), max))
}

//按分值 返回有序集合中指定分数区间的成员列表。有序集成员按分数值递增(从大到小)次序排列。
//  min < score <= max
// withscores :是否附带分值
func RedisZrevrangebyscore(key string, max, min int32, withscores bool) ([]string, error) {
	if withscores {
		return red.Strings(RedisExec("Zrevrangebyscore", key, max, min, "WITHSCORES"))
	}
	return red.Strings(RedisExec("Zrevrangebyscore", key, max, min))
}

//用于移除有序集中，指定排名(rank)区间内的所有成员。 排名rank的分值 由低到高计算
func RedisZremeangeByRank(key string, start, stop int32) (int, error) {
	return red.Int(RedisExec("ZREMRANGEBYRANK", key, start, stop))
}

//移除有序集中，指定分数（score）区间内的所有成员。
func RedisZremrangebyScore(key string, min, max int32) (int, error) {
	return red.Int(RedisExec("ZREMRANGEBYSCORE", key, min, max))
}

//====================================================================================
