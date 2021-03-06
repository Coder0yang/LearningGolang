package redistest

import (
	"github.com/garyburd/redigo/redis"
)

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 20000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

var pool = newPool()

func getConn() redis.Conn {
	c := pool.Get()
	return c
}

func releaseConn() error {
	return nil
}

func DoGet(userId string) {
	c := pool.Get()
	c.Do("INCR", userId)
	defer c.Close()
}

func GetActiveConnNum() int {
	return pool.ActiveCount()
}

func PipeGet(userId string) {
	c := pool.Get()
	c.Send("INCR", userId)
	defer c.Close()
}
