package packages

import "fmt"

//TRADITIONAL FOR
func TradFor(num int) {
	for i := 0; i < num; i++ {
		fmt.Print(i * i)
		fmt.Print(" ")
	}
	fmt.Println()
}

//WHILE LOOP
func WhileFor(num int) int {
	total := 1
	i := 1
	for i <= num {
		total = total * i
		i++
	}
	return total
}

//INFINITE LOOP
func InfiniteFor(num int) {
	k := 0
	for {
		if k%2 == 0 {
			k++
			continue
		} else {
			fmt.Println(k)
			k++
		}
		if k >= num {
			break
		}
	}
}
