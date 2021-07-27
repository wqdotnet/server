package cfg

import (

	// "github.com/google/wire"

	"reflect"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitViperConfig 初始化viper
func InitViperConfig(CfgPath string, CfgType string) {
	log.Infof("loanding config [%s]  [%s]", CfgPath, CfgType)

	v := viper.New()
	v.AddConfigPath(CfgPath)
	v.SetConfigType(CfgType)

	reflectField(&GameCfg, v)

	// v.WatchConfig()
	// v.OnConfigChange(fileChanged)
}

// func fileChanged(e fsnotify.Event) {
// 	log.Info("Config file changed:", e.Name)
// }

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
