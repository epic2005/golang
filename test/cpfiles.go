package main
//import "fmt"
import "os"
import "io"


func CopyFile(dstName, srcName string)(written int64, err error){
    src, err := os.Open(srcName)
    if err != nil{
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil{
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}


func main(){
    CopyFile("2.txt", "1.txt")
}
