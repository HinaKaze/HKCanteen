package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HinaKaze/canteen/content"
	"github.com/HinaKaze/canteen/server/handler"
)

var Quit chan int

func main() {
	//	defer func() {
	//		if x := recover(); x != nil {
	//			log.Debugf("Reconnect from panic [%+v]", x)
	//			Init()
	//		}
	//	}()
	content.LoadServerContent()
	http.HandleFunc("/", handler.Welcome)
	http.HandleFunc("/sign_in", handler.Signin)
	http.HandleFunc("/sign_in_submit", handler.SigninSubmit)
	http.HandleFunc("/order_manage", handler.OrderManage)
	//http.HandleFunc("/meal_confirm", handler.MealConfirm)
	//http.HandleFunc("/meal_cancel", handler.MealCancel)
	//http.HandleFunc("/meal_list", handler.MealList)
	http.HandleFunc("/order_submit", handler.OrderCreate)
	http.HandleFunc("/order_list", handler.OrderList)
	http.HandleFunc("/order_own_list", handler.MyOrderList)

	http.HandleFunc("/order_join", handler.JoinOrder)
	http.HandleFunc("/order_quit", handler.QuitOrder)
	http.HandleFunc("/order_detail", handler.OrderDetail)
	http.HandleFunc("/order_run", handler.OrderRun)
	http.HandleFunc("/order_finish", handler.OrderFinish)
	http.HandleFunc("/order_close", handler.OrderClose)

	http.HandleFunc("/fuli_list", handler.FuliList)

	http.HandleFunc("/account_detail", handler.GetMyAccountLog)
	http.HandleFunc("/charge_view", handler.ChargeView)
	http.HandleFunc("/charge_money", handler.ChargeMoney)

	http.HandleFunc("/favicon.ico", handler.Favicon)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./client/img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./client/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./client/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./client/fonts"))))
	err := http.ListenAndServe(fmt.Sprintf(":%d", content.GetWebPort()), nil) //设置监听的端口
	if err != nil {
		panic(err.Error())
	}
	log.Println("WEB SERVER START")
	<-Quit
}
