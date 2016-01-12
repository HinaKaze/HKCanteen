package dao

import (
	"fmt"
	"time"
)

type DAOOrder struct {
	Id          int64
	UserId      int64 //下单者
	PayerId     int64 //付款人
	InventoryId int64
	Desc        string
	Status      string //pending,waiting,finished,closed
	TotalPrice  float64
	Time        time.Time
}

func (r *DAOOrder) SaveToDB() error {
	result, err := GetDBConn().Exec("Insert into `order` (userid,payerid,inventoryid,`desc`,`status`,totalprice,`time`) values (?,?,?,?,?,?,?)", r.UserId, r.PayerId, r.InventoryId, r.Desc, r.Status, r.TotalPrice, r.Time)
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

func (r *DAOOrder) UpdateToDB() error {
	result, err := GetDBConn().Exec("Update `order` set userid=?,payerid=?,inventoryid=?,desc=?,`status`=?,totalprice=?,time=? where id=?", r.UserId, r.PayerId, r.InventoryId, r.Desc, r.Status, r.TotalPrice, r.Time, r.Id)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (r *DAOOrder) FetchFromDB(id int64) error {
	rows, err := GetDBConn().Query("select id,userid,payerid,inventoryid,`desc`,`status`,totalprice,time from order where id=?", id)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&r.Id, &r.UserId, &r.PayerId, &r.InventoryId, &r.Desc, &r.Status, &r.TotalPrice, &r.Time)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetOrderList(status ...string) (orders []DAOOrder) {
	var statusStr string
	for _, s := range status {
		statusStr += "\"" + s + "\","
	}
	statusStr = statusStr[:len(statusStr)-1]

	sql := fmt.Sprintf("select id,userid,payerid,inventoryid,`desc`,`status`,totalprice,time from `order` where status in(%s)", statusStr)

	rows, err := GetDBConn().Query(sql)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var r DAOOrder
		err = rows.Scan(&r.Id, &r.UserId, &r.PayerId, &r.InventoryId, &r.Desc, &r.Status, &r.TotalPrice, &r.Time)
		if err != nil {
			panic(err.Error())
		}
		orders = append(orders, r)
	}
	return
}
