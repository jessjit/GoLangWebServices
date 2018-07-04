package Service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dog/Structures"
	jwt "github.com/dgrijalva/jwt-go"
)

func Signuphandler(w http.ResponseWriter, r *http.Request) {
	var NewUser Structures.Newuserprofile
	_ = json.NewDecoder(r.Body).Decode(&NewUser)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": NewUser.Validusername,
		"password": NewUser.Validuserpassword,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(Structures.Jwttoken{Token: tokenString})
}
