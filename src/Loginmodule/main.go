package main

import (
	"Loginmodule/controller"
	"Loginmodule/data"
)

func main() {
	r := controller.Createrouter()
	data.Activatedb("userdb", "usercollection")
	controller.Runserver(r)
}
