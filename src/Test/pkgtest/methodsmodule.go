package pkgtest

import "fmt"

type Details struct {
	firstname string
	age       int
}
type Profile struct {
	profilename string
	salary      int
}

type showstuff interface {
	display()
}

func saysomething(s showstuff) {
	s.display()
}

func (d Details) display() {
	fmt.Println(d.firstname, d.age*2)
}
func (p Profile) display() {
	fmt.Println(p.profilename)
	fmt.Println(p.salary)
}
func Main1() {
	DETAILS := Details{
		"Jessjit",
		19}
	PROFILE := Profile{
		"Software Intern",
		500}
	fmt.Println("This is a a methods module")
	DETAILS.display()
	PROFILE.display()
	saysomething(DETAILS)
	saysomething(PROFILE)
}
