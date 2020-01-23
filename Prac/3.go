package main

func main() {
	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes)
	println(bytes, str2)
}
