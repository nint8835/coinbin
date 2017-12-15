package main

import (
	"flag"
	"os"
	"fmt"
	"github.com/nint8835/coinbin"
)

func main() {

	fromName := flag.String("from", "none", "The currency to convert from.")
	toName := flag.String("to", "usd", "The currency to convert to.")
	currencyAmount := flag.Float64("amount", 1.0, "The amount of the currency to convert.")

	flag.Parse()

	if *fromName == "none" {
		fmt.Fprint(os.Stderr, "No from currency provided.")
		os.Exit(1)
	}


	coin, err := coinbin.GetCoin(*fromName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving information for coin %s: %s", *fromName, err)
		os.Exit(1)
	}

	if *toName == "usd" {
		coinVal, err := coinbin.GetCoinValue(*fromName, *currencyAmount)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while attempting to retrieve coin value: %s", err)
			os.Exit(1)
		}
		fmt.Printf("%g %s -> %g United States Dollars", *currencyAmount, coin.Name, coinVal.USDValue)
	} else {
		coinTwo, err := coinbin.GetCoin(*toName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error retrieving information for coin %s: %s", *fromName, err)
			os.Exit(1)
		}
		exchange, err := coinbin.GetCoinExchangeValue(*fromName, *toName, *currencyAmount)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error retrieving exchange information: %s", err)
			os.Exit(1)
		}
		fmt.Printf("%g %s -> %g %s", *currencyAmount, coin.Name, exchange.Value, coinTwo.Name)
	}



}
