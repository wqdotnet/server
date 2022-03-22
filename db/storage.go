package db

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetStorageInfo 获取数据
func GetStorageInfo(tabname string, field string, value interface{}, document interface{}) error {
	// //读取缓存
	// if err := GetStruct(fmt.Sprintf("%v_%v", tabname, value), document); err != nil && document != nil {
	// 	return nil
	// }

	filter := bson.D{primitive.E{Key: field, Value: value}}
	return FindOneBson(tabname, document, filter)
}

//SaveStorageInfo save
func SaveStorageInfo(tabname string, key interface{}, document interface{}) {
	SetStruct(fmt.Sprintf("%v_%v", tabname, key), document)
}
