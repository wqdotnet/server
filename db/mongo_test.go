package db

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestConnectPool(t *testing.T) {
	InitMongodb("slggame", "mongodb://localhost:27017")
}

func TestInsertOne(t *testing.T) {
	InsertOne("cron_log", Testdata{
		Name: "Ash",
		Age:  18,
	})
	log.Info("TestInsertOne")
}

func TestFindObject(t *testing.T) {
	var obj Testdata
	FindOneObject("cron_log", "name", "Ash", &obj)
	log.Info("TestFindObject", obj)
}

func TestDelete(t *testing.T) {
	num := Delete("cron_log", "name", "Ash")
	log.Info("TestDelete num:", num)
}

type Testdata struct {
	Name string
	Age  int32
}
