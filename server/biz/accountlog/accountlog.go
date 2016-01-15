package accountlog

import (
	"HKCanteen/server/dao"
)

type AccountLog struct {
	User dao.DAOUser
	Logs []dao.DAOAccountLog
}

func GetMyAccountLog(userId int64) (log AccountLog) {
	err := log.User.FetchFromDB(userId)
	if err != nil {
		panic(err.Error())
	}
	if log.User.Id <= 0 {
		return
	}
	log.Logs = dao.GetMyAccountLogs(userId)
	return
}
