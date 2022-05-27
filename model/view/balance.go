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

// Encode a user into []byte //TODO
func (jc *BalanceCodec) Encode(value interface{}) ([]byte, error) {
	if _, isBalance := value.(*Balance); !isBalance {
		return nil, fmt.Errorf("codec requires value *Deposit, got %T", value) //TODO
	}
	return json.Marshal(value)
}

// Decode a user from []byte to it's go representation. //TODO
func (jc *BalanceCodec) Decode(data []byte) (interface{}, error) {
	var (
		c   Balance
		err error
	)
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling Wallet: %v", err) //TODO
	}
	return &c, nil
}
