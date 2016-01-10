package content

import (
	"blog/tools/iniparse"
	"database/sql"
	//	"fmt"
)

var serverContent struct {
	WebPort   int
	DBContent dbContent
	DBConn    *sql.DB
}

type dbContent struct {
	Address  string
	User     string
	Password string
	Port     string
	Schema   string
}

func LoadServerContent() {
	iniparse.DefaultParse("./content/config.ini")
	s, ok := iniparse.GetSection("WEB")
	if ok {
		serverContent.WebPort = s.GetIntValue("webPort")
	}
	//	s, ok = iniparse.GetSection("DB")
	//	if ok {
	//		serverContent.DBContent.Address, _ = s.GetValue("address")
	//		serverContent.DBContent.User, _ = s.GetValue("user")
	//		serverContent.DBContent.Password, _ = s.GetValue("password")
	//		serverContent.DBContent.Port, _ = s.GetValue("port")
	//		serverContent.DBContent.Schema, _ = s.GetValue("schema")
	//		dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", serverContent.DBContent.User, serverContent.DBContent.Password, serverContent.DBContent.Address, serverContent.DBContent.Port, serverContent.DBContent.Schema)
	//		dbt, err := sql.Open("mysql", dbConnectionString)
	//		if err != nil {
	//			panic(err.Error())
	//		}
	//		serverContent.DBConn = dbt
	//	}
}

func GetWebPort() int {
	return serverContent.WebPort
}
