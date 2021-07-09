/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"reflect"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// reflexCmd represents the reflex command
var reflexCmd = &cobra.Command{
	Use:   "reflex",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		reflex()
	},
}

func init() {
	rootCmd.AddCommand(reflexCmd)
}

var execfuncname reflect.Value

func sssssdc() {

	execfuncname.Call(nil)
}

func testfun() {

}

func reflex() {
	execfuncname = reflect.ValueOf(testfun)
	t1 := time.Now()
	for i := 0; i < 10000; i++ {
		sssssdc()
	}
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)

	//funtype:= reflect.ValueOf(f).Type(),

}

func reflectField(structName interface{}, v *viper.Viper) {
	t := reflect.ValueOf(structName)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Fatal("Check type error not Struct")
		return
	}

	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		fieldname := t.Type().Field(i).Name
		typename := t.Field(i).Type().Name()
		field := t.Field(i).Interface()

		log.Info("init config :", fieldname)
		v.SetConfigName(fieldname)

		if err := v.ReadInConfig(); err != nil {
			log.Fatalf("read: %v [%v][%v]", err, typename, fieldname)
		}

		if err := v.UnmarshalExact(&field); err != nil {
			log.Fatal("err:", err)
		}
		t.FieldByName(fieldname).Set(reflect.ValueOf(field))
	}
}
