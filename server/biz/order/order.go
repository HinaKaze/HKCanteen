package order

import (
	"time"

	"HKCanteen/server/dao"
)

type Order struct {
	Order   dao.DAOOrder
	Creator dao.DAOUser
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

func ChangeOrderStatus(orderId int64, status string) (order dao.DAOOrder) {
	order.FetchFromDB(orderId)
	order.Status = status
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
