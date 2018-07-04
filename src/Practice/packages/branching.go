package packages

import "fmt"

//IF STATEMENTS
func Ifstatement(x int, enter bool) {
	if x := x * x; enter {
		fmt.Println("The squared number is: ", x)
	} else {
		fmt.Println("The number is: ", x)
	}
}

//SWITCH STATEMENTS
func Switch(x int, str string) {
	switch x {
	case 1:
		fmt.Println(str)
	case 2:
		fmt.Println(str + str)
	case 3:
		fmt.Println(str + str + str)
	default:
		fmt.Println("Did not match.")
	}
}

//SWITCH ON TYPE
func Typeswitch(x interface{}) {
	switch x.(type) { //didnt use t because we can work with just declaring x
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("Unknown")
	}
}

//BOOLEAN SWITCH
func BoolSwitch(x int) {
	switch { //didnt use t because we can work with just declaring x
	case x == 1:
		fmt.Println("One")
	case x == 2:
		fmt.Println("Two")
	default:
		fmt.Println("Three")
	}
}
