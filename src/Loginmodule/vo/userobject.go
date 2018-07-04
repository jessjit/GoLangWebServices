package vo

type Userobject struct {
	User User `json:"user"`
}

func (object Userobject) Getuser() User {
	return object.User
}

func (object *Userobject) Setobject(user User) {
	object.User = user
}
