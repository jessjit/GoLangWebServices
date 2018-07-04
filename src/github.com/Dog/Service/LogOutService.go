package Service

import (
	"fmt"
	"net/http"
)

func Logouthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Thank you. Redirecting to Intro Page")
	http.Redirect(w, r, "/intro", 301)
}
