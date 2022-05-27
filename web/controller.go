package web

import (
	"deposits/model"
	"deposits/model/view"
	"deposits/proto/deposit"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/lovoo/goka"
	"net/http"
)

func DetailC(viewBalance *goka.View, viewThreshold *goka.View) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		walletID := r.URL.Query().Get("wallet_id")

		valueBalance, err := viewBalance.Get(walletID)
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err)
			return
		}

		valueThreshold, err := viewThreshold.Get(walletID)
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err)
			return
		}

		balance := valueBalance.(*view.Balance)
		threshold := valueThreshold.(*view.Threshold)

		response := model.DetailResponse{
			WalletID:       balance.WalletID,
			Balance:        balance.Balance,
			AboveThreshold: threshold.AboveThreshold,
		}

		WriteSuccessResponse(w, 200, response, nil)
	}
}

func DepositC(emitter *goka.Emitter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload deposit.Deposit

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&payload)
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err)
			return
		}

		// convert payload to proto
		b, err := proto.Marshal(&payload)
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err)
			return
		}

		err = emitter.EmitSync(payload.WalletId, b)
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err)
			return
		}

		WriteSuccessResponse(w, 200, model.DepositResponse{WalletID: payload.WalletId, Status: true}, nil)
	}
}
