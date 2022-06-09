package main

import "fmt"

func main() {
	fmt.Println("No. of integers")
	var num int
	fmt.Scanln(&num)

	list := make([]int, num)

	for i := 0; i < num; i++ {
		fmt.Println("enter digits")
		var digits int
		fmt.Scanln(&digits)
		list[i] = digits
	}
	fmt.Println("list: ", list)

	l1 := []int{}
	l2 := []int{}

	for i := 0; i < num; i++ {
		if list[i]%2 == 0 {
			l1 = append(l1, list[i])
		} else {
			l2 = append(l2, list[i])
		}
	}
	fmt.Println(l1)
	fmt.Println(l2)
	// if (digits % 2) == 0 {
	// 	fmt.Println("even no.")

	// } else {
	// 	fmt.Println("odd no.")
	// }

}
