package metadata

import (
	"encoding/json"
	"fmt"
)

type SQLTokenTransfer struct {
	TxHash          string
	LogIndex        int64
	Timestamp       int64
	Height          int64
	Token           string
	From            string
	To              string
	Value           string
}

func (this *SQLTokenTransfer) Marshal()string {
	return fmt.Sprintf("INSERT INTO tokenTransfer(TxHash,LogIndex,Timestamp,Height,Token,Sender,SenderTo,Value) VALUES" +
	"('%s',%d,'%s',%d,'%s','%s','%s','%s');", this.TxHash,this.LogIndex,int2time(this.Timestamp),this.Height,this.Token,this.From,this.To,this.Value)
}

func UnMarshalTokenTransfer(str []byte) (*SQLTokenTransfer, error) {
	info := new(SQLTokenTransfer)
	if err := json.Unmarshal(str, info); err != nil {
		return nil, err
	}
	return info, nil
}
