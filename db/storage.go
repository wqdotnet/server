package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetInfo 获取数据
func GetInfo(tabname string, field string, value interface{}, document interface{}) error {
	//读取缓存

	filter := bson.D{primitive.E{Key: field, Value: value}}
	return FindOneBson(tabname, document, filter)
}
