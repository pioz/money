package banks_test

import (
	"os"
	"testing"

	"github.com/pioz/money"
	"github.com/pioz/money/banks"
	"github.com/stretchr/testify/assert"
)

func TestFreecurrencyBank(t *testing.T) {
	apiKey := os.Getenv("FREECURRENCY_API_KEY")
	if apiKey == "" {
		t.Skip("Missing https://freecurrencyapi.net/ api key: run test with `FREECURRENCY_API_KEY=your-api-key go test`")
	}
	freeBank, err := banks.NewFreecurrencyBank([]money.Currency{money.EUR, money.USD, money.ANG}, apiKey, nil)
	assert.Nil(t, err)
	m, err := freeBank.NewMoney(100, "EUR")
	assert.Nil(t, err)

	ex, err := m.ExchangeTo("USD")
	assert.Nil(t, err)
	t.Log(ex.Format())

	rate, err := freeBank.GetExchangeRate("EUR", "USD")
	assert.Nil(t, err)
	assert.True(t, rate != 0)

	rate, err = freeBank.GetExchangeRate("USD", "EUR")
	assert.Nil(t, err)
	assert.True(t, rate != 0)

	_, err = freeBank.GetExchangeRate("EUR", "ANG")
	assert.NotNil(t, err)
	assert.Equal(t, "bank does not support exchange from EUR to ANG", err.Error())

	_, err = m.ExchangeTo("ANG")
	assert.NotNil(t, err)
	assert.Equal(t, "bank does not support exchange from EUR to ANG", err.Error())
}

func TestFreecurrencyapiBankInvalidApiKey(t *testing.T) {
	_, err := banks.NewFreecurrencyBank(money.AllCurrencies, "", nil)
	assert.NotNil(t, err)
	assert.Equal(t, "error to fetch 'AED' exchange rates: HTTP status 429", err.Error())
}
