package packages

import "fmt"

func Slices(i int, k int) {
	var slice []int
	slice = make([]int, i)
	for j := 1; j <= i; j++ {
		slice[j-1] = i * j
	}
	fmt.Println(slice[i/2])
	fmt.Println(slice)
	slice = append(slice, 10, 100, 1000)      //appendig elements
	slice = append(slice[:k], slice[k+1:]...) //deleting elements
	fmt.Println(slice)
}
