package db

import (
	"time"

	red "github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

//Redis redis
type Redis struct {
	pool *red.Pool
}

var redis *Redis

var auitid int32

//StartRedis 初始化
func StartRedis(address string) {
	log.Infof("StartRedis  create redis.pool:  [%v]", address)
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
				red.DialDatabase(0),
				//red.DialPassword(""),
			)
		},
	}
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
		for _, v := range args {
			parmas = append(parmas, v)
		}
	}
	return con.Do(cmd, parmas...)
}

//GetAutoID 获取自增id
func GetAutoID(tabname string) int32 {
	autoid, err := red.Int(RedisExec("incr", tabname))
	if err != nil {
		log.Error(err)
	}
	return int32(autoid)
}
