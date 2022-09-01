package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var count int

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		// time.Sleep(5 * time.Second)
		count++
		if count > 3 {
			c.JSON(http.StatusOK, gin.H{
				"_id":                           "6310211ea9edc30f1eb29f3e",
				"countryCode":                   "SG",
				"entityCode":                    "E8634660",
				"entityName":                    "EVA CPO QA",
				"kWh":                           0,
				"currency":                      "SGD",
				"connectivity":                  "BB3_OCPP",
				"connectivityChargingSessionId": "788",
				"state":                         "STARTED",
				"stopReason":                    "UNKNOWN",
				"startDateTime":                 "2022-09-01T03:03:58.368Z",
				"endDateTime":                   "",
				"authenticationToken":           "",
				"location":                      "",
				"evseId":                        "CPO_A_CHARGER_2",
				"connectorId":                   "1",
				"meterId":                       "",
				"connectorPowerType":            "AC_1_PHASE",
				"paymentMethodCode":             "FREE_OF_CHARGE",
				"paymentTransactionId":          "",
				"paymentRefundTransactionId":    "",
				"paymentState":                  "PRE_AUTH_PENDING",
				"chargingPeriods":               "",
				"tariffs":                       "",
				"totalCost":                     "",
				"updatedAt":                     "2022-09-01T03:04:41.413Z",
				"deletedAt":                     "null",
				"VCID":                          "",
				"createdAt":                     "2022-09-01T03:03:58.368Z",
				"__v":                           0,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"state": "CREATED",
			})
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
