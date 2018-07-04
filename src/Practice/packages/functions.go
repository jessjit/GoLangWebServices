package packages

import "fmt"

//BASIC FUNCTION
func Function1(x int, y int) {
	s1 := "The result is: "
	z := x + y + x*y
	fmt.Println(s1)
	fmt.Println(z)
}

//FUCNTION WITH RETURN
func Function2(x int, y int) (string, int) {
	s1 := "The result is: "
	z := x*y + x + y
	return s1, z
}

//NAMED RETURNS
func Function3(x int, y int) (named_string string, result int) {
	s1 := "The result is: "
	result = 2 * (x + y)
	named_string = s1
	return
}

//VARIDAIC FUNCTIONS
func FunctionVar(numbers ...int) {
	var total int = 1
	for _, n := range numbers {
		total = total * n
	}
	fmt.Println(total)
}

//PASSING A FUNCTION   //can even define a type of function to replace func(string) of the passed function
//eg:-
//type Printer func(string)()
func FunctionPass(num int, str string, printstring func(string)) { //or printstring Printer
	num = num * num
	str = str + "Jessjit Singh"
	fmt.Println(num)
	printstring(str)
}
func Printmul3(str string) {
	fmt.Println(str + str + str)
}
func Printmul2(str string) {
	fmt.Println(str + str)
}
func Printnorm(str string) {
	fmt.Println(str)
}

//CLOSURE:-returns a function
func Createprintstring(s string) func(string) {
	getfunc := func(strr string) {
		fmt.Println((strr + s) + (strr + s) + (strr + s))
	}
	return getfunc
}
