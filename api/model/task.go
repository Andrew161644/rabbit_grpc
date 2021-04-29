package model

type CurrencyExchangeTask struct {
	UserId              int     `json:"user_id"`
	Value               float64 `json:"value"`
	CurrentCurrencyName string  `json:"currentCurrencyName"`
	NewCurrencyName     string  `json:"newCurrencyName"`
	Result              float64 `json:"result"`
}