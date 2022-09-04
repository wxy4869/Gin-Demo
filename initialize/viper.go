package initialize

import (
	"fmt"
	"ginDemo/global"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func InitViper() {
	rootPath, _ := os.Executable()
	rootPath = filepath.Dir(rootPath)
	v := viper.New()
	v.SetConfigFile("./config.yml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Viper 出问题啦: %s\n", err))
	}
	v.Set("root_path", "./")
	global.VP = v
}
