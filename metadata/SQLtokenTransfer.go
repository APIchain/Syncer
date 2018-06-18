package metadata

import (
	"encoding/json"
)

type SQLTokenTransfer struct {
	TxHash          string
	Timestamp       int64
	Height          int64
	Token           string
	From            string
	To              string
	Value           string
}

func (this *SQLTokenTransfer) Marshal() ([]byte, error) {
	str, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return str, nil
}

//func (this *SQLTransfer) Marshal()string {
//	return fmt.Sprintf("INSERT INTO ethTransfer(TxHash,Timestamp,Height,Sender,SendTo,Value) " +
//		"VALUES ('%s','%s',%d,'%s','%s','%s');",this.TxHash,int2time(this.Timestamp),this.Height,this.From,this.To,this.Value)
//}

func UnMarshalTokenTransfer(str []byte) (*SQLTokenTransfer, error) {
	info := new(SQLTokenTransfer)
	if err := json.Unmarshal(str, info); err != nil {
		return nil, err
	}
	return info, nil
}
