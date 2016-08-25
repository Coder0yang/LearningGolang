package main
import "fmt"
func modify(array [10] int){
    array[0] = 10
    fmt.Println("In modify(),array values:",array)
}

func main(){
    array := [10]int{1,2,3,4,5,6,7,8,9,0}
    modify(array)
    fmt.Println("In main(),array values:",array)
}
