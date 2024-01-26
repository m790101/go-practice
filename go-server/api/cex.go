package api

import (
	"encoding/json"
	"fmt"
	dataTypes "hw/go-server/model"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*dataTypes.Rate, error) {
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))

	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		bodyByte, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var cryptoRate dataTypes.Rate
		// json := string(bodyByte)
		err = json.Unmarshal(bodyByte, &cryptoRate)
		// json.Marshaler(bodyByte)
		if err != nil {
			return nil, err
		}
		// fmt.Println(json)

	} else {
		return nil, fmt.Errorf("Currency %s not found", currency)
	}
	rate := dataTypes.Rate{Currency: upCurrency, Price: 20.0}
	return &rate, nil

}
