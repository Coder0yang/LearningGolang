package main

import (
    "fmt"
)

type student struct{
    name string
    snum int
    class string
}

func (stu *student)setinfor(_name string,_snum int,_class string){
    stu.name = _name
    stu.snum = _snum
    stu.class = _class
}

func (stu student)showinfor(){
    fmt.Printf("name:%s\nsnum:%d\nclass:%s\n",stu.name,stu.snum,stu.class)
}

func main(){
    var yang student
    yang.setinfor("yang",25,"网络13-2班")
    yang.showinfor()
}
