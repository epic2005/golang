package main

func test(a, b int){
	defer println("dispose1 ...")
	println(a / b)
}

func main(){
	test(10, 0)
}