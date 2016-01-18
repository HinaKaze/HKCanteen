package user

import (
	"HKCanteen/server/dao"
)

const (
	ResultSuccess     = 1
	ResultNotMatch    = 2
	ResultUserInvalid = 3
)

func Login(username string, password string) (user dao.DAOUser, result int) {
	user = dao.GetUserByUsername(username)
	if user.Id <= 0 {
		return user, ResultUserInvalid
	}
	if user.Username == username && user.Password == password {
		return user, ResultSuccess
	} else {
		return user, ResultNotMatch
	}
}

func GetUserList() (users []dao.DAOUser) {
	users = dao.GetAllUsers()
	return
}
