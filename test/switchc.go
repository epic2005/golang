package main

import "fmt"

func main(){
    //i:=1
    var i int = 3 
    switch i{
        case 1:
            fmt.Println("one")
        case 2:
            fmt.Println("two")
        case 3:
            fmt.Println("three")
        default:
            fmt.Println("invalid value1")
    }
}
