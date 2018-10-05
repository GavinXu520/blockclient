package common


import (
	"github.com/spf13/viper"
	"gopkg.in/redis.v5"
	"time"
)

func SetUpRedis() {

	addr := viper.GetString("redis.addr")
	password := viper.GetString("redis.password")
	database := viper.GetInt("redis.database")
	maxActive := viper.GetInt("redis.maxActive")
	idleTimeout := time.Duration(viper.GetInt("redis.idleTimeout")) * time.Second

	client := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    password,
		DB:          database,
		MaxRetries:  3,
		IdleTimeout: idleTimeout,
		PoolSize:    maxActive,

	})

	_, err := client.Ping().Result()
	if err != nil {
		panic("failed to connect redis:" + err.Error())
	}
	Redis = client
}

var (
	Redis *redis.Client
)
