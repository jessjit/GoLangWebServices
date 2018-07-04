package Service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dog/Structures"
	jwt "github.com/dgrijalva/jwt-go"
)

func Loginhandler(w http.ResponseWriter, r *http.Request) { //creates a token and passes it back to the user
	var User Structures.Userprofile
	_ = json.NewDecoder(r.Body).Decode(&User)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": User.Username,
		"password": User.Userpassword,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(Structures.Jwttoken{Token: tokenString})
}
