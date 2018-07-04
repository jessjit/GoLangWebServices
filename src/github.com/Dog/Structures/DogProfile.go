package Structures

//Basic Dog Profile
type Dogprofile struct {
	Dogname string   `json:"dogname"`
	Age     int      `json:"age"`
	Breed   string   `json:"breed"`
	History *History `json:"history"`
}

//History of Medical Checkups
type History struct {
	Numberofvisits int    `json:"numberofvisits"`
	Lastshown      string `json:"lastshown"`
}
