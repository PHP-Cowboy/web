package initialize

import (
	"web/global"

	"github.com/spf13/viper"
)

func InitConfig() {
	v := viper.New()

	v.SetConfigFile("config.json")

	err := v.ReadInConfig()
	if err != nil {
		panic("读取配置文件失败:" + err.Error())
	}

	//fmt.Println(content) //字符串 - yaml
	//想要将一个json字符串转换成struct，需要去设置这个struct的tag
	err = v.Unmarshal(global.ServerConfig)
	if err != nil {
		panic("解析配置失败:" + err.Error())
	}
	//fmt.Printf("config: %+v", global.ServerConfig)
}
