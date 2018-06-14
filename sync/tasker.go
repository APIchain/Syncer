package sync

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Syncer/common/config"
	"io/ioutil"
	"net/http"
	"math/big"
	"github.com/Syncer/metadata"
)

type Task struct {
	Height int
	Result chan BlockFetchResult
}

type BlockFetchWoker struct {
	client http.Client
	task   chan Task
}

func NewBlockFetcher() BlockFetchWoker {
	worker := BlockFetchWoker{}
	worker.task = make(chan Task, 10)

	return worker
}

type BlockFetchResult struct {
	Block metadata.BlockJson
	Rerr  error
}

func (self *BlockFetchWoker) FetchBlock(height int, future chan BlockFetchResult) {
	task := Task{
		Height: height,
		Result: future,
	}
	self.task <- task
}

func (self *BlockFetchWoker) Start() {

	go func() {
		for {
			select {
			case task := <-self.task:
				if task.Height == -1 {
					task.Result <- BlockFetchResult{Block: metadata.BlockJson{}, Rerr: nil}
					return
				}

				blockJson, err := self.getBlock(task.Height)

				task.Result <- BlockFetchResult{Block: blockJson, Rerr: err}

			}
		}
	}()
}

func (self *BlockFetchWoker) getBlock(height int) (metadata.BlockJson, error) {
	var rpcResp metadata.RemoteRpcResut
	data := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["%s",true],"id":1}`,toBlockNumArg(big.NewInt(int64(height))))
	body := bytes.NewReader([]byte(data))
	url := fmt.Sprintf("http://%s:%s",config.Parameters.SyncServer , config.Parameters.SyncServerPort)
	resp, err := self.client.Post(url, "application/json", body)
	if err != nil {
		fmt.Println("http.Post, getblock failed, height", 10000, err)
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http.Post, getblock failed, height", 10000, err)
	}
	//fmt.Printf("%s\n",buf)

	err = json.Unmarshal(buf, &rpcResp)
	if err != nil {
		return metadata.BlockJson{}, err
	}
	return rpcResp.Result, nil
}


