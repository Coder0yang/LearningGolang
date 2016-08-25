<<<<<<< HEAD
package RedisCtrl
=======
package main
>>>>>>> b10514a02d9a0e099da7f19f0668315ba13a44c4

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

<<<<<<< HEAD
func DoRedis(c redis.Conn) {
=======
func main() {
	c, _ := redis.Dial("tcp", "localhost:6379")
>>>>>>> b10514a02d9a0e099da7f19f0668315ba13a44c4
	c.Send("MULTI")
	c.Send("INCR", "foo0")
	c.Send("INCR", "foo1")
	c.Send("INCR", "foo2")
	c.Send("INCR", "foo3")
	c.Send("HINCRBY", "map", "key1", "1")
	r, _ := c.Do("EXEC")
	fmt.Println(r)
}
