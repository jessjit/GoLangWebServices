package main

import (
	"Test/pkgtest"
	"fmt"
)

//structures
type sal struct {
	number     int
	salutation string
}

//constant variables and concept of iota
const (
	A        = iota     //A = 0
	B        = iota     //B = 1
	Pi       = 3.14     //2
	Language = "Golang" //3
	C        = iota     //C = 4
)

//type switch case
func typeswitch(x interface{}) {
	switch x.(type) { //didnt use t because we can work with just declaring x
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("Unknown")
	}
}

//switch case
func getprefix(name string) (prefix string) {
	switch name {
	case "J", "j":
		prefix = "A"
	case "Jess":
		prefix = "B"
		fallthrough
	case "Jessjit":
		prefix = "C"
	}
	return prefix
}

//functions
func greetme(dagreeting sal) {
	fmt.Println(dagreeting.number)
	fmt.Println(dagreeting.salutation)
}

//returning functions + if statement
func returngreeting(number int, name string, isFormal bool) (string, int) { //multiple returns
	if isFormal {
		return name, number
	}
	var x string = "Yooo wassuppp"
	return x, number
}

//named returns
func namedreturn(number int, name string) (usrname string, usrno int) {
	usrname = name
	usrno = number * 2
	return
}

//variadic function
func variadic(number int, name ...string) (a int, b string, c string, d string) {
	fmt.Println(len(name))
	a = number
	b = name[0]
	c = "Jessjit"
	d = "Jess"
	return
}
func main() {
	typeswitch("Hello There")                                                            //type switch case
	typeswitch(88)                                                                       //type switch case
	typeswitch(34.45)                                                                    //type switch case
	var switchs = sal{1, "Jess"}                                                         //switch case
	fmt.Print(getprefix(switchs.salutation))                                             //switch case
	fmt.Print(" ")                                                                       //switch case
	fmt.Println(switchs.salutation)                                                      //switch case
	var kaafis = sal{22, "Variadic Function"}                                            //variadaic function
	fmt.Println(variadic(kaafis.number, kaafis.salutation))                              //variadic function
	var sss = sal{36, "Jessjit's returned named function"}                               //named returns
	fmt.Println(namedreturn(sss.number, sss.salutation))                                 //named returns
	var ss = sal{49, "Jessjit Singh's returned function"}                                //returning functions
	fmt.Println(returngreeting(ss.number, ss.salutation, true))                          //returning funtions
	var s = sal{69, "Jessjit Singh's Function"}                                          //functions
	greetme(s)                                                                           //functions
	fmt.Println(A, B, C)                                                                 //basic user types
	fmt.Println(Pi)                                                                      //basic user types
	fmt.Println(Language)                                                                //basic user types
	var structure = sal{87, "Hello Everyone"}                                            //structures
	fmt.Println("Number:", structure.number, "Greetings:", structure.salutation)         //structures
	var ultastructure = sal{salutation: "This is ulta calling", number: 99}              //structures
	fmt.Println("Number:", ultastructure.number, "Greetings:", ultastructure.salutation) //structures

	//normal stuff

	var a, b, c int = 2, 3, 4
	message := "This is a string variable"
	var greeting *string = &message
	fmt.Println(message, greeting)
	fmt.Print(a, b, c)
	fmt.Println(" are the values of a,b and c")
	fmt.Println("Test")
	pkgtest.Test()
	pkgtest.Hello()
	pkgtest.World1()
	pkgtest.Forloop() //loops
	fmt.Println()
	pkgtest.Maps(1)                 //maps
	pkgtest.Maps(3)                 //maps
	pkgtest.Maps(5)                 //maps
	pkgtest.Shorthandmap("Jatinn")  //maps
	pkgtest.Shorthandmap("Jatin")   //maps
	pkgtest.Shorthandmap("Harshit") //maps
	pkgtest.Slices(3)               //slices
	mainslice := []int{0, 1, 2, 3, 4, 5}
	pkgtest.Slicingslices(mainslice)

	//Methods and Interfaces

	pkgtest.Main1()
}
