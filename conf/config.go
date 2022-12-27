package conf

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Conf 全局变量，用来保存配置的所有信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	//*SnowFlakeConfig `mapstructure:"snowflake"`
}

type LogConfig struct {
	Level      string `mapstructure:"Level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`    //单个文件最大内存大小
	MaxAge     int    `mapstructure:"max_age"`     //每 max_age 天切割一次
	MaxBackUps int    `mapstructure:"max_backups"` //多于 max_backups 个日志文件后，清理较旧的日志
	Mode       string `mapstructure:"log_mode"`    //模式 dev 开发模式
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`     //主机名
	User         string `mapstructure:"user"`     //用户名
	Password     string `mapstructure:"password"` //登录密码
	DbName       string `mapstructure:"dbname"`   //数据库名
	Port         int    `mapstructure:"port"`     //端口号
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`     //主机名
	Password string `mapstructure:"password"` //登录密码
	Port     string `mapstructure:"port"`     //端口号
	DB       string `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init(fileName string) (err error) {
	// 方法1：直接指定配置文件路径（相对路径或者绝对路径）
	// 相对路径：相对执行的可执行文件的相对路径
	//viper.SetConfigFile("./config/config.yaml")
	// 绝对路径：系统中实际的文件路径
	//viper.SetConfigFile("/Users/admin/Projects/douyin/config/config.yaml")

	// 方法二：指定配置文件名和配置文件的位置，viper自行查找可用的配置文件
	// 配置文件名不需要带后缀
	// 配置文件位置可配置多个
	//viper.SetConfigName("config") // 指定配置文件名称（不需要带后缀）
	//viper.AddConfigPath(".")      // 指定查找配置文件的路径（这里指相对路径）

	//直接指定配置文件路径
	viper.SetConfigFile(fileName)
	// 读取配置信息
	err = viper.ReadInConfig()
	if err != nil {
		// 读取配置文件失败
		fmt.Printf("read config file failed err = %v\n", err)
		return err
	}
	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err = %v\n", err)
		return err
	}
	// listen the Config
	viper.WatchConfig()
	// 配置文件发生修改
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
		// 把读取到的配置信息反序列化到 Conf 变量中
		if err := viper.Unmarshal(Conf); err != nil {
			zap.L().Error("viper.Unmarshal failed err = ", zap.Error(err))
			return
		}
	})
	return
}
