package main
import "fmt"

type PersionInfo struct{
    ID string
    Name string
    Address string
}

func main(){
    var personDB map[string]PersionInfo
    personDB = make(map[string] PersionInfo)

    //插入数据
    personDB["12345"] = PersionInfo{"12345","Tom","Room 203"}
    personDB["12346"] = PersionInfo{"12346","Jack","Room 104"}

    //从map中查找"12345"的信息
    delete(personDB,"12345")
    personDB["12345"] = PersionInfo{"12345","Tom","Room 203"}
    person,ok := personDB["12345"]
    if ok {
        fmt.Println("Found the person",person.Name,"with ID 12345")
    } else {
        fmt.Println("Can't found the person 12345")
    }
}
