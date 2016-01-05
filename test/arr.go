package main

import "fmt"


func main(){
    var a [5]int
    fmt.Println("array a:", a)
    a[1] = 10
    a[3] = 30

    fmt.Println("assign:",a)
    fmt.Println("len: ",len(a))
}
