package cfg

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func init() {
	InitViperConfig("../../config", "json")

	//viper.AddConfigPath("./config")
	//viper.SetConfigName("mapinfo")

	log.Info("MapInfo :", len(GlobalCfg.MapInfo.Areas))
}

func TestMapCfg(t *testing.T) {

	// assert.Equal(t, CheckBigMapConfig(), true)
	// assert.Equal(t, AreasIsBeside(2427, 3449), true)
	// assert.Equal(t, AreasIsBeside(2427, 2725), true)
	// assert.Equal(t, AreasIsBeside(2427, 3966), false)
}
