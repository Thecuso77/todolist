package server

import (
	"database/sql"
	"log"
)

var MySqlDB *sql.DB

//var RedisDB *redis.Client

func InitMySqlDb() {
	var err error

	MySqlDB, err = sql.Open("mysql", "digitalprofile:root@tcp(goapp_db)/digitalprofile?parseTime=true")
	if err != nil {
		log.Panic(err)
	}

	if err = MySqlDB.Ping(); err != nil {
		log.Panic(err)
	}
	log.Printf("Connection to mysql successful")

	defer MySqlDB.Close()
}

//func InitRedisDb() {
//	var err error
//	ctx := context.Background()
//
//	RedisDB = redis.NewClient(&redis.Options{
//		Addr:     os.Getenv("REDIS_URL"),
//		Password: "",
//		DB:       0,
//	})
//
//	_, err = RedisDB.Ping(ctx).Result()
//
//	if err != nil {
//
//		log.Fatal("error", err)
//	}
//	log.Printf("Connection to redis successful")
//}
