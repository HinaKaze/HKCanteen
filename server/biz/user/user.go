package user

import (
	"HKCanteen/server/dao"
	"log"
)

const (
	ResultSuccess     = 1
	ResultNotMatch    = 2
	ResultUserInvalid = 3
)

func Login(username string, password string) (userId int64, result int) {
	u := dao.GetUserByUsername(username)
	log.Printf("u[%+v]", u)
	if u.Id <= 0 {
		return -1, ResultUserInvalid
	}
	if u.Username == username && u.Password == password {
		return u.Id, ResultSuccess
	} else {
		return -1, ResultNotMatch
	}
}
