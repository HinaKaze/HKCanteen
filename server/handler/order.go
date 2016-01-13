package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"HKCanteen/server/biz/applicant"
	"HKCanteen/server/biz/order"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	if IsLogin(r) {
		err := r.ParseForm()
		if err != nil {
			panic(err.Error())
		}
		desc := r.Form.Get("desc")
		userId := GetUserId(r)
		o := order.CreateOrder(userId, desc)
		if o.Id < 0 {
			panic("创建订单失败")
		}
		w.Write([]byte{})
	} else {
		w.Write([]byte("Nigger请登录"))
	}

}

func OrderList(w http.ResponseWriter, r *http.Request) {
	if IsLogin(r) {
		orders := order.GetRunningOrderList()
		t, _ := template.ParseFiles("./client/order_list.html")
		t.Execute(w, orders)
	} else {
		w.Write([]byte("Nigger请登录"))
	}
}

func JoinOrder(w http.ResponseWriter, r *http.Request) {
	if IsLogin(r) {
		err := r.ParseForm()
		if err != nil {
			panic(err.Error())
		}
		orderIdStr := r.Form.Get("orderid")
		orderId, _ := strconv.Atoi(orderIdStr)
		applicant.JoinOrder(GetUserId(r), int64(orderId))
		w.Write([]byte{})
	} else {
		w.Write([]byte("Nigger请登录"))
	}
}

func QuitOrder(w http.ResponseWriter, r *http.Request) {
	if IsLogin(r) {
		err := r.ParseForm()
		if err != nil {
			panic(err.Error())
		}
		orderIdStr := r.Form.Get("orderid")
		orderId, _ := strconv.Atoi(orderIdStr)
		applicant.QuitOrder(GetUserId(r), int64(orderId))
		w.Write([]byte{})
	} else {
		w.Write([]byte("Nigger请登录"))
	}
}

func OrderDetail(w http.ResponseWriter, r *http.Request) {
	if IsLogin(r) {
		err := r.ParseForm()
		if err != nil {
			panic(err.Error())
		}
		orderIdStr := r.Form.Get("orderid")
		orderId, _ := strconv.Atoi(orderIdStr)

		var resp struct {
			Order      order.Order
			StatusFlag struct {
				Pending  bool
				Running  bool
				Closed   bool
				Finished bool
			}
		}
		resp.Order = order.GetOrderDetail(int64(orderId))
		if resp.Order.Order.Status == "pending" {
			resp.StatusFlag.Pending = true
		} else if resp.Order.Order.Status == "running" {
			resp.StatusFlag.Running = true
		} else if resp.Order.Order.Status == "closed" {
			resp.StatusFlag.Closed = true
		} else if resp.Order.Order.Status == "finished" {
			resp.StatusFlag.Finished = true
		}

		t, err := template.ParseFiles("./client/order_detail.html")
		if err != nil {
			panic(err.Error())
		}
		t.Execute(w, resp)
	} else {
		w.Write([]byte("Nigger请登录"))
	}
}

func OrderRun(w http.ResponseWriter, r *http.Request) {
	if IsLogin(r) {
		err := r.ParseForm()
		if err != nil {
			panic(err.Error())
		}
		orderIdStr := r.Form.Get("orderid")
		orderId, _ := strconv.ParseInt(orderIdStr, 10, 64)
		totalPriceStr := r.Form.Get("totalprice")
		totalPrice, _ := strconv.ParseFloat(totalPriceStr, 64)
		order.OrderRun(orderId, totalPrice)
		w.Write([]byte{})
	} else {
		w.Write([]byte("Nigger请登录"))
	}
}

func OrderFinish(w http.ResponseWriter, r *http.Request) {
	if IsLogin(r) {
		err := r.ParseForm()
		if err != nil {
			panic(err.Error())
		}
		orderIdStr := r.Form.Get("orderid")
		orderId, _ := strconv.Atoi(orderIdStr)
		order.OrderFinish(int64(orderId))
		w.Write([]byte{})
	} else {
		w.Write([]byte("Nigger请登录"))
	}
}

func OrderClose(w http.ResponseWriter, r *http.Request) {
	if IsLogin(r) {
		err := r.ParseForm()
		if err != nil {
			panic(err.Error())
		}
		orderIdStr := r.Form.Get("orderid")
		orderId, _ := strconv.Atoi(orderIdStr)
		order.OrderClose(int64(orderId))
		w.Write([]byte{})
	} else {
		w.Write([]byte("Nigger请登录"))
	}
}
