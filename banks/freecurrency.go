package banks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pioz/money"
)

const endpoint = "https://freecurrencyapi.net/api/v2/latest"

type response struct {
	Data money.ExchangeRates `json:"data"`
}

func NewFreecurrencyBank(currencies []money.Currency, apiKey string, cache money.ExchangeRatesTableCache) (*money.Bank, error) {
	return money.NewBank(currencies, func() (money.ExchangeRatesTable, error) {
		table := make(money.ExchangeRatesTable)
		for _, currency := range currencies {
			toRates, err := getExchangeRatesTable(apiKey, currency.IsoCode)
			if err != nil {
				return nil, err
			}
			table[currency.IsoCode] = toRates
		}
		return table, nil
	}, cache)
}

func getExchangeRatesTable(apiKey, baseCurrency string) (money.ExchangeRates, error) {
	url := fmt.Sprintf("%s?apikey=%s&base_currency=%s", endpoint, apiKey, baseCurrency)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	contentType := resp.Header.Get("content-type")
	if !strings.Contains(contentType, "application/json") {
		return nil, nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := response{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error to fetch '%s' exchange rates: HTTP status %d", baseCurrency, resp.StatusCode)
	}
	return r.Data, nil
}
