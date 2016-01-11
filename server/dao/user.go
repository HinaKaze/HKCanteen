package dao

func GetUserByUsername(username string) (user DAOUser) {
	sql := `select id,username,password,nickname from user where username=? `
	rows, err := canteenDBConn.Query(sql, username)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.NickName)
		if err != nil {
			panic(err.Error())
		}
	}
	return
}
