package main

import "fmt"
func main(){
    str := "hello world"
    n := len(str)
    for i:=0;i<n;i++{
        ch := str[i]
        fmt.Printf("%d----%c\n",i,ch)
    }
}
