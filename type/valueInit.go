package main

import (
    "fmt"
)

type OO struct{
    A int
    B string
    C bool
}

func main(){
    var m OO
    fmt.Println(m.A,"--",m.B,"--",m.C)
}
