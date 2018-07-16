package main

import (
	"Loginmodule/controller"
	"Loginmodule/data"
	"Loginmodule/logs"
)

func main() {
	r := controller.Createrouter()
	data.Activatedb("userdb", "usercollection")
	logs.Setlogger()
	controller.Runserver(r)
}
