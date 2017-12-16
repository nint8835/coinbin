package main

import (
	"flag"
	"os"
	"fmt"
	"github.com/nint8835/coinbin"
	"github.com/fatih/color"
)

func main() {

	fromName := flag.String("from", "none", "The currency to convert from.")
	toName := flag.String("to", "usd", "The currency to convert to.")
	currencyAmount := flag.Float64("amount", 1.0, "The amount of the currency to convert.")
	noColour := flag.Bool("no-colour", false, "Disable coloured output.")

	flag.Parse()

	if *noColour {
		color.NoColor = true
	}
	red := color.New(color.FgHiRed)
	blue := color.New(color.FgHiBlue)



	if *fromName == "none" {
		red.Fprint(os.Stderr, "No from currency provided.")
		os.Exit(1)
	}
	
	coin, err := coinbin.GetCoin(*fromName)

	if err != nil {
		red.Fprintf(os.Stderr, "Error retrieving information for coin %s: %s", *fromName, err)
		os.Exit(1)
	}

	if *toName == "usd" {
		coinVal, err := coinbin.GetCoinValue(*fromName, *currencyAmount)
		if err != nil {
			red.Fprintf(os.Stderr, "Error while attempting to retrieve coin value: %s", err)
			os.Exit(1)
		}
		fmt.Printf("%s %s -> %s United States Dollars", blue.Sprintf("%g", *currencyAmount), coin.Name, blue.Sprintf("%g", coinVal.USDValue))
	} else {
		coinTwo, err := coinbin.GetCoin(*toName)
		if err != nil {
			red.Fprintf(os.Stderr, "Error retrieving information for coin %s: %s", *fromName, err)
			os.Exit(1)
		}
		exchange, err := coinbin.GetCoinExchangeValue(*fromName, *toName, *currencyAmount)
		if err != nil {
			red.Fprintf(os.Stderr, "Error retrieving exchange information: %s", err)
			os.Exit(1)
		}
		fmt.Printf("%s %s -> %s %s", blue.Sprintf("%g", *currencyAmount), coin.Name, blue.Sprintf("%g", exchange.Value), coinTwo.Name)
	}

}
