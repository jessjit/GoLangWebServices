package repositories

import (
	"Loginmodule/data"
	"Loginmodule/vo"
)

func Insertdb(user *vo.User) (val bool) {
	user.Getuserid()
	err := data.Collection.Insert(&user)
	if err != nil {
		panic(err)
		return false
	} else {
		return true
	}
}
