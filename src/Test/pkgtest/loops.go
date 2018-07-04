package pkgtest

import (
	"fmt"
)

func Forloop() {
	fmt.Println("There is only one loop in  Go ie for loop")
	fmt.Println("We will look at a basic for loop")
	Range := 10
	for i := 0; i < Range; i++ { //traditional for loop
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Println()
	var j int = 0
	for j < Range { //while loop example
		fmt.Print(j * j)
		fmt.Print(" ")
		j++
	}
	fmt.Println()
	var k int = 0
	for {
		if k%2 == 0 {
			k++
			continue
		}
		fmt.Print(k)
		k++
		fmt.Print(" ")
		if k == Range {
			break
		}
	}
}
