package main

import (
	"deposits/config"
	"deposits/threshold"
)

func main() {
	p := threshold.NewThresholdProcessor(config.GroupThreshold)
	p.RunProcessor()
}
