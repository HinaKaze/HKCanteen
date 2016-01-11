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

type DAOUser struct {
	Id            int64
	Username      string
	Password      string
	NickName      string
	AccountAmount float64 //账户余额总额
}

func (r *DAOUser) SaveToDB(db *sql.DB) error {
	result, err := db.Exec("Insert into user (username,password,nickname) values (?,?,?)", r.Username, r.Password, r.NickName)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	r.Id = id
	return nil
}

func (r *DAOUser) UpdateToDB(db *sql.DB) error {
	result, err := db.Exec("Update user set username=?,password=?,nickname=? where id=?", r.Username, r.Password, r.NickName, r.Id)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (r *DAOUser) FetchFromDB(db *sql.DB, id int) error {
	rows, err := db.Query("select id,username,password,nickname from user where id=?", id)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&r.Id, &r.Username, &r.Password, &r.NickName)
		if err != nil {
			return err
		}
	}
	return nil
}

type DAOAccountLog struct {
	Id     int64
	UserId int64
	Type   string //withdraw,charge,spend...
	Number float64
	Time   time.Time
}

type DAOOrder struct {
	Id          int64
	UserId      int64 //下单者
	PayerId     int64 //付款人
	InventoryId int64
	Desc        string
	Status      string //pending,waiting,finished,closed
	TotalPrice  float64
	Time        time.Time
}

type DAOApplicant struct {
	Id      int64
	OrderId int64
	UserId  int64
	Status  string //join,hesitation,cancel
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
