package main

import (
    "fmt"
    "sync"

    "github.com/garyburd/redigo/redis"
)

func foo()interface{}{
    a,_ := redis.Dial("tcp","localhost:6379")
    return a
}

func main(){
    newFun := func ()interface{}{
        c,_ := redis.Dial("tcp","localhost:6379")
        return c
    }
    pool := &sync.Pool{New:newFun}
    t := pool.Get() 

    /*switch v := t.(type){
    case int:
        fmt.Println(v)
    case string:
        fmt.Println(v)
    case redis.Conn:
        fmt.Println("redis.Conn")
        a,_:=redis.Int(v.Do("get","conn"))
        fmt.Printf("%d\n",a)
    case *redis.Conn:
        fmt.Println("*redis.Conn")
    }*/
    a,_:=redis.Int(t.Do("get","conn"))
    fmt.Printf("%d\n",a)

}
