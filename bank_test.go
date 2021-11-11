package money_test

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/pioz/money"
	"github.com/stretchr/testify/assert"
)

func ExampleNewMoney() {
	m, _ := money.NewMoney(100, "USD")
	fmt.Println(m.Format())
	// Output: $1.00
}

func ExampleBank_NewMoney() {
	bank, _ := money.NewBankFromStaticExchangeRatesTable([]money.Currency{money.USD}, nil)
	m, _ := bank.NewMoney(100, "USD")
	fmt.Println(m.Format())
	// Output: $1.00
}

func ExampleNewMoneyFromAmount() {
	m, _ := money.NewMoneyFromAmount(1.499, "USD")
	fmt.Println(m.Format())
	// Output: $1.50
}

func ExampleBank_NewMoneyFromAmount() {
	bank, _ := money.NewBankFromStaticExchangeRatesTable([]money.Currency{money.USD}, nil)
	m, _ := bank.NewMoneyFromAmount(1.502, "USD")
	fmt.Println(m.Format())
	// Output: $1.50
}

func TestDefaultBankCurrencies(t *testing.T) {
	assert.Equal(t, 182, len(money.DefaultBank.Currencies))
	assert.Equal(t, "EUR", money.DefaultBank.Currencies["EUR"].IsoCode)
	assert.Equal(t, "USD", money.DefaultBank.Currencies["USD"].IsoCode)
}

func TestExchangeDefaultBank(t *testing.T) {
	m, err := money.NewMoney(1234, "EUR")
	assert.Nil(t, err)

	_, err = m.ExchangeTo("USD")
	assert.NotNil(t, err)
	assert.Equal(t, "bank does not support exchange from EUR to USD", err.Error())
}

func TestNewBank(t *testing.T) {
	f := func() (money.ExchangeRatesTable, error) {
		return money.ExchangeRatesTable{
			"EUR": {"EUR": 1.0, "USD": 1.2, "GBP": 1.3},
			"AUD": {"EUR": 2.0, "USD": 2.2, "GBP": 2.3},
		}, nil
	}
	bank, err := money.NewBank([]money.Currency{money.EUR, money.USD}, f, nil)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(bank.Currencies))
	assert.Equal(t, "EUR", bank.Currencies["EUR"].IsoCode)
	assert.Equal(t, "USD", bank.Currencies["USD"].IsoCode)

	m, err := bank.NewMoney(100, "EUR")
	assert.Nil(t, err)
	assert.Equal(t, "€1,00", m.Format())

	_, err = bank.NewMoney(100, "XLN")
	assert.NotNil(t, err)
	assert.Equal(t, "bank does not support XLN currency", err.Error())

	_, err = bank.NewMoneyFromAmount(100, "XLN")
	assert.NotNil(t, err)
	assert.Equal(t, "bank does not support XLN currency", err.Error())
}

func TestUpdateRatesError(t *testing.T) {
	f := func() (money.ExchangeRatesTable, error) {
		return nil, fmt.Errorf("Cannot update exchange rates table")
	}
	bank, err := money.NewBank([]money.Currency{money.EUR, money.USD}, f, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Cannot update exchange rates table", err.Error())
	assert.NotNil(t, bank)
}

func TestGetExchangeRate(t *testing.T) {
	f := func() (money.ExchangeRatesTable, error) {
		return money.ExchangeRatesTable{
			"EUR": {"USD": 1.2, "GBP": 1.3},
			"AUD": {"USD": 2.2, "GBP": 2.3},
		}, nil
	}
	bank, err := money.NewBank([]money.Currency{money.EUR, money.USD}, f, nil)
	assert.Nil(t, err)

	r, err := bank.GetExchangeRate("EUR", "EUR")
	assert.Nil(t, err)
	assert.Equal(t, 1.0, r)

	r, err = bank.GetExchangeRate("EUR", "USD")
	assert.Nil(t, err)
	assert.Equal(t, 1.2, r)

	_, err = bank.GetExchangeRate("EUR", "XLN")
	assert.NotNil(t, err)
	assert.Equal(t, "bank does not support exchange from EUR to XLN", err.Error())

	_, err = bank.GetExchangeRate("XLN", "EUR")
	assert.NotNil(t, err)
	assert.Equal(t, "bank does not support exchange from XLN to EUR", err.Error())
}

func TestRatesCache(t *testing.T) {
	fileCache := money.ExchangeRatesTableFileCache{FilePath: "/tmp/exchange-rates-table-cache"}
	defer os.RemoveAll(fileCache.FilePath)

	c := []money.Currency{money.EUR, money.USD}
	f := func() (money.ExchangeRatesTable, error) {
		// Simutate fetch from Internet
		time.Sleep(10 * time.Millisecond)
		return money.ExchangeRatesTable{
			"EUR": {"USD": rand.Float64()},
			"USD": {"EUR": rand.Float64()},
		}, nil
	}
	bank, _ := money.NewBank(c, f, fileCache)

	rate1, err := bank.GetExchangeRate("EUR", "USD")
	assert.Nil(t, err)
	_, err = os.Stat(fileCache.FilePath)
	assert.False(t, errors.Is(err, os.ErrNotExist))

	cachedRates, err := fileCache.Read()
	assert.Nil(t, err)
	assert.Equal(t, rate1, cachedRates["EUR"]["USD"])

	bank, _ = money.NewBank(c, f, fileCache)
	rate2, _ := bank.GetExchangeRate("EUR", "USD")
	assert.Equal(t, rate1, rate2)

	time.Sleep(20 * time.Millisecond)

	rate3, _ := bank.GetExchangeRate("EUR", "USD")
	assert.NotEqual(t, rate1, rate3)
	cachedRates, err = fileCache.Read()
	assert.Nil(t, err)
	assert.Equal(t, rate3, cachedRates["EUR"]["USD"])
}

func TestIncBank(t *testing.T) {
	counter := 0.0
	var bank, _ = money.NewBank(money.AllCurrencies, func() (money.ExchangeRatesTable, error) {
		counter += 1.0
		table := make(money.ExchangeRatesTable)
		for _, fromCurrency := range money.AllCurrencies {
			table[fromCurrency.IsoCode] = make(money.ExchangeRates)
			for _, toCurrency := range money.AllCurrencies {
				if fromCurrency.IsoCode == toCurrency.IsoCode {
					continue
				}
				table[fromCurrency.IsoCode][toCurrency.IsoCode] = (float64(fromCurrency.SubunitToUnit) / float64(toCurrency.SubunitToUnit)) * counter
			}
		}
		return table, nil
	}, nil)

	m, _ := bank.NewMoney(100, "EUR")

	ex, _ := m.ExchangeTo("USD")
	assert.Equal(t, "$1.00", ex.Format())

	ex, _ = m.ExchangeTo("USD")
	assert.Equal(t, "$1.00", ex.Format())

	err := bank.UpdateExchangeRatesTable()
	assert.Nil(t, err)

	ex, _ = m.ExchangeTo("USD")
	assert.Equal(t, "$2.00", ex.Format())

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := m.ExchangeTo("USD")
			assert.Nil(t, err)
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := bank.UpdateExchangeRatesTable()
		assert.Nil(t, err)
	}()
	wg.Wait()

	ex, _ = m.ExchangeTo("USD")
	assert.Equal(t, "$3.00", ex.Format())
}
