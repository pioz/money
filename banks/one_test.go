package banks_test

import (
	"testing"

	"github.com/pioz/money"
	"github.com/pioz/money/banks"
	"github.com/stretchr/testify/assert"
)

func TestOneBank(t *testing.T) {
	oneBank, err := banks.NewOneBank(money.AllCurrencies)
	assert.Nil(t, err)

	m, err := oneBank.NewMoney(1234, "EUR")
	assert.Nil(t, err)

	exchanged, err := m.ExchangeTo("USD")
	assert.Nil(t, err)
	assert.Equal(t, 1234, exchanged.Cents)
	assert.Equal(t, "USD", exchanged.Currency)

	exchanged, err = m.ExchangeTo("JPY")
	assert.Nil(t, err)
	assert.Equal(t, 1234, exchanged.Cents)
	assert.Equal(t, "JPY", exchanged.Currency)

	exchanged, err = m.ExchangeTo("EUR")
	assert.Nil(t, err)
	assert.Equal(t, 1234, exchanged.Cents)
	assert.Equal(t, "EUR", exchanged.Currency)
}
