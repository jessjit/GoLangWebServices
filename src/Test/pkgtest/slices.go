package pkgtest

import "fmt"

type slice []int

func Slices(i int) {
	//var s []int
	//s = make([]int,3,10)  //10 total no of slots  //3 active slots
	var s []int
	s = make([]int, 3) //3 = total and capacity
	s[0] = 11
	s[1] = 121
	s[2] = 1330
	//Shorthand notation:
	scl := []int{1, 10, 100, 5}
	fmt.Println(scl[0])
	scl = append(scl, 7, 8, 9) //appending 7,8,9 to the existing slice
	fmt.Println(scl[5])
	scl = append(scl[:i], scl[i+1:]...) //deleting the ith element from the slice
}
func Slicingslices(ssll slice) {
	fmt.Println(ssll[1:3])
}
