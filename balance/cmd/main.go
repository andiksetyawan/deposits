package main

import (
	"deposits/balance"
	"deposits/config"
)

func main() {
	p := balance.NewBalanceProcessor(config.GroupBalance)
	p.RunProcessor()
}
