package goanda

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestReduceTradeSize(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Error(err.Error())
	}
	testAPIKey := os.Getenv("OANDA_API_KEY_TEST")
	testAccountID := os.Getenv("OANDA_ACCOUNT_ID_TEST")
	oanda := NewConnection(testAccountID, testAPIKey, false)
	var tradeID string
	t.Run("market order with client extension ID", func(t *testing.T) {
		order := oanda.CreateOrder(OrderPayload{
			OrderBody{
				Units:        10,
				Type:         "MARKET",
				Instrument:   "GBP_USD",
				TimeInForce:  "FOK",
				PositionFill: "DEFAULT",
			},
		})
		tradeID = order.OrderFillTransaction.TradeOpened.TradeID
	})
	t.Run("set client extensions ID for a trade", func(t *testing.T) {
		_ = oanda.SetClientExtensions(tradeID, ClientExtension{
			ClientExtensions: &ClientExtensions{
				Comment: "whateva",
				Tag:     "trade-test",
				ID:      "my_trade_2",
			},
		})
	})
	_ = oanda.ReduceTradeSize("@my_trade_2", CloseTradePayload{
		Units: "10",
	})
}
