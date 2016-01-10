package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var orderedAccouts map[interface{}]string = make(map[interface{}]string)
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

		if username == "admin" && password == "admin" {
			session, err := store.Get(r, "account")
			if err != nil {
				panic(err.Error())
			}
			session.Values["username"] = username
			session.Save(r, w)
			w.Write([]byte("登陆成功！！"))
		}
	}
}

func OrderMeal(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "account")
	if err != nil {
		panic(err.Error())
	}
	_, ok := session.Values["username"]
	if ok {
		t, _ := template.ParseFiles("./client/order_meal.html")
		t.Execute(w, nil)
	} else {
		w.Write([]byte("Nigger请登录"))
	}
}

func MealConfirm(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "account")
	if err != nil {
		panic(err.Error())
	}
	name, ok := session.Values["username"]
	if ok {
		orderedAccouts[name] = "ok"
		log.Printf("[%+v]\n", orderedAccouts)
		t, _ := template.ParseFiles("./client/order_meal.html")
		t.Execute(w, nil)
	} else {
		w.Write([]byte("Nigger请登录"))
	}

}

func MealCancel(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "account")
	if err != nil {
		panic(err.Error())
	}
	name, ok := session.Values["username"]
	if ok {
		delete(orderedAccouts, name)
		t, _ := template.ParseFiles("./client/order_meal.html")
		t.Execute(w, nil)
	} else {
		w.Write([]byte("Nigger请登录"))
	}
}

func MealList(w http.ResponseWriter, r *http.Request) {
	htmlStr := ""
	log.Printf("[%+v]\n", orderedAccouts)
	for a := range orderedAccouts {
		aStr, _ := a.(string)
		htmlStr += "<p>" + aStr + "</p>"
	}
	w.Write([]byte(htmlStr))
}

func Favicon(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(" "))
}
