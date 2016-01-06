package main

import "fmt"

func main(){
    //i:=1
    var i int = 2 
    switch i{
        case 1:
            fmt.Println("one")
        case 2:
            fmt.Println("two")
        default:
            fmt.Println("invalid value1")
    }
}
