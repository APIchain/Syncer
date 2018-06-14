package sync

import (
	"github.com/Syncer/metadata"
	"fmt"
)

func GetEtherTransfer(block metadata.BlockJson) string{
	var result string
	for _, v := range block.Transactions {
		realData := &metadata.TxTransfer{
			TxHash:    v.Hash,
			Timestamp: hexoToInt(block.Timestamp),
			Height:    hexoToInt(v.BlockNumber),
			From:      v.From,
			To:        v.To,
			Value:     hexoToString(v.Value),
		}
		return  fmt.Sprintf("%s\n%s",result,realData.Marshal())
	}
	return result
}
