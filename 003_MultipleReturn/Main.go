package main

import (
	"fmt"
)

func main() {

	fmt.Println(calc(10,12))
}

func calc(a int,b int) (sum int,multiple int)  {

	result1:=a+b
	result2:=a*b

	return result1,result2
}
