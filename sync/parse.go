package sync

import (
	"github.com/Syncer/metadata"
)

var BlockTime int64

func GetEtherTransfer(block metadata.RpcBlockJson) []string{
	var result []string
	BlockTime = HexoToInt(block.Timestamp)
	for _, v := range block.Transactions {
		realData := &metadata.SQLTransfer{
			TxHash:    v.Hash,
			Timestamp: HexoToInt(block.Timestamp),
			Height:    HexoToInt(v.BlockNumber),
			From:      v.From,
			To:        v.To,
			Value:     hexoToString(v.Value),
		}
		if realData.Value == "0" || realData.To==""{
			continue
		}
		result = append(result, realData.Marshal())
	}
	return result
}
