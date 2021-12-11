package model

import (
	"fmt"

	"mygo/Mysql/Code/task3/utils" //数据库初始化
)

type User struct {
	Uid   string
	Uname string
	Upwd  string
}

func CreatUser(name string, pwd string) error {
	sqlStr := "insert into live(userName,Pwd)values(?,?)"
	_, err := utils.DB.Exec(sqlStr, name, pwd)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}

func IsTrue(name string, pwd string) (string, string, error) {
	var (
		username string
		password string
		err      error
	)
	sqlStr := "selcet userName,Pwd from live where userName = ? and Pwd =?"
	row, _ := utils.DB.Query(sqlStr, name, pwd)
	for row.Next() {
		err = row.Scan(&username, &password)
	}
	if err != nil {
		return "", "", err
	}
	return username, password, nil
}

func Find(id string) (map[string]string, bool) {
	var (
		ID_tmp   int
		UserName string
		PassWord string
	)
	ok := false
	tmp := make(map[string]string, 2)
	sqlStr := "SELECT id,userName,Pwd FROM live"
	rows, err := utils.DB.Query(sqlStr)
	if err != nil {
		fmt.Println("查找失败", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&ID_tmp, &UserName, &PassWord)
		if err != nil {
			fmt.Println(err)
		}
		if id == ID_tmp {
			tmp["name"] = UserName
			tmp["password"] = PassWord
			ok = true
		}
	}
	return tmp, ok
}

//更新
func Update(id string, username string, password string) {
	sqlStr := "UPDATE live SET userName=?,Pwd=? WHERE id=?"
	infor, _ := Find(id)
	fmt.Println(infor)
	fmt.Println(password, username)
	if password == "" {
		password = infor["password"]
	}
	if username == "" {
		username = infor["name"]
	}
	_, err := utils.DB.Exec(sqlStr, username, password)
	if err != nil {
		fmt.Println("修改失败")
	} else {
		fmt.Println("修改成功")
	}
}

//查询数据
func QueryData(userName string, Pwd string, id string) {
	sqlStr := "SELECT *FROM live"
	rows, err := utils.DB.Query(sqlStr)
	if err != nil {
		fmt.Println("查找失败", err)
	}
	defer rows.Close()
	count := 0
	var data map[int]struct
	for rows.Next() {
		map[count]
		err = rows.Scan(&User.Uid, &User.Uname, &Upwd)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

}
