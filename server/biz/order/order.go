package order

import (
	"log"
	"time"

	"github.com/HinaKaze/canteen/server/biz/applicant"
	"github.com/HinaKaze/canteen/server/dao"
)

type Order struct {
	Order      dao.DAOOrder
	Creator    dao.DAOUser
	Applicants []applicant.Applicant
}

func CreateOrder(userId int64, desc string) (order dao.DAOOrder) {
	order.UserId = userId
	order.Desc = desc
	order.Status = "pending"
	order.Time = time.Now()
	err := order.SaveToDB()
	if err != nil {
		panic(err.Error())
	}
	return
}

func GetOrderDetail(orderId int64) (order Order) {
	log.Println(orderId)
	order.Order.FetchFromDB(orderId)
	log.Printf("+v\n", order.Order)
	if order.Order.Id <= 0 {
		return
	}
	order.Creator.FetchFromDB(order.Order.UserId)
	order.Applicants = applicant.GetApplicantsByOrderId(orderId)
	return
}

func OrderRun(orderId int64, totalprice float64) (order dao.DAOOrder) {
	order.FetchFromDB(orderId)
	order.Status = "running"
	order.TotalPrice = totalprice
	order.UpdateToDB()
	return
}

func OrderFinish(orderId int64) (order dao.DAOOrder) {
	order.FetchFromDB(orderId)
	order.Status = "finished"
	order.UpdateToDB()

	var payerCount int
	applicants := applicant.GetApplicantsByOrderId(orderId)
	for i := range applicants {
		if applicants[i].Applicant.Status == "join" {
			payerCount++
		}
	}
	var averagePay float64 = order.TotalPrice / float64(payerCount)

	for i := range applicants {
		if applicants[i].Applicant.Status == "join" {
			//update applicant info
			applicants[i].Applicant.Pay = averagePay
			applicants[i].Applicant.UpdateToDB()
			//add new log
			var log dao.DAOAccountLog
			log.UserId = applicants[i].Applicant.UserId
			log.Type = "spend"
			log.Time = time.Now()
			log.Value = averagePay
			log.OrderId = applicants[i].Applicant.OrderId
			log.SaveToDB()
			//update user account amount
			applicants[i].User.AccountAmount -= averagePay
			applicants[i].User.UpdateToDB()
		}
	}
	return
}

func OrderClose(orderId int64) (order dao.DAOOrder) {
	order.FetchFromDB(orderId)
	order.Status = "closed"
	order.UpdateToDB()
	return
}

func GetRunningOrderList() (orders []Order) {
	daoOrders := dao.GetOrderList("pending", "running")
	for i := range daoOrders {
		var o Order
		o.Creator.FetchFromDB(daoOrders[i].UserId)
		o.Order = daoOrders[i]
		orders = append(orders, o)
	}
	return
}

func GetMyOrderList(userid int64) (orders []Order) {
	daoOrders := dao.GetMyOrderList(userid)
	for i := range daoOrders {
		var o Order
		o.Creator.FetchFromDB(daoOrders[i].UserId)
		o.Order = daoOrders[i]
		orders = append(orders, o)
	}
	return
}

func ChargeMoney(userid int64, num float64) (result int) {
	if num <= 0 {
		return 1
	}
	var user dao.DAOUser
	if err := user.FetchFromDB(userid); err != nil {
		panic(err.Error())
	}
	if user.Id <= 0 || userid != user.Id {
		return 0
	}
	user.AccountAmount += num
	if err := user.UpdateToDB(); err != nil {
		panic(err.Error())
	}
	var alog dao.DAOAccountLog
	alog.UserId = userid
	alog.Time = time.Now()
	alog.OrderId = -1
	alog.Type = "charge"
	alog.Value = num
	if err := alog.SaveToDB(); err != nil {
		panic(err.Error())
	}
	return 1
}
