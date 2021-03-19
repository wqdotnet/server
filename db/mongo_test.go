package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	//"github.com/go-playground/assert/v2"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func TestConnectPool(t *testing.T) {
	StartMongodb("slggame", "mongodb://localhost:27017")
}

func TestInsertOne(t *testing.T) {
	InsertOne("cron_log", Testdata{
		Name: "Ash",
		Age:  17,
	})

	data := &Testdata{Name: "wq", Age: 18}
	InsertOne("cron_log", &data)
	log.Info("TestInsertOne")
}

func TestFindFieldMax(t *testing.T) {
	var obj Testdata
	FindFieldMax("cron_log", "age", &obj)
	log.Info("TestFindFieldMax:", obj.Age)
	assert.Equal(t, obj.Age, int32(18))
}

func TestFindBson(t *testing.T) {
	var obj Testdata
	filter := bson.D{{"name", "wq"}, {"age", 18}}
	FindOneBson("cron_log", &obj, filter)
	log.Info("TestFindObject", obj)

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
	log.Info("TestFindObject", obj)
}

func TestFind(t *testing.T) {
	filter := bson.D{{"age", 18}}

	var results []*Testdata
	cur, err := FindBson("cron_log", filter)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem Testdata
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	// Close the cursor once finished
	cur.Close(context.TODO())

	log.Info(results)
}

func TestDelete(t *testing.T) {
	num := Delete("cron_log", "name", "Ash")
	log.Info("TestDelete num:", num)
	num = Delete("cron_log", "name", "wq")
	log.Info("TestDelete num:", num)
}

type Testdata struct {
	Name string
	Age  int32
}
