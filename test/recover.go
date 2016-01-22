package main
import "fmt"


func g(i int){
    if i > 1{
        fmt.Println("panic")
        panic(fmt.Sprintf("%v",i))
    }
}

func f(){
    defer func(){
        if r:= recover(); r != nil{
            fmt.Println("Recovered in f", r)
        }
    }()


    for i:=0;i<4;i++{
        fmt.Println("calling g with", i)
        g(i)
        fmt.Println("Returned normally form g.")
    }
}

func main(){
    f()
    fmt.Println("Returned normally form f.")
    
}
