package main
import "fmt"


func f(msg string){
    fmt.Println(msg)
}

func main(){
    go f("goroutine")
    //fmt.Println("test")
    //go func(msg string){
    //    fmt.Println(msg)
    //}("going")
}
