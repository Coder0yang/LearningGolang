package main
import "fmt"

type Integer int 
func Integer_Less(a Integer,b Integer)bool{
    return a<b
}

func (a Integer)Less(b Integer)bool{
    return a<b
}

func main(){
    var a Integer = 3 
    if Integer_Less(a,2){
        fmt.Println(a,"Less 2")
    }else{
        fmt.Println(a,"Biger 2")
    }
    if a.Less(2){
        fmt.Println(a,"Less 2")
    }else{
        fmt.Println(a,"Biger 2")
    }
}
