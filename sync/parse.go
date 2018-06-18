package sync

import (
	"github.com/Syncer/metadata"
	"fmt"
)

func GetEtherTransfer(block metadata.RpcBlockJson) string{
	var result string
	for _, v := range block.Transactions {
		realData := &metadata.SQLTransfer{
			TxHash:    v.Hash,
			Timestamp: HexoToInt(block.Timestamp),
			Height:    HexoToInt(v.BlockNumber),
			From:      v.From,
			To:        v.To,
			Value:     hexoToString(v.Value),
		}
		if v.Value == "0" || v.To==""{
			continue
		}
		return  fmt.Sprintf("%s\n%s",result,realData.Marshal())
	}
	return result
}
