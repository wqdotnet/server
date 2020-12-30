package cfg

import (

	// "github.com/google/wire"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitViperConfig 初始化viper
func InitViperConfig(CfgPath string, CfgType string) {
	log.Infof("loanding config [%s]  [%s]", CfgPath, CfgType)

	v := viper.New()
	v.AddConfigPath(CfgPath)
	v.SetConfigType(CfgType)

	v.SetConfigName("mapinfo")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.UnmarshalExact(&GlobalCfg.MapInfo); err != nil {
		panic(err)
	}

	//reflectField(&GlobalCfg)

	v.WatchConfig()
	v.OnConfigChange(fileChanged)

}

func fileChanged(e fsnotify.Event) {
	log.Info("Config file changed:", e.Name)
}

// func reflectField(structName interface{}) {
// 	t := reflect.ValueOf(structName)
// 	//v := reflect.ValueOf(structName)
// 	if t.Kind() == reflect.Ptr {
// 		t = t.Elem()
// 	}
// 	if t.Kind() != reflect.Struct {
// 		log.Println("Check type error not Struct")
// 		return
// 	}

// 	fieldNum := t.NumField()
// 	for i := 0; i < fieldNum; i++ {

// 		filename := t.Field(i).Type().Name()
// 		field := t.Field(i)
// 		// filetype := t.Field(i).Type()
// 		// file := reflect.New(filetype)

// 		field.FieldByName("")

// 		viper.SetConfigName(filename)

// 		if err := viper.ReadInConfig(); err != nil {
// 			panic(err)
// 		}

// 		//bs, err := json.Marshal(v.AllSettings())
// 		err := viper.UnmarshalExact(&field)
// 		log.Debug(field)
// 		log.Error(err)

// 	}

// }
