package db

import (
	"context"

	pool "github.com/jolestar/go-commons-pool/v2"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database   string
	urlstr     string
	clientPool *pool.ObjectPool
)

//StartMongodb mongodb init
func StartMongodb(dbname string, url string) {
	log.Infof("StartMongodb  create sync.pool:[%v]   dbname:[%v]", url, dbname)
	database = dbname
	urlstr = url

	factory := pool.NewPooledObjectFactorySimple(
		func(context.Context) (interface{}, error) {
			clientOptions := options.Client().ApplyURI(url)
			client, err := mongo.Connect(context.TODO(), clientOptions)
			if err != nil {
				log.Error(err)
			}
			return client, nil
		})
	ctx := context.Background()
	clientPool = pool.NewObjectPoolWithDefaultConfig(ctx, factory)
}

func MongodbPing() (bool, error) {
	ctx := context.Background()
	obj, err := clientPool.BorrowObject(ctx)
	if err != nil {
		log.Error(err)
	}
	client := obj.(*mongo.Client)
	if err := client.Ping(context.TODO(), nil); err != nil {
		return false, err
	}
	return true, nil
}

//GetDatabase connectPool mongodb database
func GetDatabase() (*mongo.Client, *mongo.Database) {
	ctx := context.Background()
	obj, err := clientPool.BorrowObject(ctx)
	if err != nil {
		log.Error(err)
	}
	client := obj.(*mongo.Client)

	// 检查连接
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Warn(err)
		clientOptions := options.Client().ApplyURI(urlstr)
		if client, err = mongo.Connect(context.TODO(), clientOptions); err != nil {
			log.Error(err)
		}
	}

	database := client.Database(database)
	return client, database
}

//getCollection connectPool mongodb collection
func getCollection(collectionname string) (*mongo.Client, *mongo.Collection) {
	client, database := GetDatabase()
	collection := database.Collection(collectionname)
	return client, collection
}

//InsertOne 添加数据
func InsertOne(tbname string, document interface{}) {
	client, collection := getCollection(tbname)

	_, err := collection.InsertOne(context.TODO(), document)

	if err != nil {
		log.Error(err)
	}

	clientPool.ReturnObject(context.Background(), client)

}

//FindOneBson 查询数据
//filter := bson.D{{field, value}}
//filter := bson.D{primitive.E{Key:field, value}}
func FindOneBson(tbname string, document interface{}, filter interface{}) error {
	client, collection := getCollection(tbname)
	defer clientPool.ReturnObject(context.Background(), client)
	return collection.FindOne(context.TODO(), filter).Decode(document)
}

//FindBson 查找数据
func FindBson(tbname string, filter interface{}) (*mongo.Cursor, error) {
	client, collection := getCollection(tbname)
	defer clientPool.ReturnObject(context.Background(), client)

	findOptions := options.Find()

	return collection.Find(context.TODO(), filter, findOptions)
	// cur, err := collection.Find(context.TODO(), filter, findOptions)
	// if err != nil {
	// 	return 0, err
	// }

	// // Finding multiple documents returns a cursor
	// // Iterating through the cursor allows us to decode documents one at a time
	// for cur.Next(context.TODO()) {

	// 	// create a value into which the single document can be decoded
	// 	var elem Trainer
	// 	err := cur.Decode(&amp;elem)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	results = append(results, &amp;elem)
	// }

	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// // Close the cursor once finished
	// cur.Close(context.TODO())

	//return len(cur.Current), nil
}

//Update 更新数据
//	Findfield := bson.D{{"name", "Ash"}}
//	Upfield := bson.D{{"$inc", bson.D{{"age", 1}}}}
func Update(tbname string, Findfield interface{}, Upfield interface{}) (int64, error) {
	client, collection := getCollection(tbname)
	defer clientPool.ReturnObject(context.Background(), client)

	updateResult, err := collection.UpdateOne(context.TODO(), Findfield, Upfield)
	if err != nil {
		return 0, err
	}
	//log.Debug("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return updateResult.ModifiedCount, nil
}

//Delete 删除
func Delete(tbname string, field string, value interface{}) int64 {
	client, collection := getCollection(tbname)
	defer clientPool.ReturnObject(context.Background(), client)
	filter := bson.D{primitive.E{Key: field, Value: value}}
	//删除所有
	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Error(err)
	}
	//log.Debugf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	return deleteResult.DeletedCount
}

//FindFieldMax 查询最大值
func FindFieldMax(tbname string, fieldkey string, document interface{}) error {
	client, collection := getCollection(tbname)
	defer clientPool.ReturnObject(context.Background(), client)

	filter := bson.D{{}}
	findOptions := options.FindOne().SetSort(bson.D{primitive.E{Key: fieldkey, Value: -1}}).SetSkip(0)
	return collection.FindOne(context.TODO(), filter, findOptions).Decode(document)
}
