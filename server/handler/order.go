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

func GetApplicants(w http.ResponseWriter, r *http.Request) {
	if IsLogin(r) {
		err := r.ParseForm()
		if err != nil {
			panic(err.Error())
		}
		orderIdStr := r.Form.Get("orderid")
		orderId, _ := strconv.Atoi(orderIdStr)
		applicants := applicant.GetApplicantsByOrderId(int64(orderId))
		t, _ := template.ParseFiles("./client/applicant_list.html")
		t.Execute(w, applicants)
	} else {
		w.Write([]byte("Nigger请登录"))
	}
}
