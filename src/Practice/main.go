package main

import (
	"Practice/packages"
	"fmt"
)

func main() {
	packages.Hello()
	packages.Variables()
	packages.Function1(12, 23)
	fmt.Println(packages.Function2(7, 8))
	fmt.Println(packages.Function3(3, 13))
	packages.FunctionVar(3, 4, 5, 6)
	packages.FunctionVar(12, 11)
	packages.FunctionVar(3, 4, 7, 8, 6, 5, 12)
	packages.FunctionPass(17, "Hello Everyone. ", packages.Printmul2)
	packages.FunctionPass(17, "Hello Everyone. ", packages.Printmul3)
	packages.FunctionPass(17, "Hello Everyone. ", packages.Printnorm)
	packages.FunctionPass(12, "Hello There", packages.Createprintstring("!!!"))
	packages.Ifstatement(42, true)
	packages.Switch(5, "jhasj")
	packages.Switch(1, "Hello")
	packages.Switch(2, "Jessjit ")
	packages.Typeswitch(212)
	packages.Typeswitch("Jessjit")
	packages.Typeswitch(34.45)
	packages.BoolSwitch(2)
	packages.TradFor(5)
	fmt.Println(packages.WhileFor(5))
	packages.InfiniteFor(15)
	packages.Slices(5, 3)
	packages.Maps(10, 5)
	//CONCURRENCY
	c := make(chan string)
	go packages.Count("sheep", c)
	for message := range c {
		fmt.Println(message)
	}
	c1 := make(chan int, 3)
	c1 <- 1
	c1 <- 2
	c1 <- 3
	number := <-c1
	fmt.Println(number)
	number = <-c1
	fmt.Println(number)
	number = <-c1
	fmt.Println(number)

	//BASIC CONCURRENCY QUESTION
	inchan := make(chan string)
	scll := []string{"Good", "Boys", "are", "cool"}
	go packages.CountVal(scll, inchan)
	for msg := range inchan {
		fmt.Println(msg)
	}
}

//FUNCTIONS FOR CONCURRENCY/SELECT
func worker(jobs <-chan int, result chan<- int) {
	for n := range jobs {
		result <- fib(n)
	}
}
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
