package main

import "fmt"

func fun(a *int,b *int){
    *a , *b = *b , *a
}

func main(){
    a := 1
    b := 2
    fmt.Println(a,b)
    fun(&a,&b)
    fmt.Println(a,b)
}
