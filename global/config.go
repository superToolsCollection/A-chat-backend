package global

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

/**
* @Author: super
* @Date: 2020-09-23 08:36
* @Description:
**/

var (
	SensitiveWords []string

	MessageQueueLen = 1024
)

func initConfig() {
	viper.SetConfigName("chatroom")
	viper.AddConfigPath(RootDir + "/config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	SensitiveWords = viper.GetStringSlice("sensitive")
	MessageQueueLen = viper.GetInt("message-queue")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		viper.ReadInConfig()

		SensitiveWords = viper.GetStringSlice("sensitive")
	})
}
