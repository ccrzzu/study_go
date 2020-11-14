package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func insertData(db *sql.DB) {
	stmt, _ := db.Prepare(`INSERT INTO person (id, name, age, ts) VALUES (?, ?, ?, ?)`)

	rows, err := stmt.Query(0, "cui", 20, time.Now().UnixNano()/1e6)
	time.Sleep(time.Duration(1) * time.Second)
	defer stmt.Close()

	for i := 1; i < 100; i++ {
		rows, err := stmt.Query(i, "cui", 20, time.Now().UnixNano()/1e6)
		if err != nil {
			log.Fatalf("insert data error: %v\n", err)
		}
		rows.Close()
		time.Sleep(time.Duration(6) * time.Second)
	}
	if err != nil {
		log.Fatalf("insert data error: %v\n", err)
	}

	//var result int
	////rows, err = stmt.Query(2, "test", 19)
	//rows.Scan(&result)
	log.Printf("insert result %v\n", "ok")
	rows.Close()
}

func main() {
	//第⼀步：打开数据库,格式是 ⽤户名：密码@/数据库名称？编码⽅式
	db, err := sql.Open("mysql", "rtic_test_db_pro:82jxlsem505q9s5s57ermb@tcp(rtic-test-db.cgn9ujqteylm.us-west-1.rds.amazonaws.com:3306)/rtic_test_db?charset=utf8")
	//关闭数据库
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//real exe
	insertData(db)

}
