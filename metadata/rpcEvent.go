package metadata

type RemoteRpcResutToken struct {
	Jsonrpc string
	Id      int
	Result  []RpcEvent
}

type RpcEvent struct {
	// Consensus fields:
	// address of the contract that generated the event
	Address string `json:"address" gencodec:"required"`
	// list of topics provided by the contract.
	Topics []string `json:"topics" gencodec:"required"`
	// supplied by the contract, usually ABI-encoded
	Data string `json:"data" gencodec:"required"`

	// Derived fields. These fields are filled in by the node
	// but not secured by consensus.
	// block in which the transaction was included
	BlockNumber string `json:"blockNumber"`
	// hash of the transaction
	TxHash string `json:"transactionHash" gencodec:"required"`
	// index of the transaction in the block
	TxIndex string `json:"transactionIndex" gencodec:"required"`
	// hash of the block in which the transaction was included
	BlockHash string `json:"blockHash"`
	// index of the log in the receipt
	Index string `json:"logIndex" gencodec:"required"`
}
