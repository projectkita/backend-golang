package initial

import (
	"github.com/garyburd/redigo/redis"
	"github.com/projectkita/project-harapan-backend-golang/src/config"
	"github.com/projectkita/project-harapan-backend-golang/src/lib/database"
)

var APP Module

// Module is a package struct
type Module struct {
	Config    config.AppConfig
	DB        *database.DBsql
	RedisPool *redis.Pool
	Debug     func(...interface{})
}
