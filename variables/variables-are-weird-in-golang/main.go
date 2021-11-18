// Followed tuto : https://www.youtube.com/watch?v=sZoRSbokUE8

package main

import "fmt"

func main() {
	const a int = 10
	fmt.Println(a)

	//a = 20
	//fmt.Println(a)

	temp := 3
	fmt.Printf("%v, %T\n", temp, temp)

	temp2 := 3.01
	fmt.Printf("%v, %T\n", temp2, temp2)

	temp3 := 3.
	fmt.Printf("%v, %T\n", temp3, temp3)

	var (
		sp  = "oka"
		age = 999
		pw  = "noth"
	)
	fmt.Println(sp, age, pw)
}
