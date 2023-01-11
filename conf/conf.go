package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"strings"
)

type Mysql struct {
	Host        string
	Port        int
	Database    string
	Username    string
	Password    string
	Charset     string
	ParseTime   bool `toml:"parse_time"`
	Loc         string
	ReadTimeout string `toml:"read_timeout"`
	MaxIdel     int    `toml:"max_idel"`
	MaxActive   int    `toml:"max_active"`
}

type Redis struct {
	IP       string
	Port     int
	Database int
}

type Server struct {
	IP   string
	Port int
}

// Path ffmpeg_path和static_path
type Path struct {
	FfmpegPath       string `toml:"ffmpeg_path"`
	StaticSourcePath string `toml:"static_source_path"`
}

type Config struct {
	DB     Mysql `toml:"database"`
	RDB    Redis `toml:"redis"`
	Server `toml:"server"`
	Path   `toml:"path"`
}

// Info 全局参数配置
var Info Config

func init() {
	// toml加载配置文件xxx.toml
	if _, err := toml.DecodeFile("/Users/inke226101m/workspace/src/github.com/longsuizhi/douyin/app/config/config.toml", &Info); err != nil {
		panic(err)
	}
	//去除左右的空格
	Info.Server.IP = strings.Trim(Info.Server.IP, " ")
	Info.RDB.IP = strings.Trim(Info.RDB.IP, " ")
	Info.DB.Host = strings.Trim(Info.DB.Host, " ")
}

// DBConnectString 填充得到数据库连接字符串
func DBConnectString() string {
	arg := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s&readTimeout=%s",
		Info.DB.Username, Info.DB.Password, Info.DB.Host, Info.DB.Port, Info.DB.Database,
		Info.DB.Charset, Info.DB.ParseTime, Info.DB.Loc, Info.DB.ReadTimeout)
	log.Println(arg)
	return arg
}
