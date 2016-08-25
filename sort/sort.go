package main

import(
    "os"
    "fmt"
    "bufio"
    "strconv"
)

func main(){
    reader := bufio.NewReader(os.Stdin)
    data,_,_ := reader.ReadLine()
    command := strconv.Atoi(data)
    fmt.Println(command)
}
