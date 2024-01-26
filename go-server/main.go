package main

import "hw/go-server/api"

func main() {

	rate, err := api.GetRate("BTC")
	if err != nil {
		panic(err)
	}
	println(rate.Currency, rate.Price)
}
