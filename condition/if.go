package main
import "fmt"
func main(){
    myArrary := [10]int{0,1,2,3,4,5,6,7,8,9}
    var mySlice []int = myArrary[2:6]
    fmt.Println("Elements of myArrary: ")
    for _,v := range mySlice{
        fmt.Print(v , " ")
        if v == 1{
            fmt.Println("111")
        }else if v == 2{
            fmt.Println("222")
        }else if v == 3{
            fmt.Println("333")
        }else if v == 4{
            fmt.Println("444")
        }else{
            fmt.Println("other")
        }
    }
    fmt.Println()
}
