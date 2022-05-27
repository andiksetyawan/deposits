package model

type DepositRequest struct {
	WalletID string  `json:"wallet_id"`
	Amount   float64 `json:"amount"`
}
