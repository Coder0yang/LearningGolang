package conredis

import "github.com/garyburd/redigo/redis"

func Doredis(c redis.Conn,cmd string,args ...interface{})(interface{},error){
    pending := c.pending
    var argSlice []int = args[:num]
    v,err := c.Do(cmd,argSlice)
    return v,err
}
