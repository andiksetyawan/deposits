package view

import (
	"encoding/json"
	"fmt"
	"time"
)

type Threshold struct {
	WalletID       string    `json:"wallet_id"`
	AboveThreshold bool      `json:"above_threshold"`
	Histories      []History `json:"histories"`
}

type History struct {
	CreatedAt time.Time `json:"created_at"`
	Amount    float64   `json:"amount,omitempty"`
}

// ThresholdCodec This codec allows marshalling (encode) and unmarshalling (decode) the Balance to and from the
// Group table.
type ThresholdCodec struct{}

// Encode a Threshold into []byte
func (jc *ThresholdCodec) Encode(value interface{}) ([]byte, error) {
	if _, isThreshold := value.(*Threshold); !isThreshold {
		return nil, fmt.Errorf("codec requires value *Threshold, got %T", value)
	}
	return json.Marshal(value)
}

// Decode a Threshold from []byte
func (jc *ThresholdCodec) Decode(data []byte) (interface{}, error) {
	var (
		c   Threshold
		err error
	)
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling Threshold: %v", err) //TODO
	}
	return &c, nil
}
