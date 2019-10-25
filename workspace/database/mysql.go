package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func chaeckerr(err error){
	if err != nil{
		fmt.Println(err)
		return
	}
}


func main()  {
	bT := time.Now()
	db, err := sql.Open("mysql", "root:mysql@tcp(192.168.133.128:3306)/tset?charset=utf8")
	chaeckerr(err)
	stmt,err := db.Prepare("INSERT userinfo set uasername= ?,departname=?,created=?")
	chaeckerr(err)
	res,err := stmt.Exec("张三", "研发部门","2017-09-09")
	chaeckerr(err)
	id,err := res.LastInsertId()
	chaeckerr(err)
	fmt.Println(id)
	eT := time.Since(bT)
	println(eT)

}
