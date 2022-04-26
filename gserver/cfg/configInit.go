package cfg

import (

	// "github.com/google/wire"

	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitViperConfig 初始化viper
func InitViperConfig(cfgPath string, cfgType string) *viper.Viper {
	logrus.Infof("loanding config [%s]  [%s]", cfgPath, cfgType)

	v := viper.New()
	v.AddConfigPath(cfgPath)
	v.SetConfigType(cfgType)

	cfg := &cfgCollection{}
	reflectField(cfg, cfgPath, cfgType, v)
	saveCfg(cfg)
	return v
}

func reflectField(structName interface{}, cfgPath, cfgType string, v *viper.Viper) {
	t := reflect.ValueOf(structName)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		logrus.Fatal("Check type error not Struct")
		return
	}

	fieldNum := t.NumField()

	for i := 0; i < fieldNum; i++ {
		fieldname := t.Type().Field(i).Name
		typename := t.Field(i).Type().Name()
		field := t.Field(i).Interface()

		logrus.Info("load init config =>:", fieldname)
		v.SetConfigName(fieldname)

		if err := v.ReadInConfig(); err != nil {
			//viper 库无法加载 "[{}]" 格式json
			if cfgType == "json" {
				jsonFile, e1 := os.Open(cfgPath + "/" + fieldname + "." + cfgType)
				defer jsonFile.Close()
				if e1 != nil {
					logrus.Fatalf("fiel: [%v] err:[%v]", jsonFile, e1)
				}
				jsda, err := ioutil.ReadAll(jsonFile)
				if err != nil {
					logrus.Fatalf("ReadAll: [%v] [%v][%v]", err, typename, field)
				}

				newdata := reflect.New(reflect.TypeOf(field)).Interface()
				if err := json.Unmarshal(jsda, newdata); err != nil {
					logrus.Fatalf("unmarshal: [%v] [%v][%v]", err, typename, field)
				}

				t.FieldByName(fieldname).Set(reflect.ValueOf(newdata).Elem())
				continue
			}
			logrus.Fatalf("err:  [%v]   ", err)
		}

		if err := v.UnmarshalExact(&field); err != nil {
			logrus.Fatalf("err:  [%v]   [%v] ", err, field)
		}

		t.FieldByName(fieldname).Set(reflect.ValueOf(field))
	}
}

func WatchConfig(configDir string, run func(in fsnotify.Event)) {
	initWG := sync.WaitGroup{}
	initWG.Add(1)
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			logrus.Fatal(err)
		}
		defer watcher.Close()
		// we have to watch the entire directory to pick up renames/atomic saves in a cross-platform way

		if err != nil {
			logrus.Printf("error: %v\n", err)
			initWG.Done()
			return
		}

		eventsWG := sync.WaitGroup{}
		eventsWG.Add(1)
		go func() {
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok { // 'Events' channel is closed
						eventsWG.Done()
						return
					}
					// we only care about the config file with the following cases:
					// 1 - if the config file was modified or created
					// 2 - if the real path to the config file changed (eg: k8s ConfigMap replacement)
					const writeOrCreateMask = fsnotify.Write | fsnotify.Create
					if event.Op&writeOrCreateMask != 0 {
						if run != nil {
							run(event)
						}
					}

				case err, ok := <-watcher.Errors:
					if ok { // 'Errors' channel is not closed
						logrus.Printf("watcher error: %v\n", err)
					}
					eventsWG.Done()
					return
				}
			}
		}()
		watcher.Add(configDir)
		initWG.Done()   // done initializing the watch in this go routine, so the parent routine can move on...
		eventsWG.Wait() // now, wait for event loop to end in this go-routine...
	}()
	initWG.Wait() // make sure that the go routine above fully ended before returning
}
