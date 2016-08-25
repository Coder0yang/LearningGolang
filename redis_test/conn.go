package redistest

import (
	"github.com/garyburd/redigo/redis"
)

//var conn redis.Conn
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

/*
func Init()error{
    pool = newPool()
}
*/

func getConn() redis.Conn {
	c := pool.Get()
	return c
}

func releaseConn(c *redis.Conn) {
	(*c).Close()
}

func DoGet(userId string) {
	c := getConn()
	c.Do("INCR", userId)
	c.Close()
}

func PipeGet(userId string) {
	c := getConn()
	c.Send("INCR", userId)
	c.Close()
}
