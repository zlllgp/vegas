package viper

import (
	"github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/fsnotify/fsnotify"
	"github.com/kr/pretty"
	"github.com/spf13/viper"
	"github.com/zlllgp/vegas/internal/application/config"
	"os"
	"path/filepath"
	"sync"
)

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join("./conf", GetEnv()))           //working dir
	viper.AddConfigPath(filepath.Join("/home/admin/conf", GetEnv())) //server path
	if err := viper.ReadInConfig(); err != nil {
		klog.Errorf("viper.ReadInConfig error: %v", err)
		panic(err)
	}

	if err := viper.Unmarshal(&config.GlobalConfig); err != nil {
		klog.Errorf("Viper: viper.Unmarshal error: %v", err)
		panic(err)
	}
	if err := validator.Validate(config.GlobalConfig); err != nil {
		klog.Error("Viper: validate config error - %v", err)
		panic(err)
	}
	config.GlobalConfig.Env = GetEnv()
	klog.Infof("Viper: init success")
	pretty.Printf("Viper: %+v\n", config.GlobalConfig)

	//watch
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		klog.Infof("Viper: Config file changed: %s", e.Name)
		if err := viper.Unmarshal(&config.GlobalConfig); err != nil {
			klog.Errorf("Viper: viper.Unmarshal error: %v", err)
		}
	})
}

var (
	//conf *config.Config
	once sync.Once
)

// GetConf gets configuration instance
func GetConf() *config.Config {
	once.Do(Init)
	return &config.GlobalConfig
}

func GetEnv() string {
	e := os.Getenv("GO_ENV")
	if len(e) == 0 {
		return "test"
	}
	return e
}

func LogLevel() klog.Level {
	level := config.GlobalConfig.Kitex.LogLevel
	switch level {
	case "trace":
		return klog.LevelTrace
	case "debug":
		return klog.LevelDebug
	case "info":
		return klog.LevelInfo
	case "notice":
		return klog.LevelNotice
	case "warn":
		return klog.LevelWarn
	case "error":
		return klog.LevelError
	case "fatal":
		return klog.LevelFatal
	default:
		return klog.LevelInfo
	}
}
