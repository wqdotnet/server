package db

import (
	"context"

	pool "github.com/jolestar/go-commons-pool/v2"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database   string
	urlstr     string
	clientPool *pool.ObjectPool
)

//InitMongodb mongodb init
func InitMongodb(dbname string, url string) {
	log.Info("init mongodb sync.pool")
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

//getDatabase connectPool mongodb database
func getDatabase() (*mongo.Client, *mongo.Database) {
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
	client, database := getDatabase()
	collection := database.Collection(collectionname)
	return client, collection
}

//InsertOne 添加数据
func InsertOne(tbname string, document interface{}) {
	client, collection := getCollection(tbname)

	insertResult, err := collection.InsertOne(context.TODO(), document)

	if err != nil {
		log.Fatal(err)
	}
	log.Info("Inserted a single document: ", insertResult)

	clientPool.ReturnObject(context.Background(), client)
}

//FindOneObject 查询数据
func FindOneObject(tbname string, field string, value interface{}, document interface{}) {
	client, collection := getCollection(tbname)
	filter := bson.D{{field, value}}

	err := collection.FindOne(context.TODO(), filter).Decode(document)

	if err != nil {
		log.Fatal("FindObject error:", err)
	}

	log.Debugf("Found a single document: %+v\n", document)
	clientPool.ReturnObject(context.Background(), client)
}

//Delete 删除
func Delete(tbname string, field string, value interface{}) int64 {
	client, collection := getCollection(tbname)
	filter := bson.D{{field, value}}
	//删除所有
	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	clientPool.ReturnObject(context.Background(), client)
	return deleteResult.DeletedCount
}
