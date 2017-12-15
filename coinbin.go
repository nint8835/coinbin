package coinbin

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"errors"
)

type Coin struct {
	Name     string  `json:"name"`
	Ticker   string  `json:"ticker"`
	USDValue float64 `json:"usd"`
	BTCValue float64 `json:"btc"`
	Rank     int     `json:"rank"`
}

func (c Coin) GetValue(amount float64) (CoinValue, error) {
	return GetCoinValue(c.Name, amount)
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

type CoinExchangeRate struct {
	ExchangeRate float64 `json:"exchange_rate"`
}

type CoinExchangeRateResponse struct {
	CoinExchangeRate CoinExchangeRate `json:"coin"`
}

type CoinExchangeValue struct {
	ExchangeRate float64 `json:"exchange_rate"`
	Value        float64 `json:"value"`
	ResultCoin   string `json:"value.coin"`
}

type CoinExchangeValueResponse struct {
	CoinExchangeValue CoinExchangeValue `json:"coin"`
}

func GetCoin(name string) (Coin, error) {
	url := fmt.Sprintf("https://coinbin.org/%s", name)
	response, err := http.Get(url)
	if err != nil {
		return Coin{}, err
	}
	if response.StatusCode == 500 {
		return Coin{}, errors.New("specified coin does not exist")
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
	if response.StatusCode == 500 {
		return CoinValue{}, errors.New("specified coin does not exist")
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

func GetCoinExchangeRate(from string, to string) (CoinExchangeRate, error) {
	url := fmt.Sprintf("https://coinbin.org/%s/to/%s", from, to)
	response, err := http.Get(url)
	if err != nil {
		return CoinExchangeRate{}, err
	}
	if response.StatusCode == 500 {
		return CoinExchangeRate{}, errors.New("specified coin does not exist")
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return CoinExchangeRate{}, err
	}

	var coinExchangeRateResponse CoinExchangeRateResponse

	err = json.Unmarshal(data, &coinExchangeRateResponse)
	if err != nil {
		return CoinExchangeRate{}, err
	}

	return coinExchangeRateResponse.CoinExchangeRate, err
}

func GetCoinExchangeValue(from string, to string, amount float64) (CoinExchangeValue, error) {
	url := fmt.Sprintf("https://coinbin.org/%s/%f/to/%s", from, amount, to)
	response, err := http.Get(url)
	if err != nil {
		return CoinExchangeValue{}, err
	}
	if response.StatusCode == 500 {
		return CoinExchangeValue{}, errors.New("specified coin does not exist")
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return CoinExchangeValue{}, err
	}

	var coinExchangeValueResponse CoinExchangeValueResponse

	err = json.Unmarshal(data, &coinExchangeValueResponse)
	if err != nil {
		return CoinExchangeValue{}, err
	}

	return coinExchangeValueResponse.CoinExchangeValue, err
}