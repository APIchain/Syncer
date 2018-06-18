package main

import (
	"fmt"
	"github.com/Syncer/tokenSync"
	//"github.com/Syncer/db"
	"os"
	"strconv"
	"github.com/Syncer/common/log"
)

func main() {
	log.Init(log.Path, log.Stdout)
	var NEO_START_HEIGHT = os.Getenv("NEO_START_HEIGHT")
	startHeight, _ := strconv.ParseInt(NEO_START_HEIGHT, 10, 64)

	worker := tokenSync.NewTokenFetcher()
	worker.Start()
	completed := make(chan tokenSync.TokenFetchResult, 1)
	for i := 3500000; i < 4000000; i++ {
		worker.FetchToken(i+int(startHeight), completed)
		tokenTransferResult := <-completed
		if tokenTransferResult.Rerr != nil {
			fmt.Println("tokenTransferResult failed,err=",tokenTransferResult.Rerr)
			return
		}
		var sqltxt []string
		fmt.Println("Block Height=",i)
		str:=tokenSync.GetTokenTransfer(tokenTransferResult.TokenTransfer)
		if str!=""{
			sqltxt=append(sqltxt, str)
		}
		if len(sqltxt)>0{
			fmt.Println(sqltxt)
			//db.ExecBatch(sqltxt)
		}
	}

}