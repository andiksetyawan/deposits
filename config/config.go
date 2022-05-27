package config

import "github.com/lovoo/goka"

var (
	Brokers                    = []string{"127.0.0.1:9092"}
	Topic          goka.Stream = "deposit"
	GroupBalance   goka.Group  = "balance"
	GroupThreshold goka.Group  = "threshold"
	Tmc                        = goka.NewTopicManagerConfig()
)

func init() {
	Tmc.Table.Replication = 1
	Tmc.Stream.Replication = 1
}
