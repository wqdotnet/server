package db

import (
	"context"
	"fmt"
	"testing"

	//"github.com/go-playground/assert/v2"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestConnectPool(t *testing.T) {
	StartMongodb("gamedemo", "mongodb://admin:123456@localhost:27017")
	b, e := MongodbPing()
	fmt.Println(b, e)
}

func TestInsertOne(t *testing.T) {
	InsertOne("cron_log", Testdata{
		Name: "Ash",
		Age:  17,
	})

	data := &Testdata{Name: "wq", Age: 18}
	InsertOne("cron_log", &data)
	logrus.Info("TestInsertOne")
}

func TestFindFieldMax(t *testing.T) {
	var obj Testdata
	FindFieldMax("cron_log", "age", &obj)
	logrus.Info("TestFindFieldMax:", obj.Age)
	assert.Equal(t, obj.Age, int32(18))
}

func TestFindBson(t *testing.T) {
	var obj Testdata
	filter := bson.D{{"name", "wq"}, {"age", 18}}
	FindOneBson("cron_log", &obj, filter)
	logrus.Info("TestFindObject", obj)

}

func TestUpdate(t *testing.T) {
	filter := bson.D{{"name", "Ash"}}
	// $inc 加减
	updatefilter := bson.D{{"$set", bson.D{{"age", 18}}}}
	Update("cron_log", filter, updatefilter)

}

func TestFindOne(t *testing.T) {
	var obj Testdata
	list := make(map[string]interface{})
	list["name"] = "Ash"
	//list["age"] = 18
	FindOneBson("cron_log", &obj, list)
	logrus.Info("TestFindObject", obj)
}

func TestFind(t *testing.T) {
	filter := bson.D{{"age", 18}}

	var results []*Testdata
	cur, err := FindBson("cron_log", filter)
	if err != nil {
		logrus.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem Testdata
		err := cur.Decode(&elem)
		if err != nil {
			logrus.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		logrus.Fatal(err)
	}
	// Close the cursor once finished
	cur.Close(context.TODO())

	logrus.Info(results)
}

func TestDelete(t *testing.T) {
	num := Delete("cron_log", "name", "Ash")
	logrus.Info("TestDelete num:", num)
	num = Delete("cron_log", "name", "wq")
	logrus.Info("TestDelete num:", num)
}

type Testdata struct {
	Name string
	Age  int32
}

// func mongodb() {
// 	type trainer struct {
// 		Name string
// 		Age  int
// 		City string
// 	}

// 	var (
// 		client *mongo.Client
// 		err    error
// 	)

// 	// 建立mongodb连接
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	if client, err = mongo.Connect(context.TODO(), clientOptions); err != nil {
// 		logrus.Error(err)
// 		return
// 	}

// 	// 检查连接
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		logrus.Fatal(err)
// 	}
// 	logrus.Info("Connected to MongoDB!")

// 	// 2, 选择数据库my_db
// 	database := client.Database("gamedemo")

// 	// 3, 选择表my_collection
// 	collection := database.Collection("cron_log")
// 	// 4, 插入记录(bson)
// 	ash := trainer{"Ash", 10, "Pallet Town"}
// 	misty := trainer{"Misty", 10, "Cerulean City"}
// 	brock := trainer{"Brock", 15, "Pewter City"}
// 	insertResult, err := collection.InsertOne(context.TODO(), ash)
// 	if err != nil {
// 		logrus.Fatal(err)
// 	}
// 	logrus.Info("Inserted a single document: ", insertResult)

// 	//插入列表数据
// 	trainers := []interface{}{misty, brock}
// 	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
// 	if err != nil {
// 		logrus.Fatal(err)
// 	}
// 	logrus.Info("Inserted multiple documents: ", insertManyResult.InsertedIDs)

// 	// 更新
// 	filter := bson.D{primitive.E{Key: "name", Value: "Ash"}}
// 	update := bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "age", Value: 1}}}}
// 	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		logrus.Fatal(err)
// 	}
// 	logrus.Infof("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

// 	//查找
// 	filter2 := bson.D{primitive.E{Key: "name", Value: "Ash"}}
// 	var result trainer
// 	err = collection.FindOne(context.TODO(), filter2).Decode(&result)
// 	if err != nil {
// 		logrus.Fatal(err)
// 	}
// 	logrus.Infof("Found a single document: %+v\n", result)

// 	//删除所有
// 	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
// 	if err != nil {
// 		logrus.Fatal(err)
// 	}
// 	logrus.Infof("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
// }

// func objectPool() {
// 	factory := pool.NewPooledObjectFactorySimple(
// 		func(context.Context) (interface{}, error) {
// 			clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 			client, err := mongo.Connect(context.TODO(), clientOptions)
// 			if err != nil {
// 				logrus.Error(err)
// 			}
// 			return client, nil
// 		})

// 	ctx := context.Background()
// 	p := pool.NewObjectPoolWithDefaultConfig(ctx, factory)

// 	obj, err := p.BorrowObject(ctx)
// 	if err != nil {
// 		logrus.Error(err)
// 	}

// 	client := obj.(*mongo.Client)
// 	err = client.Ping(context.TODO(), nil)
// 	fmt.Println(err)

// 	err = p.ReturnObject(ctx, obj)
// 	if err != nil {
// 		logrus.Error(err)
// 	}
// }
