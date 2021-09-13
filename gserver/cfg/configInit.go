package cfg

import (

	// "github.com/google/wire"

	"reflect"
	"sync"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitViperConfig 初始化viper
func InitViperConfig(cfgPath string, cfgType string) *viper.Viper {
	log.Infof("loanding config [%s]  [%s]", cfgPath, cfgType)

	v := viper.New()
	v.AddConfigPath(cfgPath)
	v.SetConfigType(cfgType)

	cfg := &cfgCollection{}
	reflectField(cfg, v)
	saveCfg(cfg)
	return v
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

		log.Info("load init config :", fieldname)
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

func WatchConfig(configDir string, run func(in fsnotify.Event)) {
	initWG := sync.WaitGroup{}
	initWG.Add(1)
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()
		// we have to watch the entire directory to pick up renames/atomic saves in a cross-platform way

		if err != nil {
			log.Printf("error: %v\n", err)
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
						log.Printf("watcher error: %v\n", err)
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
