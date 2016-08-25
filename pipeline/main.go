package main
import (
    "fmt"

    "github.com/garyburd/redigo/redis"
)

func main(){
    c,err := redis.Dial("tcp","127.0.0.1:6379")
    if err != nil {
        fmt.Println(err)
        return
    }
    c.Send("HINCRBY","pipetest","num","1")
    c.Send("HINCRBY","pipetest","num","1")
    c.Send("HINCRBY","pipetest","num","1")
    c.Send("HINCRBY","pipetest","num","1")
    c.Flush()
    c.Send("HGETALL","pipetest")
    c.Flush()
    str,_ := c.Receive()
    fmt.Printf("%d\n",str)
    c.Flush()
    str,_ = c.Receive()
    fmt.Printf("%s\n",str)
}
