package Structures

//Users Information
type Userprofile struct {
	Username     string `json:"username"`
	Userpassword string `json:"password"`
}

//New Users Account Information
type Newuserprofile struct {
	Emailid           string `json:"emailid"`
	Firstname         string `json:"firstname"`
	Lastname          string `json:"lastname"`
	Age               int    `json:"age"`
	Validusername     string `json:"validusername"`
	Validuserpassword string `json:"validuserpassword"`
}
