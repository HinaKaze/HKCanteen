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

func Login(username string, password string) int {
	u := dao.GetUserByUsername(username)
	log.Printf("u[%+v]", u)
	if u.Id <= 0 {
		return ResultUserInvalid
	}
	if u.Username == username && u.Password == password {
		return ResultSuccess
	} else {
		return ResultNotMatch
	}
}
