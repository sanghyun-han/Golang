package main

func main() {

	k := 3
	sum := 0
	names := []string{"shhan", "shan", "sanghyun"}

	if k == 1 {
		println("One")
	} else if k == 2 {
		println("Two")
	} else {
		println("Others")
	}

	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	for idx, name := range names {
		println(idx, name)
	}
}
