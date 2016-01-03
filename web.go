// frgm
package main

import (
	"fmt"
	"log"
	"net/http"

	"blog/content"
	"blog/server/handler"
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
	//defer conn.Close()
}
