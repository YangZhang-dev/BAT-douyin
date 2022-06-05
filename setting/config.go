package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 通过全局变量Conf来存储配置信息
var Conf = new(Config)

type Config struct {
	*AppConfig   `mapstructure:"app"`
	*MySQLConfig `mapstructure:"mysql"`
	*PprofConfig `mapstructure:"pprof"`
}

// AppConfig 项目配置文件
type AppConfig struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Ip   string `mapstructure:"ip"`
	Port int64  `mapstructure:"port"`
}

// MySQLConfig mysql配置文件
type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"dbname"`
	Port     int    `mapstructure:"port"`
}

// PprofConfig pprof配置文件
type PprofConfig struct {
	Port int64  `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

func Init(settingPath string) (err error) {
	//导入viper来读取配置文件

	//设置文件类型和文件路径
	viper.SetConfigType("yaml")
	viper.SetConfigFile(settingPath)
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return err
	}

	//将配置信息反序列化到结构体上
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		return err
	}

	//监听配置文件
	viper.WatchConfig()
	//当配置文件被被更改时
	viper.OnConfigChange(func(in fsnotify.Event) {
		//被更改后执行的回调函数
		fmt.Println("config info be changed...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Println("viper.Unmarshal failed")
			return
		}
	})

	fmt.Printf("reading config success...\ninfo：%v\nAppConfig：%v\n", Conf, Conf.AppConfig)
	return nil
}
