package handler

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"HKCanteen/server/biz/user"
	"HKCanteen/server/common"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("23333"))

func Welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./client/index.html")
	t.Execute(w, nil)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "account")
	if err != nil {
		panic(err.Error())
	}
	_, ok := session.Values["username"]
	if ok {
		w.Write([]byte("nigger你已经登陆了"))
	} else {
		t, _ := template.ParseFiles("./client/signin.html")
		t.Execute(w, nil)
	}
}

func SigninSubmit(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "account")
	if err != nil {
		panic(err.Error())
	}
	_, ok := session.Values["username"]
	if ok {
		w.Write([]byte("nigger你已经登陆了"))
	} else {
		err := r.ParseForm()
		if err != nil {
			panic(err.Error())
		}
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		myself, result := user.Login(username, password)
		switch result {
		case user.ResultSuccess:
			session, err := store.Get(r, "account")
			if err != nil {
				panic(err.Error())
			}
			session.Values["username"] = myself.Username
			session.Values["userid"] = myself.Id
			session.Values["nickname"] = myself.NickName
			session.Values["privilege"] = myself.Privilege
			log.Println("Pri" + strconv.Itoa(myself.Privilege))
			session.Save(r, w)
			w.Write([]byte("登陆成功！！"))
		case user.ResultNotMatch:
			w.Write([]byte("账户密码不匹配！！"))
		case user.ResultUserInvalid:
			w.Write([]byte("用户不存在！！"))
		default:
			w.Write([]byte("未知错误！！"))
		}
	}
}

//func MealConfirm(w http.ResponseWriter, r *http.Request) {
//	session, err := store.Get(r, "account")
//	if err != nil {
//		panic(err.Error())
//	}
//	name, ok := session.Values["username"]
//	if ok {
//		orderedAccouts[name] = "ok"
//		log.Printf("[%+v]\n", orderedAccouts)
//		t, _ := template.ParseFiles("./client/order_meal.html")
//		t.Execute(w, nil)
//	} else {
//		w.Write([]byte("Nigger请登录"))
//	}

//}

//func MealCancel(w http.ResponseWriter, r *http.Request) {
//	session, err := store.Get(r, "account")
//	if err != nil {
//		panic(err.Error())
//	}
//	name, ok := session.Values["username"]
//	if ok {
//		delete(orderedAccouts, name)
//		t, _ := template.ParseFiles("./client/order_meal.html")
//		t.Execute(w, nil)
//	} else {
//		w.Write([]byte("Nigger请登录"))
//	}
//}

//func MealList(w http.ResponseWriter, r *http.Request) {
//	htmlStr := ""
//	log.Printf("[%+v]\n", orderedAccouts)
//	for a := range orderedAccouts {
//		aStr, _ := a.(string)
//		htmlStr += "<p>" + aStr + "</p>"
//	}
//	w.Write([]byte(htmlStr))
//}

func FuliList(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "account")
	if err != nil {
		panic(err.Error())
	}
	_, ok := session.Values["username"]
	if ok {
		t, _ := template.ParseFiles("./client/fuli_list.html")
		t.Execute(w, nil)
	} else {
		w.Write([]byte("Nigger请登录"))
	}
}

func Favicon(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(" "))
}

func IsLogin(r *http.Request) bool {
	session, err := store.Get(r, "account")
	if err != nil {
		panic(err.Error())
	}
	_, ok := session.Values["username"]
	return ok
}

func CheckPrivilege(r *http.Request, p common.PrivilegeSignal) bool {
	session, err := store.Get(r, "account")
	if err != nil {
		panic(err.Error())
	}
	pi, ok := session.Values["privilege"]
	if !ok {
		return false
	}
	pnum, ok := pi.(int64)
	if !ok {
		return false
	}
	privilege := common.PrivilegeSignal(pnum)
	if privilege&p > 0 {
		return true
	}
	return false
}

func GetUserId(r *http.Request) int64 {
	session, err := store.Get(r, "account")
	if err != nil {
		panic(err.Error())
	}
	val, _ := session.Values["userid"]
	userid, _ := val.(int64)
	return userid
}
