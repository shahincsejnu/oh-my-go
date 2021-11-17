// Followed : https://gobyexample.com/variables

package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c = 1, 2
	fmt.Println(b, c)

	var d, e int = 3, 4
	fmt.Println(d, e)

	var f int
	var g string
	var h bool
	fmt.Println(g, f, h)

	i := 10
	j := "oka"
	k := true
	fmt.Println(i, j, k)
}
