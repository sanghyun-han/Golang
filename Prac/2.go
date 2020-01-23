package main

import "fmt"

func main() {
	// Raw String Litheral
	rawLitheral := `arirang\n
	arirang\n
	arariyo`

	// Interpreted String literal
	interliteral := "ariarang\n arirang arariyo"

	fmt.Println(rawLitheral)
	fmt.Println()
	fmt.Println(interliteral)
}
