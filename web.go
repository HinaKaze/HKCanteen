// frgm
package main

import (
	"fmt"
	//"encoding/json"
	//"errors"
	//"fmt"
	// "github.com/gorilla/mux"
	"html/template"

	"hkweb/content"

	"github.com/gorilla/securecookie"
	// "log"
	//"net"
	//"net/http"
	"palmjoi.com/fantasyrealm/log"
	//"strconv"
	"net/http"
	//"time"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

var Quit chan int

func main() {
	//	defer func() {
	//		if x := recover(); x != nil {
	//			log.Debugf("Reconnect from panic [%+v]", x)
	//			Init()
	//		}
	//	}()
	content.LoadServerContent()
	http.HandleFunc("/", welcome)
	//	http.HandleFunc("/player", PlayerExec)
	//	http.HandleFunc("/pquery", PQueryExec)
	//	http.HandleFunc("/pmodify", PModifyExec)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./client/img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./client/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./client/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./client/fonts"))))
	err := http.ListenAndServe(fmt.Sprintf(":%d", content.GetWebPort()), nil) //设置监听的端口
	if err != nil {
		panic(err.Error())
	}
	log.Debugf("WEB SERVER START")
	<-Quit
	//defer conn.Close()
}

//func PlayerExec(w http.ResponseWriter, r *http.Request) {

//	r.ParseForm()
//	response := PlayerHandle(r.FormValue("code"), r.FormValue("pid"))
//	w.Write([]byte(response))
//}
//func PModifyExec(w http.ResponseWriter, r *http.Request) {

//	r.ParseForm()
//	response := PModifyHandle(r.FormValue("code"), r.FormValue("pid"), r.FormValue("num"))
//	w.Write([]byte(response))
//}
//func PQueryExec(w http.ResponseWriter, r *http.Request) {
//	//	defer func() {
//	//		if x := recover(); x != nil {
//	//			log.Debugf("Reconnect from panic [%+v]", x)
//	//			Init()
//	//		}
//	//	}()
//	r.ParseForm()
//	response := PQueryHandle(r.FormValue("code"), r.FormValue("param"))
//	w.Write([]byte(response))
//}

//func getUserName(request *http.Request) (userName string) {
//	//	if cookie, err := request.Cookie("session"); err == nil {
//	//		cookieValue := make(map[string]string)
//	//		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
//	//			userName = cookieValue["name"]
//	//		}
//	//	}
//	return userName
//}

//func setSession(userName string, response http.ResponseWriter) {
//	value := map[string]string{
//		"name": userName,
//	}
//	if encoded, err := cookieHandler.Encode("session", value); err == nil {
//		cookie := &http.Cookie{
//			Name:  "session",
//			Value: encoded,
//			Path:  "/",
//		}
//		http.SetCookie(response, cookie)
//	}
//}

//func clearSession(response http.ResponseWriter) {
//	cookie := &http.Cookie{
//		Name:   "session",
//		Value:  "",
//		Path:   "/",
//		MaxAge: -1,
//	}
//	http.SetCookie(response, cookie)
//}

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./client/index.html")
	log.Debugf("welcome invoke")
	t.Execute(w, nil)
}
