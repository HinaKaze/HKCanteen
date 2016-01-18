package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"HKCanteen/server/biz/accountlog"
	"HKCanteen/server/biz/applicant"
	"HKCanteen/server/biz/order"
	"HKCanteen/server/biz/user"
	"HKCanteen/server/common"
)

func OrderManage(w http.ResponseWriter, r *http.Request) {
	if IsLogin(r) {
		t, _ := template.ParseFiles("./client/order.html")
		t.Execute(w, nil)
	} else {
		w.Write([]byte("Nigger请登录"))
	}
}

func OrderCreate(w http.ResponseWriter, r *http.Request) {
	if IsLogin(r) {
		if !CheckPrivilege(r, common.PrivilegeOrderCreate) {
			w.Write([]byte("你没有发起活动的权限"))
			return
		}
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

func MyOrderList(w http.ResponseWriter, r *http.Request) {
	if IsLogin(r) {
		orders := order.GetMyOrderList(GetUserId(r))
		t, _ := template.ParseFiles("./client/order_own_list.html")
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
				IsCreator bool
				Pending   bool
				Running   bool
				Closed    bool
				Finished  bool
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
		if resp.Order.Creator.Id == GetUserId(r) {
			resp.StatusFlag.IsCreator = true
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

func GetMyAccountLog(w http.ResponseWriter, r *http.Request) {
	if IsLogin(r) {
		accountLog := accountlog.GetMyAccountLog(GetUserId(r))
		t, err := template.ParseFiles("./client/account_log.html")
		if err != nil {
			panic(err.Error())
		}
		t.Execute(w, accountLog)
	} else {
		w.Write([]byte("Nigger请登录"))
	}
}

func ChargeView(w http.ResponseWriter, r *http.Request) {
	if IsLogin(r) {
		t, err := template.ParseFiles("./client/charge_money.html")
		if err != nil {
			panic(err.Error())
		}
		users := user.GetUserList()
		t.Execute(w, users)
	} else {
		w.Write([]byte("Nigger请登录"))
	}
}

func ChargeMoney(w http.ResponseWriter, r *http.Request) {
	if IsLogin(r) {
		if !CheckPrivilege(r, common.PrivilegeChargeMoney) {
			w.Write([]byte("少年你想多了，你没有权限"))
			return
		}
		err := r.ParseForm()
		if err != nil {
			panic(err.Error())
		}
		numStr := r.Form.Get("charge_num")
		chargeNum, err1 := strconv.ParseFloat(numStr, 64)
		userStr := r.Form.Get("charge_user")
		chargeUser, err2 := strconv.ParseInt(userStr, 10, 64)
		if err1 != nil || err2 != nil {
			w.Write([]byte("充值失败，输入参数错误"))
			return
		}
		result := order.ChargeMoney(chargeUser, chargeNum)
		if result > 0 {
			w.Write([]byte("充值成功"))
		} else {
			w.Write([]byte("充值失败"))
		}
	} else {
		w.Write([]byte("Nigger请登录"))
	}
}
