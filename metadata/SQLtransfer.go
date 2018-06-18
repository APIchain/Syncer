package metadata

import (
	"encoding/json"
	"fmt"
	"time"
)

type SQLTransfer struct {
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

func (this *SQLTransfer) Marshal()string {
	return fmt.Sprintf("INSERT INTO etherTransfer(TxHash,Timestamp,Height,Sender,SendTo,Value) " +
		"VALUES ('%s','%s',%d,'%s','%s','%s');",this.TxHash,int2time(this.Timestamp),this.Height,this.From,this.To,this.Value)
}

func UnMarshalTxTransfer(str []byte) (*SQLTransfer, error) {
	info := new(SQLTransfer)
	if err := json.Unmarshal(str, info); err != nil {
		return nil, err
	}
	return info, nil
}
