package view

import (
	"encoding/json"
	"fmt"
)

type Balance struct {
	WalletID string  `json:"wallet_id"`
	Balance  float64 `json:"balance"`
}

// BalanceCodec This codec allows marshalling (encode) and unmarshalling (decode) the Balance to and from the
// Group table.
type BalanceCodec struct{}

// Encode a Balance into []byte
func (jc *BalanceCodec) Encode(value interface{}) ([]byte, error) {
	if _, isBalance := value.(*Balance); !isBalance {
		return nil, fmt.Errorf("codec requires value *Balance, got %T", value)
	}
	return json.Marshal(value)
}

// Decode a Balance from []byte to it's go representation.
func (jc *BalanceCodec) Decode(data []byte) (interface{}, error) {
	var (
		c   Balance
		err error
	)
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling Balance: %v", err)
	}
	return &c, nil
}
