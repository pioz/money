package money

import (
	"fmt"
	"log"
	"math"
)

// FetchExchangeRatesTableFunc is the signature of the function to fetch an
// exchange rates table. It take no parameters and returns the exchange rates
// table or an error.
type FetchExchangeRatesTableFunc func() (ExchangeRatesTable, error)

// A Bank makes it possible to create money. It define in which currencies money
// can be created and the exchange rates table to exchange monetary value in a
// currency to another.
type Bank struct {
	// Map by currency ISO code of all currencies supported by the bank
	Currencies              map[string]Currency
	ExchangeRatesTable      ExchangeRatesTable
	exchangeRatesTableCache ExchangeRatesTableCache
	fetchExchangeRatesTable FetchExchangeRatesTableFunc
}

// The DefaultBank supports all currencies (money.AllCurrencies), but it is
// unable to exchange currencies (it does not have an exchange rates table). It
// is helpful to work with currencies if there is no need to exchange them.
var DefaultBank, _ = NewBank(AllCurrencies, nil, nil)

// NewBank creates a new bank that supports the currencies in the slice
// currencies and uses fetch to fetch the exchange rates table. If cache is not
// nil, it is used to set the exchange rates table immediately, while a go
// routine in background updates the table using fetch. Returns an error if
// fetch returns error.
func NewBank(currencies []Currency, fetch FetchExchangeRatesTableFunc, cache ExchangeRatesTableCache) (*Bank, error) {
	bank := &Bank{
		Currencies:              make(map[string]Currency),
		ExchangeRatesTable:      make(ExchangeRatesTable),
		exchangeRatesTableCache: cache,
		fetchExchangeRatesTable: fetch,
	}
	for _, fromCurrency := range currencies {
		bank.Currencies[fromCurrency.IsoCode] = fromCurrency
		bank.ExchangeRatesTable[fromCurrency.IsoCode] = make(ExchangeRates)
	}

	return bank, bank.UpdateExchangeRatesTable()
}

// NewBankFromStaticExchangeRatesTable is a conveniently function to create a
// new bank that use a static exchange rates table.
func NewBankFromStaticExchangeRatesTable(currencies []Currency, table ExchangeRatesTable) (*Bank, error) {
	return NewBank(currencies, func() (ExchangeRatesTable, error) {
		return table, nil
	}, nil)
}

// GetExchangeRate returns the exchange rate to convert the currency with ISO
// code fromCurrencyIsoCode to the currency with ISO code toCurrencyIsoCode.
// Returns an error if the bank does not support one of the two currencies or if
// it does not support the exchange between the two currencies, meaning that the
// exchange rates table does not have the exchange rate.
func (bank *Bank) GetExchangeRate(fromCurrencyIsoCode, toCurrencyIsoCode string) (float64, error) {
	if fromCurrencyIsoCode == toCurrencyIsoCode {
		return 1.0, nil
	}
	exchangeRates := bank.ExchangeRatesTable[fromCurrencyIsoCode]
	rate := exchangeRates[toCurrencyIsoCode]
	if rate == 0.0 {
		return 0.0, fmt.Errorf("bank does not support exchange from %s to %s", fromCurrencyIsoCode, toCurrencyIsoCode)
	}
	return rate, nil
}

// UpdateExchangeRatesTable updates the bank exchange rates table by calling the
// fetch function. If fetch is nil, it has no effect. Returns an error if fetch
// returns error.
func (bank *Bank) UpdateExchangeRatesTable() error {
	if bank.fetchExchangeRatesTable == nil {
		return nil
	}

	if bank.exchangeRatesTableCache == nil {
		return bank.blockingUpdateExchangeRatesTable()
	}

	table, err := bank.exchangeRatesTableCache.Read()
	if err != nil {
		return bank.blockingUpdateExchangeRatesTable()
	}
	bank.setExchangeRatesTable(table)
	go func() {
		err := bank.blockingUpdateExchangeRatesTable()
		if err != nil {
			log.Println(err)
		}
	}()
	return nil
}

// NewMoney creates a new money of value given in the fractional unit of the
// given currency. For example, in the US dollar currency the fractional unit is
// cents, and there are 100 cents in one US dollar. So given the Money
// representation of one US dollar, the fractional interpretation is 100.
// Returns error if the currency is not supported by the bank.
func (bank *Bank) NewMoney(cents int, currencyIsoCode string) (*Money, error) {
	_, err := bank.getCurrency(currencyIsoCode)
	if err != nil {
		return nil, err
	}
	return &Money{Cents: cents, Currency: currencyIsoCode, bank: bank}, nil
}

// NewMoney creates a new money of value given in the unit of the given
// currency. Returns error if the currency is not supported by the bank.
func (bank *Bank) NewMoneyFromAmount(amount float64, currencyIsoCode string) (*Money, error) {
	currency, err := bank.getCurrency(currencyIsoCode)
	if err != nil {
		return nil, err
	}
	cents := int(math.Round(amount * float64(currency.SubunitToUnit)))
	return &Money{Cents: cents, Currency: currencyIsoCode, bank: bank}, nil
}

// NewMoney creates a new money of value given in the fractional unit of the
// given currency using the default bank. Returns error if the currency is not
// supported by the bank.
func NewMoney(cents int, currencyIsoCode string) (*Money, error) {
	return DefaultBank.NewMoney(cents, currencyIsoCode)
}

// NewMoney creates a new money of value given in the unit of the given currency
// using the default bank. Returns error if the currency is not supported by the
// bank.
func NewMoneyFromAmount(amount float64, currencyIsoCode string) (*Money, error) {
	return DefaultBank.NewMoneyFromAmount(amount, currencyIsoCode)
}

// Private functions

func (bank *Bank) getCurrency(currencyIsoCode string) (Currency, error) {
	currency, found := bank.Currencies[currencyIsoCode]
	if !found {
		return currency, fmt.Errorf("bank does not support %s currency", currencyIsoCode)
	}
	return currency, nil
}

func (bank *Bank) blockingUpdateExchangeRatesTable() error {
	table, err := bank.fetchExchangeRatesTable()
	if err != nil {
		return err
	}
	bank.setExchangeRatesTable(table)
	return nil
}

func (bank *Bank) setExchangeRatesTable(table ExchangeRatesTable) {
	for fromCurrencyIsoCode, fromRates := range table {
		fromCurrency, found := bank.Currencies[fromCurrencyIsoCode]
		if !found {
			continue
		}
		for toCurrencyIsoCode, rate := range fromRates {
			if fromCurrencyIsoCode == toCurrencyIsoCode {
				continue
			}
			toCurrency, found := bank.Currencies[toCurrencyIsoCode]
			if !found {
				continue
			}
			bank.ExchangeRatesTable[fromCurrency.IsoCode][toCurrency.IsoCode] = rate
		}
	}
	if bank.exchangeRatesTableCache != nil {
		err := bank.exchangeRatesTableCache.Write(table)
		if err != nil {
			log.Println(err)
		}
	}
}
