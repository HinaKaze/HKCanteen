package model

import (
	"HKCanteen/server/dao"
)

type User struct {
	dao.DAOUser
	Account      dao.DAOAccount
	PendingOrder dao.DAOOrder
}
