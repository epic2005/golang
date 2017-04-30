package main

func test(a, b int){
	defer println("dispose2 ...")
	println(a / b)
}

func main(){
	test(10, 0)
}