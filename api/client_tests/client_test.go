package client_tests


import (
	. "github.com/Andrew161644/currency_exchange_grpc/api/client"
	. "github.com/Andrew161644/currency_exchange_grpc/api/model"
	"log"
	"testing"
)

var host = "amqp://guest:guest@localhost:5672/"
func TestCanSendTask(t *testing.T)  {
	var res, err = ExchangerRPC(CurrencyExchangeTask{
		UserId:              1,
		Value:               156,
		CurrentCurrencyName: "USD",
		NewCurrencyName:     "RUB",
		Result:              0,
	},host)
	if err!=nil {
		log.Fatal(err)
	}else {
		log.Println(res)
	}
}