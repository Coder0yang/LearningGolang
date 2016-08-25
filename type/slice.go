package main
import "fmt"
func main(){
    var myArrary [10]int = [10]int{1,2,3,4,5,6,7,8,9,0}
    //基于数组创建一个数组切片
    var mySlice []int = myArrary[:5]
    fmt.Println("Elements of myArrary: ")
    for _,v := range myArrary{
        fmt.Print(v, " ")
    }
    fmt.Println("\nElements of mySlice: ")
    for _,v := range mySlice{
        fmt.Print(v," ")
    }
    fmt.Println()
    var m_slice []int = myArrary[:len(myArrary)]
    for _,v := range m_slice{
        fmt.Print(v," ")
    }
    fmt.Println()
}
