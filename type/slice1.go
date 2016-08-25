package main
import "fmt"

func main(){
    myslice := make([]int , 5 , 10)
    fmt.Println("len(myslice)",len(myslice))
    fmt.Println("cap(myslice)",cap(myslice))
    for i:=0;i<len(myslice);i++{
        myslice[i] = i
    }
    myslice = append(myslice,9,8,7)
    myslice = append(myslice,9,8,7)
    myslice = append(myslice,9,8,7)
    myslice = append(myslice,9,8,7)
    for _,v := range(myslice){
        fmt.Printf("%d ",v)
    }
    fmt.Println()
    fmt.Println("len(myslice)",len(myslice))
    fmt.Println("cap(myslice)",cap(myslice))
}
