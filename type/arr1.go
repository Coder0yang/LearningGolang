package main

import "fmt"

func main(){
    var arr [5]int 
    for j , v := range(arr){
        fmt.Printf("%d  %d\n",j,v)
    }
    for i:=0;i<5;i++{
        arr[i] = i
    }
    for j , v := range(arr){
        fmt.Printf("%d  %d\n",j,v)
    }
}
