package db

import (
	"database/sql"
	"github.com/Syncer/common/log"
	"github.com/Syncer/common/config"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func clearTransaction(tx *sql.Tx) {
	err := tx.Rollback()
	if err != sql.ErrTxDone && err != nil {
		log.Fatal(err)
	}
}

func ExecBatch(txs []string) {
	dsn:=fmt.Sprintf("junjie:i3yap24k@tcp(%s:3306)/etherdb?charset=utf8mb4",config.Parameters.SyncServer)
	//db, err := sql.Open("mysql", "junjie:i3yap24k@tcp(35.198.225.135:3306)/etherdb?charset=utf8mb4")
	db, err := sql.Open("mysql", dsn)
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
