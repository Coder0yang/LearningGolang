package main

import (
    "fmt"
    "reflect"
)

func main(){
    var x float64 = 1.1
    fmt.Println("reflect.Valueof:",reflect.ValueOf(x))
    fmt.Println("reflect.Type:",reflect.TypeOf(x))

    v := reflect.ValueOf(x)

    fmt.Println("reflect.Type:",v.Type())
    fmt.Println("actual value:",v.Float())
    fmt.Println("kind is float64?",v.Kind() == reflect.Float64)
}
