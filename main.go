package main

import (
	"fmt"
	"github.com/Syncer/sync"
	"github.com/Syncer/db"
	"os"
	"strconv"
	"github.com/Syncer/common/log"
	"github.com/Syncer/tokenSync"
)

var Tokencompleted chan tokenSync.TokenFetchResult
var Ethcompleted chan sync.BlockFetchResult


func main() {
	log.Init(log.Path, log.Stdout)
	var NEO_START_HEIGHT = os.Getenv("NEO_START_HEIGHT")
	startHeight, _ := strconv.ParseInt(NEO_START_HEIGHT, 10, 64)
	// Ether blockFetchWoker
	blockFetchWoker := sync.NewBlockFetcher()
	blockFetchWoker.Start()
	Ethcompleted = make(chan sync.BlockFetchResult, 1)

	//tokenWorker
	tokenWorker := tokenSync.NewTokenFetcher()
	tokenWorker.Start()
	Tokencompleted = make(chan tokenSync.TokenFetchResult, 1)
	for i := 0; i < 4000000; i++ {
		var sqltxt []string
		//Get Ether
		result1:=etherHandler(blockFetchWoker,i+int(startHeight))
		if result1!=nil || len(result1)>0{
			sqltxt=append(sqltxt, result1...)
		}
		//Get Token
		result2:=tokenHandler(tokenWorker,i+int(startHeight))
		if result2!=nil || len(result2)>0{
			sqltxt=append(sqltxt, result2...)
		}
		log.Infof("[Height %d] will inset %d row",i+int(startHeight),len(sqltxt))
		if len(sqltxt)>0{
			for _, v := range sqltxt {
				log.Infof("%s",v)
			}
			db.ExecBatch(sqltxt)
		}
	}
}

func etherHandler(blockFetchWoker sync.BlockFetchWoker,startHeight int)[]string{
	blockFetchWoker.FetchBlock(int(startHeight), Ethcompleted)
	blockResult := <-Ethcompleted
	if blockResult.Rerr != nil {
		fmt.Println(blockResult.Rerr)
		return nil
	}
	var sqltxt []string
	if len(blockResult.Block.Transactions) > 0{
		str:=sync.GetEtherTransfer(blockResult.Block)
		if str!=nil{
			sqltxt=append(sqltxt, str...)
		}
	}
	return sqltxt
}

func tokenHandler(worker tokenSync.TokenFetchWoker,startHeight int) []string{
	worker.FetchToken(int(startHeight), Tokencompleted)
	tokenTransferResult := <-Tokencompleted
	if tokenTransferResult.Rerr != nil {
		log.Error("tokenTransferResult failed,err=",tokenTransferResult.Rerr)
		return nil
	}
	var sqltxt []string
	str:=tokenSync.GetTokenTransfer(tokenTransferResult.TokenTransfer)
	if str!=nil{
		sqltxt=append(sqltxt, str...)
	}
	return sqltxt
}