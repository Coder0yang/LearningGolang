package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, _ := redis.Dial("tcp", "localhost:6379")
	c.Send("MULTI")
	c.Send("INCR", "foo0")
	c.Send("INCR", "foo1")
	c.Send("INCR", "foo2")
	c.Send("INCR", "foo3")
	c.Send("HINCRBY", "map", "key1", "1")
	r, _ := c.Do("EXEC")
	fmt.Println(r)
}
