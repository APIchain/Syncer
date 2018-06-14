package db

import (
	"database/sql"
	"github.com/Syncer/common/log"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func clearTransaction(tx *sql.Tx) {
	err := tx.Rollback()
	if err != sql.ErrTxDone && err != nil {
		log.Fatal(err)
	}
}

//type DbWorker struct {
//	Dsn      string
//	Db       *sql.DB
//	UserInfo userTB
//}
//type userTB struct {
//	Id   int
//	Name sql.NullString
//	Age  sql.NullInt64
//}
//
//func main() {
//	var err error
//	dbw := DbWorker{
//		Dsn: "junjie:i3yap24k@tcp(35.198.225.135:3306)/etherdb?charset=utf8mb4",
//	}
//	dbw.Db, err = sql.Open("mysql", dbw.Dsn)
//	if err != nil {
//		panic(err)
//		return
//	}
//	defer dbw.Db.Close()
//
//	dbw.DoTx()
//}
//
//func (dbw *DbWorker) DoTx() {
//	stmt, _ := dbw.Db.Prepare(`INSERT INTO user (name, age) VALUES (?, ?)`)
//	defer stmt.Close()
//
//	ret, err := stmt.Exec("xys", 23)
//	if err != nil {
//		fmt.Printf("insert data error: %v\n", err)
//		return
//	}
//	if LastInsertId, err := ret.LastInsertId(); nil == err {
//		fmt.Println("LastInsertId:", LastInsertId)
//	}
//	if RowsAffected, err := ret.RowsAffected(); nil == err {
//		fmt.Println("RowsAffected:", RowsAffected)
//	}
//}
//func main(){
//	var str []string
//	log.Init(log.Path, log.Stdout)
//	str=append(str, "INSERT INTO user (name, age) VALUES ('zhang',17 )")
//	str=append(str, "INSERT INTO user (name, age) VALUES ('wang',18 );")
//	str=append(str, "INSERT INTO user (name, age) VALUES ('wang',18 )1;")
//
//	ExecBatch(str)
//}

func ExecBatch(txs []string) {
	db, err := sql.Open("mysql", "junjie:i3yap24k@tcp(35.198.225.135:3306)/etherdb?charset=utf8mb4")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer clearTransaction(tx)
	for _, v := range txs {
		fmt.Println("v=",v)
		_, err = tx.Exec(v)
		if err != nil {
			clearTransaction(tx)
			log.Warnf("[ExecBatch] failed,Error=%s,SQL=%s",err,txs)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		// tx.Rollback() 此时处理错误，会忽略doSomthing的异常
		log.Fatal(err)
	}

}
