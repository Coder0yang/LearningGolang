package main
import "fmt"
import "conredis"

func main(){
    str1 := "tcp"
    str2 := "127.0.0.1:6379"
    c,_ := conredis.Connect(str1,str2)
    for i := 0 ;i < 3;i++{
        fmt.Println("----------")
        //v,err := c.Do("LPUSH","num",i)
        v,err := conredis.Doredis(c,"LPUSH","num",i)
        if err != nil{
            fmt.Println(err)
            break
        }
        fmt.Println(v)
    }

    defer c.Close()
}
