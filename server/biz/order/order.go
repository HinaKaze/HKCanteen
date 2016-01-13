package order

import (
	"log"
	"time"

	"HKCanteen/server/biz/applicant"
	"HKCanteen/server/dao"
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
	return
}

func OrderClose(orderId int64) (order dao.DAOOrder) {
	order.FetchFromDB(orderId)
	order.Status = "closed"
	order.UpdateToDB()
	return
}

func GetRunningOrderList() (orders []Order) {
	daoOrders := dao.GetOrderList("pending", "waiting")
	for i := range daoOrders {
		var o Order
		o.Creator.FetchFromDB(daoOrders[i].UserId)
		o.Order = daoOrders[i]
		orders = append(orders, o)
	}
	return
}
