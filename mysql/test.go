package main

import (
	"database/sql"
	"fmt"

	//time
	_ "github.com/Go-SQL-Driver/MYSQL"
)

func main() {
	db, err := sql.Open("mysql", "")
	checkErr(err)

	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?,uid=?")
	checkErr(err)

	res, err := stmt.Exec("Tom", "研发部门", "2012-12-09", "0")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("插入记录id：")
	fmt.Println(id)

	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("lisa", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("受影响行数：")
	fmt.Println(affect)

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	/*stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)*/

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
