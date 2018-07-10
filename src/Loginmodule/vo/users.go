package vo

type Usersobject struct {
	users []User `json:"user"`
}

func (object Usersobject) GetUsers() []User {
	return object.users
}

func (object *Usersobject) SetUsers(users []User) {
	object.users = users
}
