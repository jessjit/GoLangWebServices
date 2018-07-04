package packages

import "fmt"

type Str string

func Variables() {
	//VARIABLES
	var integer int = 10
	integer2 := 100
	fmt.Println(integer, integer2)
	//CONSTANT VARIABLES
	const (
		a = iota
		b = iota
		c = 3.14
		d = iota
	)
	//POINTERS
	var pointer *int = &integer
	var pointer2 *int = &integer2
	*pointer = *pointer * 10
	*pointer2 = *pointer2 / 5
	fmt.Println(*pointer, *pointer2)
	fmt.Println(integer, integer2)
	//BASIC USER TYPES
	var string1 Str = "My name is Jessjit"
	fmt.Println(string1)
}
