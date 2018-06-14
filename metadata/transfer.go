package metadata

import (
	"encoding/json"
	"fmt"
	"time"
)

type TxTransfer struct {
	TxHash          string   `json:"TxHash" `
	Timestamp       int64    `json:"Timestamp" `
	Height          int64    `json:"Height" `
	From            string   `json:"From"     gencodec:"required"`
	To              string   `json:"To"       gencodec:"required"`
	Value           string   `json:"Value"    gencodec:"required"`
}

func int2time(timxe int64) string{
	tm := time.Unix(timxe, 0)

	return tm.Format("2006-01-02 15:04:05")
}

func (this *TxTransfer) Marshal()string {
	return fmt.Sprintf("INSERT INTO ethTransfer(TxHash,Timestamp,Height,Sender,SendTo,Value) " +
		"VALUES ('%s','%s',%d,'%s','%s','%s');",this.TxHash,int2time(this.Timestamp),this.Height,this.From,this.To,this.Value)
}

func UnMarshalTxTransfer(str []byte) (*TxTransfer, error) {
	info := new(TxTransfer)
	if err := json.Unmarshal(str, info); err != nil {
		return nil, err
	}
	return info, nil
}
