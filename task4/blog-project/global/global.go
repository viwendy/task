package global

import (
	"blog-main/config"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DBClient *gorm.DB
var RedisClient *redis.Client

func InitMysql() {
	mysqlConfig := config.AppConfig.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("gorm连接mysql数据库失败：", err)
	}
	DBClient = db
	log.Println("mysql数据库连接成功")
}

func InitRedis() {
	redisConfig := config.AppConfig.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	RedisClient = rdb
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("redis连接失败：", err)
	}
	log.Println("ping:", pong)
	log.Println("redis连接成功")
}

func GetDB() *gorm.DB {
	return DBClient
}

func CloseDB() {
	if DBClient != nil {
		sqlDB, err := DBClient.DB()
		if err != nil {
			sqlDB.Close()
		}
	}
}

func CloseRedis() {
	if RedisClient != nil {
		err := RedisClient.Close()
		if err != nil {
			log.Println("redis关闭失败：", err)
		}
		log.Println("redis关闭成功")
	}
}

func ReturnJson(ctx *gin.Context, httpStatus, code int, msg string, data interface{}) {
	response := gin.H{
		"code":    code,
		"message": msg,
		"data":    data,
	}
	ctx.JSON(httpStatus, response)
}
