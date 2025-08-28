package global

import (
	"bookstore-go/config"
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var mysqlClient *gorm.DB
var RedisClient *redis.Client

func InitMysql(mysqlConfig config.Mysql) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DatabaseName)

	dbClient, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(fmt.Errorf("init mysql client error:%v", err))
	}

	mysqlClient = dbClient
	log.Println("init mysql success")
}

func InitRedis(redisConfig config.Redis) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	str, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("init redis client error:", err)
	}
	fmt.Println("redis init success:", str)

	RedisClient = client
}

func GetDB() *gorm.DB {
	return mysqlClient
}

func CloseDB() {
	if mysqlClient != nil {
		sqlDB, err := mysqlClient.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
}
