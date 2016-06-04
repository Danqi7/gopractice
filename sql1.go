package main

import (
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:&l/t4uDgBz2g@/test")
	checkErr(err)
	
	//insert
	stmt, err := db.Prepare("INSERT userinfo SET username=?, departname=?,created=?")
	checkErr(err)
	
	res, err := stmt.Exec("Danqi", "GoogleX", "2018-6-1")
	checkErr(err)
	
	id, err := res.LastInsertId()
	checkErr(err)
	
	fmt.Println(id)

	//update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("danqiliao", id)
	checkErr(err)
	
	affect, err := res.RowsAffected()
	checkErr(err)
	
	fmt.Println(affect)

	//query
	rows, err := db.Query("SELECT * FROM userinfo")	
	checkErr(err)
	
	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created string
		err = rows.Scan(&uid, &username, &departname, &created)	
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(departname)
		fmt.Println(created)	
	}
	
	//delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	

	res, err = stmt.Exec(id)
	checkErr(err)
	
	affect, err = res.RowsAffected()
	checkErr(err)
	
	fmt.Println(affect)
	
	db.Close()		 
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}
