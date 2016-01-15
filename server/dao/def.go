package dao

import (
	"database/sql"
	"time"
)

var canteenDBConn *sql.DB

func Init(db *sql.DB) {
	canteenDBConn = db
}

func GetDBConn() *sql.DB {
	return canteenDBConn
}

type DAOI interface {
	SaveToDB(db *sql.DB) error
	UpdateToDB(db *sql.DB) error
	FetchFromDB(db *sql.DB, id int) error
}

type DAOInventory struct {
	Id           int64
	Desc         string
	CommodityIds []int64
}

type DAOCompany struct {
	Id   int64
	Name string
	Desc string
}

type DAOCommodity struct {
	Id          int64
	CompanyId   int64
	Name        string
	Price       float64
	Desc        string
	SaleStatus  string //pending_sale,for_sale,closed
	EvaluteRank int    //评价等级（对顾客的评价取平均值）0-100
}

/*顾客对商品的评价*/
type DAOEvaluation struct {
	Id          int64
	CommodityId int64
	Score       int //0-100
	Comment     string
	Time        time.Time
}
