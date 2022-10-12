package main

import (
	"fmt"
	"hello/calculation"	
)

func main()  {
	fmt.Println("hallo world")

	result := calculation.Add(9, 9)

	fmt.Println(result)

}
