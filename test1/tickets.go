package main

import "fmt"
import "time"
import "math/rand"
import "runtime"

var total_tickets int32 = 10;

func sell_tickets(i int){
    for{
        if total_tickets > 0 {
            time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
            total_tickets--
            fmt.Println("id:",i, "ticket:", total_tickets)
        }else{
            break
        }
    }
}


func main(){
    runtime.GOMAXPROCS(4)
    rand.Seed(time.Now().Unix())

    for i:=0;i<5;i++{
        go sell_tickets(i)
    }
    var input string
    fmt.Scanln(&input)
    fmt.Println(total_tickets, "done")
}
