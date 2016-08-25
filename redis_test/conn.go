package redistest

import (
    "github.com/garyburd/redigo/redis"
)

var conn redis.Conn

func Init(redisIp string,redisPort string)error{
    var err error
    conn,err = redis.Dial(redisIp,redisPort)
    return err
}

func DoGet(userId string)string{
    str,_ := redis.String(conn.Do("HGET",userId,"ad1"))
    return str
}

func PipeGet(userId string)string{
    conn.Send("HGET",userId,"ad1")
    conn.Flush()
    str,_ := redis.String(conn.Receive())
    return str
}
func Release(){
    conn.Close()
}
