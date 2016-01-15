package dao

type DAOUser struct {
	Id            int64
	Username      string
	Password      string
	NickName      string
	AccountAmount float64 //账户余额总额
	Privilege     int     //
}

func (r *DAOUser) SaveToDB() error {
	result, err := GetDBConn().Exec("Insert into user (username,password,nickname,accountamount,privilege) values (?,?,?,?,?)", r.Username, r.Password, r.NickName, r.AccountAmount, r.Privilege)
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

func (r *DAOUser) UpdateToDB() error {
	result, err := GetDBConn().Exec("Update user set username=?,password=?,nickname=?,accountamount=?,privilege=? where id=?", r.Username, r.Password, r.NickName, r.AccountAmount, r.Privilege, r.Id)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (r *DAOUser) FetchFromDB(id int64) error {
	rows, err := GetDBConn().Query("select id,username,password,nickname,accountamount,privilege from user where id=?", id)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&r.Id, &r.Username, &r.Password, &r.NickName, &r.AccountAmount, &r.Privilege)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetUserByUsername(username string) (user DAOUser) {
	sql := `select id,username,password,nickname,accountamount,privilege from user where username=? `
	rows, err := canteenDBConn.Query(sql, username)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.NickName, &user.AccountAmount, &user.Privilege)
		if err != nil {
			panic(err.Error())
		}
	}
	return
}
