package metadata

import "github.com/gin-gonic/gin/json"

type RemoteRpcResutEth struct {
	Jsonrpc string
	Id      int
	Result  RpcBlockJson
}



type RpcBlockJson struct {
	Author           string           `"&json:author&"`
	Difficulty       string           `"&json:difficulty&"`
	ExtraData        string           `"&json:extraData&"`
	GasLimit         string           `"&json:gasLimit&"`
	GasUsed          string           `"&json:gasUsed&"`
	Hash             string           `"&json:hash&"`
	LogsBloom        string           `"&json:logsBloom&"`
	Miner            string           `"&json:miner&"`
	MixHash          string           `"&json:mixHash&"`
	Nonce            string           `"&json:nonce&"`
	Number           string           `"&json:number&"`
	ParentHash       string           `"&json:parentHash&"`
	ReceiptsRoot     string           `"&json:receiptsRoot&"`
	SealFields       []string         `"&json:sealFields&"`
	Sha3Uncles       string           `"&json:sha3Uncles&"`
	Size             string           `"&json:size&"`
	StateRoot        string           `"&json:stateRoot&"`
	Timestamp        string           `"&json:timestamp&"`
	TotalDifficulty  string           `"&json:totalDifficulty&"`
	TransactionsRoot string           `"&json:transactionsRoot&"`
	Transactions     []rpcTransaction `json:"transactions"`
}

func (this *RpcBlockJson) Marshal() ([]byte, error) {
	str, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return str, nil
}

type rpcTransaction struct {
	BlockHash        string `"&json:blockHash&"`
	BlockNumber      string `"&json:blockNumber&"`
	ChainId          string `"&json:chainId&"`
	Condition        string `"&json:condition&"`
	Creates          string `"&json:creates&"`
	From             string `"&json:from&"`
	Gas              string `"&json:gas&"`
	GasPrice         string `"&json:gasPrice&"`
	Hash             string `"&json:hash&"`
	Input            string `"&json:input&"`
	Nonce            string `"&json:nonce&"`
	PublicKey        string `"&json:publicKey&"`
	R                string `"&json:r&"`
	Raw              string `"&json:raw&"`
	S                string `"&json:s&"`
	StandardV        string `"&json:standardV&"`
	To               string `"&json:to&"`
	TransactionIndex string `"&json:transactionIndex&"`
	V                string `"&json:v&"`
	Value            string `"&json:value&"`
}