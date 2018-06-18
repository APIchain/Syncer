package tokenSync

import (
	"github.com/Syncer/metadata"
	"fmt"
	"github.com/Syncer/common/log"
	"github.com/Syncer/sync"
)

func GetTokenTransfer(logs *[]metadata.RpcEvent) []string{
	if len(*logs)< 1{
		return nil
	}
	var result []string
	for _, v := range *logs {
		var from, to, value string
		var err error
		if len(v.Topics) < 3 {
			if len(v.Topics) == 1 && len(v.Data) == 32*3 {
				detail := fmt.Sprintf("%x", v.Data)
				from, err = marchalAddr(detail[:64])
				if err != nil {
					log.Errorf("[marchalAddr] from failed with hash=%s,err=%s", v.TxHash, err)
					continue
				}
				to, err = marchalAddr(detail[64:128])
				if err != nil {
					log.Errorf("[marchalAddr] from failed with hash=%s,err=%s", v.TxHash, err)
					continue
				}
				value, err = marchalInt(detail[128:])
				if err != nil {
					log.Errorf("[marchalInt2] value failed with hash=%s,err=%s", v.TxHash, err)
					continue
				}
			} else {
				log.Errorf("error with hash=%s topic=%d,data=%d", v.TxHash, len(v.Topics), len(v.Data))
				continue
			}
		} else {
			from, err = marchalAddr(v.Topics[1][2:])
			if err != nil {
				log.Errorf("[marchalAddr] from failed with hash=%s,err=%s", v.TxHash, err)
				continue
			}
			to, err = marchalAddr(v.Topics[2][2:])
			if err != nil {
				log.Errorf("[marchalAddr] from failed with hash=%s,err=%s", v.TxHash, err)
				continue
			}
			if len(v.Topics) < 4 {
				value = fmt.Sprintf("%d",HexoToInt(v.Data))
				//value, err = marchalInt(fmt.Sprintf("%x", v.Data))
				//if err != nil {
				//	log.Errorf("[marchalInt2] value failed with hash=%s,err=%s", v.TxHash, err)
				//	continue
				//}
				//log.Infof("hash=%s,from=%s,to=%s,value=%s\n", v.TxHash, from, to, value)
			} else {
				value, err = marchalInt(v.Topics[3][2:])
				if err != nil {
					log.Errorf("[marchalInt3] value failed with hash=%s,err=%s", v.TxHash, err)
					continue
				}
				log.Infof("hash=%s,from=%s,to=%s,value=%s", v.TxHash, from, to, value)
			}
		}
		order := &metadata.SQLTokenTransfer{
			TxHash:          v.TxHash,
			LogIndex:        HexoToInt(v.Index),
			Token:           v.Address,
			From:            from,
			To:              to,
			Value:           value,
			Height:          HexoToInt(v.BlockNumber),
			Timestamp:       sync.BlockTime,
		}
		if order.Value!="0"{
			result=append(result, order.Marshal())
		}
	}

	return  result
}
