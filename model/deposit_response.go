package model

type DepositResponse struct {
	WalletID string `json:"wallet_id"`
	Status   bool   `json:"status"`
}
