package dao

import (
	"fmt"
	"strconv"
	"time"
)

type DAOAccountLog struct {
	Id      int64
	UserId  int64
	Type    string //withdraw,charge,spend...
	Value   float64
	Time    time.Time
	OrderId int64
}

func (r *DAOAccountLog) SaveToDB() error {
	result, err := GetDBConn().Exec("Insert into account_log (userid,type,value,time,orderid) values (?,?,?,?,?)", r.UserId, r.Type, r.Value, r.Time, r.OrderId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	r.Id = id
	return nil
}

func (r *DAOAccountLog) UpdateToDB() error {
	result, err := GetDBConn().Exec("Update account_log set userid=?,type=?,value=?,time=?,orderid=? where id=?", r.UserId, r.Type, r.Value, r.Time, r.OrderId, r.Id)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (r *DAOAccountLog) FetchFromDB(id int64) error {
	rows, err := GetDBConn().Query("select id,userid,type,value,time,orderid from account_log where id=?", id)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&r.Id, &r.UserId, &r.Type, &r.Value, &r.Time, &r.OrderId)
		if err != nil {
			return err
		}
		r.Value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", r.Value), 64)
	}
	return nil
}

func GetMyAccountLogs(userid int64) (logs []DAOAccountLog) {
	rows, err := GetDBConn().Query("select id,userid,type,value,time,orderid from account_log where userid=?", userid)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var log DAOAccountLog
		err = rows.Scan(&log.Id, &log.UserId, &log.Type, &log.Value, &log.Time, &log.OrderId)
		if err != nil {
			panic(err.Error())
		}
		log.Value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", log.Value), 64)
		logs = append(logs, log)
	}
	return
}
