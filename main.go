package main

import (
    "fmt"
)

var c, python, java bool

func add(x,y int)int{
	return x+y
}

func swap(y, z string) (string, string) {
	return z, y
}

func split(sum int)(x, y int){
	x = sum * 4 /9 
	y = sum - x
return
}


func main() {
    // fmt.Println(add(3,4))
	// fmt.Println(swap("jatin","foujdar"))
	// a, b := swap("jatin","foujdar")
	// fmt.Println(a, b)
	// fmt.Println(split(17))
	 i := 0
	fmt.Println(i,c, python,java)
}
