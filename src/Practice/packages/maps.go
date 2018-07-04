package packages

import "fmt"

func Maps(x int, y int) {
	var map1 map[int]int
	map1 = make(map[int]int)
	for i := 0; i < x; i++ {
		map1[i] = i * i
	}
	fmt.Println(map1)
	fmt.Println(map1[1])
	map1[x] = x * x
	fmt.Println(map1[x])
	delete(map1, y)
	fmt.Println(map1)
}
