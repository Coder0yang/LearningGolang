package main
import "fmt"

type Rect struct{
    x,y float64
    width,height float64
}

func (r *Rect) Area() float64{
    return r.width * r.height
}

func NewRect(x,y,width,height float64) *Rect{
    return &Rect{x,y,width,height}
}

func main(){
    Rec := NewRect(1,1,2,2)
    b := Rec.Area()
    fmt.Println(b)
}
