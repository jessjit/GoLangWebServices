package Service

import (
	"fmt"
	"net/http"
)

func Introhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to the website")
	fmt.Println("Signup:s,Login:l")
	fmt.Println("What would yoiu like to do? Enter the code:")
	var choice string
	fmt.Scanf("%s", choice)
	if choice == "s" {
		http.Redirect(w, r, "/sigup", 301)
	} else if choice == "l" {
		http.Redirect(w, r, "/login", 301)
	} else {
		fmt.Println("Wrong code. Go home")
	}
}
