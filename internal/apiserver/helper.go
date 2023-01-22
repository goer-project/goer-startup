package apiserver

import (
	"github.com/spf13/viper"

	"goer-startup/internal/apiserver/store"
	"goer-startup/internal/pkg/log"
	genericapiserver "goer-startup/internal/pkg/server"
	"goer-startup/pkg/db"
)

var cfgFile string

const (
	// defaultConfigName 指定了服务的默认配置文件名.
	defaultConfigName = "apiserver.yaml"
)

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	genericapiserver.LoadConfig(cfgFile, defaultConfigName)
}

// logOptions 从 viper 中读取日志配置，构建 `*log.Options` 并返回.
// 注意：`viper.Get<Type>()` 中 key 的名字需要使用 `.` 分割，以跟 YAML 中保持相同的缩进.
func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}

// initStore 读取 db 配置，创建 gorm.DB 实例，并初始化 store 层.
func initStore() error {
	dbOptions := &db.MySQLOptions{
		Host:                  viper.GetString("mysql.host"),
		Username:              viper.GetString("mysql.username"),
		Password:              viper.GetString("mysql.password"),
		Database:              viper.GetString("mysql.database"),
		MaxIdleConnections:    viper.GetInt("mysql.max-idle-connections"),
		MaxOpenConnections:    viper.GetInt("mysql.max-open-connections"),
		MaxConnectionLifeTime: viper.GetDuration("mysql.max-connection-life-time"),
		LogLevel:              viper.GetInt("mysql.log-level"),
	}

	ins, err := db.NewMySQL(dbOptions)
	if err != nil {
		return err
	}

	_ = store.NewStore(ins)

	return nil
}
