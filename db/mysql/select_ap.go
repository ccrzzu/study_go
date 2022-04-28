package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func selectData1(db *sql.DB, id int, c chan int64) {
	row := db.QueryRow(`SELECT ts From person where id = ?`, id)
	var ts int64
	row.Scan(&ts)

	for ts == 0 {
		row = db.QueryRow(`SELECT ts From person where id = ?`, id)
		row.Scan(&ts)
	}

	now := time.Now().UnixNano() / 1e6
	duration := now - ts
	fmt.Println("id:", id, "now:", now, "ts:", ts, "duration:", duration)
	c <- duration
	//fmt.Println(rows.Err())
	//for rows.Next() {
	//	fmt.Println("end in")
	//	err = rows.Scan(&ts)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	now := time.Now().UnixNano() / 1e6
	//	duration := now - ts
	//	fmt.Println("id:", id, "now:", now, "ts:", ts, "duration:", duration)
	//	c <- duration
	//}
	//rows.Close()
}

func consumer1(c, f chan int64) {
	var sum int64
	var i int
	for i < 100 {
		if v, ok := <-c; ok {
			sum += v // 阻塞，直到生产者放入数据后继续读取数据
			//fmt.Println("sum:", sum)
			i++
		} else {
			continue
		}
	}
	f <- sum //发送数据，通知main函数已接受完成
}

func main() {
	//第⼀步：打开数据库,格式是 ⽤户名：密码@/数据库名称？编码⽅式
	db, err := sql.Open("mysql", "rtic_test_db_pro:82jxlsem505q9s5s57ermb@tcp(rtic-test-db-ap.c6zuo4pslptg.ap-southeast-1.rds.amazonaws.com:3306)/rtic_test_db?charset=utf8")
	//关闭数据库
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	buf := make(chan int64, 100)
	flg := make(chan int64)

	//real exe

	for i := 0; i < 100; i++ {
		go selectData1(db, i, buf)
	}
	go consumer1(buf, flg)

	//var str string
	//fmt.Scan(&str)
	sum, ok := <-flg //等待接受完成
	if ok {
		fmt.Println("sum:", sum, "avg ms:", sum/100)
	}

}
