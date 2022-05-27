package main

import (
	"context"
	"deposits/config"
	"deposits/model/view"
	"deposits/web"
	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
	"log"
	"net/http"
)

func main() {
	conf := goka.DefaultConfig()
	conf.Consumer.Offsets.Initial = sarama.OffsetOldest
	goka.ReplaceGlobalConfig(conf)

	tm, err := goka.NewTopicManager(config.Brokers, goka.DefaultConfig(), config.Tmc)
	if err != nil {
		log.Fatalf("Error creating topic manager: %v", err)
	}
	err = tm.EnsureStreamExists(string(config.Topic), 8)
	if err != nil {
		log.Fatalf("Error creating kafka topic %s: %v", config.Topic, err)
	}

	err = tm.EnsureTableExists(string(config.GroupBalance+"-table"), 8)
	if err != nil {
		log.Fatalf("Error creating kafka table %s: %v", config.Topic, err)
	}
	err = tm.EnsureTableExists(string(config.GroupThreshold+"-table"), 8)
	if err != nil {
		log.Fatalf("Error creating kafka table %s: %v", config.Topic, err)
	}

	viewBalance, err := goka.NewView(config.Brokers, goka.GroupTable(config.GroupBalance), new(view.BalanceCodec))
	if err != nil {
		panic(err)
	}
	go viewBalance.Run(context.Background()) //TODO handle err

	viewThreshold, err := goka.NewView(config.Brokers, goka.GroupTable(config.GroupThreshold), new(view.ThresholdCodec))
	if err != nil {
		panic(err)
	}
	go viewThreshold.Run(context.Background()) //TODO handle err

	emitter, err := goka.NewEmitter(config.Brokers, config.Topic, new(codec.Bytes))
	if err != nil {
		panic(err)
	}
	defer emitter.Finish()

	router := mux.NewRouter()
	router.HandleFunc("/deposit", web.DepositC(emitter)).Methods("POST")
	router.HandleFunc("/detail", web.DetailC(viewBalance, viewThreshold)).Methods("GET")

	log.Printf("Listen port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
