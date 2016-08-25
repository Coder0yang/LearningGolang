package conredis

import "fmt"
//import "errors"
import "github.com/garyburd/redigo/redis"

func Connect(str1,str2 string) (redis.Conn,error) {
    c,err := redis.Dial(str1,str2)
    if err != nil{
        fmt.Println(err)
        return nil,err
    }
    return c,err
}
