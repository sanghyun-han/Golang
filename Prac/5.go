package main

func main() {
	total, num := sum(1, 2, 3, 4, 6, 7, 8)
	println("total is", total, "num is", num)
}
func sum(nums ...int) (count int, total int) {
	println(total)
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}
