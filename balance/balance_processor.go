package balance

import (
	"context"
	"deposits/config"
	"deposits/model/view"
	"deposits/proto/deposit"
	"github.com/golang/protobuf/proto"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
	"log"
)

type balanceProcessor struct {
	group goka.Group
}

func NewBalanceProcessor(group goka.Group) *balanceProcessor {
	return &balanceProcessor{group: group}
}

func (proc *balanceProcessor) callback(ctx goka.Context, msg interface{}) {
	//decode from proto msg
	var depositMsg deposit.Deposit
	proto.Unmarshal(msg.([]byte), &depositMsg)

	var balance *view.Balance
	if val := ctx.Value(); val != nil {
		balance = val.(*view.Balance)
	} else {
		balance = new(view.Balance)
		balance.WalletID = depositMsg.WalletId
	}

	balance.Balance += depositMsg.Amount
	ctx.SetValue(balance)
	log.Printf("[balance created] wallet_id: %s amount: %v total_balance: %v\n", ctx.Key(), depositMsg.Amount, balance.Balance)
}

func (proc *balanceProcessor) RunProcessor() {
	g := goka.DefineGroup(proc.group,
		goka.Input(config.Topic, new(codec.Bytes), proc.callback),
		goka.Persist(new(view.BalanceCodec)),
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
