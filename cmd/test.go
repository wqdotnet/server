package cmd

import (
	"context"
	"fmt"

	pool "github.com/jolestar/go-commons-pool/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test demo",
	Long:  `test demo`,
	Run:   exectest,
}

func init() {
	rootCmd.AddCommand(testCmd)
}

func exectest(cmd *cobra.Command, args []string) {

	mongodb()
	//objectPool()

}

func mongodb() {
	type trainer struct {
		Name string
		Age  int
		City string
	}

	var (
		client *mongo.Client
		err    error
	)

	// 建立mongodb连接
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	if client, err = mongo.Connect(context.TODO(), clientOptions); err != nil {
		log.Error(err)
		return
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Connected to MongoDB!")

	// 2, 选择数据库my_db
	database := client.Database("slggame")

	// 3, 选择表my_collection
	collection := database.Collection("cron_log")
	// 4, 插入记录(bson)
	ash := trainer{"Ash", 10, "Pallet Town"}
	misty := trainer{"Misty", 10, "Cerulean City"}
	brock := trainer{"Brock", 15, "Pewter City"}
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Inserted a single document: ", insertResult)

	//插入列表数据
	trainers := []interface{}{misty, brock}
	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	// 更新
	filter := bson.D{{"name", "Ash"}}
	update := bson.D{{"$inc", bson.D{{"age", 1}}}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	//查找
	filter2 := bson.D{{"name", "Ash"}}
	var result trainer
	err = collection.FindOne(context.TODO(), filter2).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Found a single document: %+v\n", result)

	//删除所有
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

func objectPool() {
	factory := pool.NewPooledObjectFactorySimple(
		func(context.Context) (interface{}, error) {
			clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
			client, err := mongo.Connect(context.TODO(), clientOptions)
			if err != nil {
				log.Error(err)
			}
			return client, nil
		})

	ctx := context.Background()
	p := pool.NewObjectPoolWithDefaultConfig(ctx, factory)

	obj, err := p.BorrowObject(ctx)
	if err != nil {
		log.Error(err)
	}

	client := obj.(*mongo.Client)
	err = client.Ping(context.TODO(), nil)
	fmt.Println(err)

	err = p.ReturnObject(ctx, obj)
	if err != nil {
		log.Error(err)
	}
}
