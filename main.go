package main

import (
	"fmt"
	"github.com/Syncer/sync"
	"github.com/Syncer/db"
	"os"
	"strconv"
	"github.com/Syncer/common/log"
)

func main() {
	log.Init(log.Path, log.Stdout)
	var NEO_START_HEIGHT = os.Getenv("NEO_START_HEIGHT")
	startHeight, _ := strconv.ParseInt(NEO_START_HEIGHT, 10, 64)

	worker := sync.NewBlockFetcher()
	worker.Start()
	completed := make(chan sync.BlockFetchResult, 1)
	for i := 0; i < 1000000; i++ {
		worker.FetchBlock(i+int(startHeight), completed)

		blockResult := <-completed
		if blockResult.Rerr != nil {
			fmt.Println(blockResult.Rerr)
			return
		}
		//result := ProcessBlock(blockResult.block)
		//buf, _ := json.Marshal(result)
		//str,err:=blockResult.Block.Marshal()
		//if err != nil {
		//	fmt.Println(err)
		//}
		//fmt.Println(string(str))
		fmt.Println("blockResult.Block.Number=",sync.HexoToInt(blockResult.Block.Number))
		var sqltxt []string
		if len(blockResult.Block.Transactions) > 0{
			str:=sync.GetEtherTransfer(blockResult.Block)
			if str!=""{
				sqltxt=append(sqltxt, str)
			}
		}
		if len(sqltxt)>0{
			db.ExecBatch(sqltxt)
		}
	}

}