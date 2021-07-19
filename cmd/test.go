package cmd

import (
	"context"
	"fmt"
	"time"

	pool "github.com/jolestar/go-commons-pool/v2"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type teststr struct {
	name string
	id   int
}

func exectest(cmd *cobra.Command, args []string) {

	parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	startTime, _ := parser.Parse("0 0 0 8 6 ? ")
	endTime, _ := parser.Parse("0 0 0 15 6 ? ")

	fmt.Printf("%v   %v ", startTime.Next(time.Now()), endTime.Next(time.Now()))
	fmt.Println()
	stime := startTime.Next(time.Now())
	etime := endTime.Next(time.Now())
	fmt.Printf("%v   %v ", startTime.Next(time.Now()).Unix(), endTime.Next(time.Now()).Unix())
	fmt.Println()
	fmt.Println(stime.Unix() > etime.Unix())

	//slice()
	//time.Sleep(time.Second * 10)
	//objectPool()

	Record := make(map[uint32]uint32)
	Record[2] = 34

}

func slice() {
	var ss []string
	fmt.Printf("[ local print ]\t:\t length:%v\taddr:%p\tisnil:%v\n", len(ss), ss, ss == nil)
	fmt.Println("func print", ss)
	//切片尾部追加元素append elemnt
	for i := 0; i < 10; i++ {
		ss = append(ss, fmt.Sprintf("s%d", i))
	}
	fmt.Printf("[ local print ]\t:\tlength:%v\taddr:%p\tisnil:%v\n", len(ss), ss, ss == nil)
	fmt.Println("after append", ss)
	//删除切片元素remove element at index
	index := 5
	ss = append(ss[:index], ss[index+1:]...)
	fmt.Println("after delete", ss)
	//在切片中间插入元素insert element at index;
	//注意：保存后部剩余元素，必须新建一个临时切片
	rear := append([]string{}, ss[index:]...)
	ss = append(ss[0:index], "inserted")
	ss = append(ss, rear...)
	fmt.Println("after insert\n", ss)

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
	database := client.Database("gamedemo")

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
	filter := bson.D{primitive.E{Key: "name", Value: "Ash"}}
	update := bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "age", Value: 1}}}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	//查找
	filter2 := bson.D{primitive.E{Key: "name", Value: "Ash"}}
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
