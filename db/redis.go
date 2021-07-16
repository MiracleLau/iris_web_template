package db

import (
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"sync"
	"web_template/config"
)

var once sync.Once

func NewRedis() *redis.Database {
	var database *redis.Database
	once.Do(func() {
		rd := config.Config.Redis
		database = redis.New(redis.Config{
			Network:   rd.Network,
			Addr:      rd.Addr + ":" + rd.Port,
			Prefix:    rd.Prefix,
			Password:  rd.Password,
			MaxActive: 20,
		})
	})

	return database
}
