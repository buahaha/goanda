package goanda

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestOrderClientExtensions(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Error(err.Error())
	}
	testAPIKey := os.Getenv("OANDA_API_KEY_TEST")
	testAccountID := os.Getenv("OANDA_ACCOUNT_ID_TEST")
	oanda := NewConnection(testAccountID, testAPIKey, false)
	tradeID := "myTrade"
	// oanda.CreateOrder(OrderPayload{
	// 	OrderBody{
	// 		Units:        1,
	// 		Instrument:   "EUR_USD",
	// 		Type:         "STOP",
	// 		PositionFill: "DEFAULT",
	// 		TimeInForce:  "GTC",
	// 		Price:        "1.4",
	// 		ClientExtensions: &OrderClientExtensions{
	// 			ID:      tradeID,
	// 			Comment: "Nutek and Oanda on the Go!",
	// 			Tag:     "strategy1",
	// 		},
	// 	},
	// })
	order := oanda.GetOrder("@" + tradeID)
	if order.Order.Price != "1.40000" {
		t.Error("Wrong price")
	}
	if order.Order.ClientExtensions.Tag != "strategy1" {
		t.Error("Wrong tag")
	}
}
