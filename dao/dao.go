package dao

import (
	"douyin/conf"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var (
	SvDB    *gorm.DB
	SvRedis *redis.Client
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

func InitRedisClient() (err error) {
	SvRedis = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       10,
	})
	_, err = SvRedis.Ping().Result()
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
