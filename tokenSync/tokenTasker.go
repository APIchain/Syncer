package tokenSync

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Syncer/common/config"
	"github.com/Syncer/common/log"
	"io/ioutil"
	"net/http"
	"math/big"
	"github.com/Syncer/metadata"
)

type Task struct {
	Height int
	Result chan TokenFetchResult
}

type TokenFetchWoker struct {
	client http.Client
	task   chan Task
}

func NewTokenFetcher() TokenFetchWoker {
	worker := TokenFetchWoker{}
	worker.task = make(chan Task, 10)

	return worker
}

type TokenFetchResult struct {
	TokenTransfer *[]metadata.RpcEvent
	Rerr  error
}

func (self *TokenFetchWoker) FetchToken(height int, future chan TokenFetchResult) {
	task := Task{
		Height: height,
		Result: future,
	}
	self.task <- task
}

func (self *TokenFetchWoker) Start() {

	go func() {
		for {
			select {
			case task := <-self.task:
				if task.Height == -1 {
					task.Result <- TokenFetchResult{TokenTransfer: nil, Rerr: nil}
					return
				}

				tokenTransfer, err := self.gettokenTransfer(task.Height)

				task.Result <- TokenFetchResult{TokenTransfer: tokenTransfer, Rerr: err}

			}
		}
	}()
}

func (self *TokenFetchWoker) gettokenTransfer(height int) (*[]metadata.RpcEvent, error) {
	var rpcResp metadata.RemoteRpcResutToken
	data := fmt.Sprintf(`{"jsonrpc":"2.0","id":3156,"method":"eth_getLogs","params":[{"address":[],"fromBlock":"%s","toBlock":"%s","topics":[["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"]]}]}`,toBlockNumArg(big.NewInt(int64(height))),toBlockNumArg(big.NewInt(int64(height))))
	//fmt.Println("request data=",data)
	body := bytes.NewReader([]byte(data))
	url := fmt.Sprintf("http://%s:%s",config.Parameters.SyncServer , config.Parameters.SyncServerPort)
	resp, err := self.client.Post(url, "application/json", body)
	if err != nil {
		log.Errorf("http.Post, gettokenTransfer failed, height=%d,err=%s", height, err)
		return nil,err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("http.Post, gettokenTransfer failed, height=%d,err=%s", height, err)
		return nil,err
	}
	err = json.Unmarshal(buf, &rpcResp)
	if err != nil {
		log.Errorf("Unmarshal err=",err)
		return nil, err
	}
	return &rpcResp.Result, nil
}


