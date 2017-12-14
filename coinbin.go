package coinbin

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Coin struct {
	Name     string  `json:"name"`
	Ticker   string  `json:"ticker"`
	USDValue float64 `json:"usd"`
	BTCValue float64 `json:"btc"`
	Rank     int     `json:"rank"`
}

type CoinResponse struct {
	Coin Coin `json:"coin"`
}

type CoinValue struct {
	ExchangeRate float64 `json:"exchange_rate"`
	USDValue     float64 `json:"usd"`
}

type CoinValueResponse struct {
	CoinValue CoinValue `json:"coin"`
}

func GetCoin(name string) (Coin, error) {
	url := fmt.Sprintf("https://coinbin.org/%s", name)
	response, err := http.Get(url)
	if err != nil {
		return Coin{}, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Coin{}, err
	}

	var coinResponse CoinResponse

	err = json.Unmarshal(data, &coinResponse)
	if err != nil {
		return Coin{}, err
	}

	return coinResponse.Coin, err
}

func GetCoinValue(name string, amount float64) (CoinValue, error) {
	url := fmt.Sprintf("https://coinbin.org/%s/%f", name, amount)
	response, err := http.Get(url)
	if err != nil {
		return CoinValue{}, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return CoinValue{}, err
	}

	var coinValueResponse CoinValueResponse

	err = json.Unmarshal(data, &coinValueResponse)
	if err != nil {
		return CoinValue{}, err
	}

	return coinValueResponse.CoinValue, err
}