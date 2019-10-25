// go 操作mysql
package main

import (
	"database/sql"
	"fmt"
	//"time"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	DB, err := sql.Open("mysql","root:mysql@tcp(192.168.160.130:3306)/test?charset=utf8 ")
	checkErr(err)
	// 插入数据
	stmt, err := DB.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	//更新数据
	stmt,err = DB.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//查询数据
	rows, err := DB.Query("select * from userinfo")
	checkErr(err)

	for rows.Next(){
		var uid int
		var usernae string
		var department string
		var created string
		err = rows.Scan(&uid, &usernae, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(usernae)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = DB.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect,err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	DB.Close()

}

func checkErr(err error)  {
	if err != nil{
		panic(err)
	}
}