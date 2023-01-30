package dao

import (
	"douyin/conf"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	SvDB    *gorm.DB
	SvRedis redis.Conn
)

// Dao represents data access object
/*type Dao struct {
	svDB    *gorm.DB
	svRedis *redis.Client
}*/

/*
	func New() *Dao {
		return &Dao{
			svDB:    InitDB(),
			svRedis: InitRedis(),
		}
	}
*/

func InitDB() {
	var err error
	SvDB, err = gorm.Open("mysql", conf.DBConnectString())
	if err != nil {
		fmt.Printf("mysql connect error = %v", err)
	}
	if SvDB.Error != nil {
		fmt.Printf("database error = %v", SvDB.Error)
	}
	//SvDB.LogMode(true)
}

func InitRedisClient() error {
	redisAddr := fmt.Sprintf("%s:%d", conf.Info.RDB.IP, conf.Info.RDB.Port)
	var err error
	SvRedis, err = redis.Dial("tcp", redisAddr,
		redis.DialPassword("123456"),
	)
	if err != nil {
		return err
	}
	return nil
}

func MysqlClose() {
	SvDB.Close()
}

func RedisClose() {
	SvRedis.Close()
}

type Model struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}
