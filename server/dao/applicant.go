package dao

import (
	"fmt"
	"strconv"
)

type DAOApplicant struct {
	Id      int64
	OrderId int64
	UserId  int64
	Status  string //join,hesitation,cancel
	Comment string
	Pay     float64
}

func (r *DAOApplicant) SaveToDB() error {
	result, err := GetDBConn().Exec("Insert into applicant (orderid,userid,status,comment,pay) values (?,?,?,?,?)", r.OrderId, r.UserId, r.Status, r.Comment, r.Pay)
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

func (r *DAOApplicant) UpdateToDB() error {
	result, err := GetDBConn().Exec("Update applicant set orderid=?,userid=?,status=?,comment=?,pay=? where id=?", r.OrderId, r.UserId, r.Status, r.Comment, r.Pay, r.Id)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (r *DAOApplicant) FetchFromDB(id int64) error {
	rows, err := GetDBConn().Query("select id,orderid,userid,status,comment,pay from applicant where id=?", id)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&r.Id, &r.OrderId, &r.UserId, &r.Status, &r.Comment, &r.Pay)
		if err != nil {
			return err
		}
		r.Pay, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", r.Pay), 64)
	}
	return nil
}

func GetApplicantsByOrderId(orderId int64) (applicants []DAOApplicant) {
	rows, err := GetDBConn().Query("select id,orderid,userid,status,comment,pay from applicant where orderid=?", orderId)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var r DAOApplicant
		err = rows.Scan(&r.Id, &r.OrderId, &r.UserId, &r.Status, &r.Comment, &r.Pay)
		if err != nil {
			panic(err.Error())
		}
		r.Pay, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", r.Pay), 64)
		applicants = append(applicants, r)
	}
	return
}

func GetApplicantByOrderId(userId int64, orderId int64) (applicant DAOApplicant) {
	rows, err := GetDBConn().Query("select id,orderid,userid,status,comment,pay from applicant where orderid=? and userid=?", orderId, userId)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&applicant.Id, &applicant.OrderId, &applicant.UserId, &applicant.Status, &applicant.Comment, &applicant.Pay)
		if err != nil {
			panic(err.Error())
		}
		applicant.Pay, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", applicant.Pay), 64)
	}
	return
}
