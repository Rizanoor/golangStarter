package main

import "fmt"

func main()  {
	/** 
	age := 20

		if age > 10 {
			fmt.Println("boleh main game")
		}else {
			fmt.Println("tidak boleh main game")
		}
	*/

	score := 65
	var grade string
	
		if score <= 50 {
			grade = "E"
		} else if score <= 60{
			grade = "D"
		} else if score < 70 {
			grade = "C"
		}else {
			grade = "NULL"
		}

		fmt.Println(grade)



}