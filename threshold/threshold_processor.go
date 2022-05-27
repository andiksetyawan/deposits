package threshold

import (
	"context"
	"deposits/config"
	"deposits/model/view"
	"deposits/proto/deposit"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
	"log"
	"time"
)

type thresholdProcessor struct {
	group goka.Group
}

func NewThresholdProcessor(group goka.Group) *thresholdProcessor {
	return &thresholdProcessor{group: group}
}

func (proc *thresholdProcessor) isAboveThreshold(balance *view.Threshold) bool {
	start := time.Now()
	countAmount := float64(0)
	isAbove := false
	for i := len(balance.Histories) - 1; i >= 0; i-- {
		if (start.Unix() - balance.Histories[i].CreatedAt.Unix()) > 120 {
			isAbove = false
			break
		}

		countAmount += balance.Histories[i].Amount
		if countAmount > 10000 {
			isAbove = true
			break
		}
	}
	return isAbove
}

func (proc *thresholdProcessor) callback(ctx goka.Context, msg interface{}) {
	//decode from proto msg
	var depositMsg deposit.Deposit
	proto.Unmarshal(msg.([]byte), &depositMsg)

	var balance *view.Threshold
	if val := ctx.Value(); val != nil {
		balance = val.(*view.Threshold)
	} else {
		balance = new(view.Threshold)
		balance.WalletID = depositMsg.WalletId
	}

	balance.Histories = append(balance.Histories, view.History{
		CreatedAt: time.Now(),
		Amount:    depositMsg.Amount,
	})

	balance.AboveThreshold = proc.isAboveThreshold(balance)
	b, _ := json.Marshal(balance.Histories)
	log.Println("above_threshold :", balance.AboveThreshold)
	log.Println(string(b))

	ctx.SetValue(balance)
	log.Printf("[threshold created] wallet_id: %s amount: %v above_threshold: %v\n", ctx.Key(), depositMsg.Amount, balance.AboveThreshold)
}

func (proc *thresholdProcessor) RunProcessor() {
	g := goka.DefineGroup(proc.group,
		goka.Input(config.Topic, new(codec.Bytes), proc.callback),
		goka.Persist(new(view.ThresholdCodec)),
	)

	p, err := goka.NewProcessor(config.Brokers,
		g,
		goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(config.Tmc)),
		goka.WithConsumerGroupBuilder(goka.DefaultConsumerGroupBuilder),
	)
	if err != nil {
		panic(err)
	}

	err = p.Run(context.Background())
	if err != nil {
		panic(err)
	}
}
