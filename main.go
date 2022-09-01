package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type resp struct {
	Message string `json:"message`
}

type ChargingSession struct {
	Status     ChargingSessionState `json:"state"`
	MeterValue float64              `json:"kWh"`
	Id         string               `json:"_id"`
}

type ChargingSessionState string

const (
	Started ChargingSessionState = "STARTED"
	Aborted ChargingSessionState = "ABORTED"
	Stopped ChargingSessionState = "STOPPED"
	Pending ChargingSessionState = "PENDING"
)

func main() {

	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	url := "http://localhost:8080/ping"

	chargingSession := new(ChargingSession)

	client := resty.New().
		SetRetryWaitTime(2 * time.Second).
		SetRetryCount(5).SetLogger(sugar).SetDebug(true).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			if err != nil {
				sugar.Info(zap.String("Voltnet retries error", err.Error()))
			}
			if r.StatusCode() != 200 {
				return true
			}
			if err = json.Unmarshal(r.Body(), &chargingSession); err != nil {
				sugar.Info(zap.String("Voltnet retries during unmarshal", err.Error()), zap.String("chargingSession", string(chargingSession.Status)))
			}
			if chargingSession.Status == Started {
				return false
			}
			return true
		})

	logger.Info("", zap.String("Voltnet", "making http request to url: "+url))

	_, err := client.R().SetHeader("Accept", "application/json").SetResult(chargingSession).Get(url)

	if err != nil {
		logger.Error("", zap.String("Voltnet", "http request error: "+err.Error()))
		fmt.Println(chargingSession, err)
	}
	fmt.Println(chargingSession)
}
