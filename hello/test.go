package main

import "fmt"

var a int = 5

func main() {
	b := 2
	var c, d int
	c = a / b
	d = a % b
	fmt.Println("商:", c, "余数:", d)
}

