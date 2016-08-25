package main

import "fmt"

var re = func(a int,b int,c int)(int,int,int){
    return a,b,c
}

func main(){
    var a int
    var b int
    var c int
    a,b,c = re(1,2,3)
    fmt.Println(a,b,c)
    a,b,c = 0,0,0
    _,b,c = re(1,2,3)
    fmt.Println(a,b,c)
}
