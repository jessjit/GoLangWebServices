package vo

type Userobject struct {
	User User `json:"user"`
}

func (object Userobject) Initobj(name string, word string) Userobject {
	object.Setobject(User{
		username: name,
		password: word,
	})
	return object
}

func (object Userobject) GetUser() Userobject {
	return object
}

func (object Userobject) Getuser() User {
	return object.User
}

func (object *Userobject) Setobject(user User) {
	object.User = user
}
