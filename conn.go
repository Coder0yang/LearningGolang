package redistest

import (
	"github.com/garyburd/redigo/redis"
)

//var conn redis.Conn
func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 200,
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

/*
func Init()error{
    pool = newPool()
}
*/

func getConn() redis.Pool {
	c := pool.Get()
	return c
}

func (c *redis.Pool) releaseConn() {
	c.Close()
}

func DoGet(userId string) {
	c := getConn()
	defer c.releaseConn()
}

func PipeGet(userId string) {
}
func Release() {
}
