package model

import (
	"blog/server/dao"
)

type User struct {
	dao.DAOUser
	Account      dao.DAOAccount
	PendingOrder dao.DAOOrder
}
