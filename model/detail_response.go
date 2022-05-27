package model

type DetailResponse struct {
	WalletID       string  `json:"wallet_id"`
	Balance        float64 `json:"balance"`
	AboveThreshold bool    `json:"above_threshold"`
}
