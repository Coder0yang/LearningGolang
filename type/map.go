package main

import (
    "fmt"
)

func main(){
    m := make(map[string]int)
    m[""] = 1
    fmt.Println(m[""])
}
