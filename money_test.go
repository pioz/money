package money_test

import (
	"fmt"
	"testing"

	"github.com/pioz/money"
	"github.com/stretchr/testify/assert"
)

func TestExchange(t *testing.T) {
	bank, err := money.NewBankFromStaticExchangeRatesTable([]money.Currency{money.EUR, money.USD, money.JPY}, money.ExchangeRatesTable{
		"EUR": {"EUR": 1.1, "USD": 1.2, "JPY": 103.3},
		"USD": {"EUR": 1.4, "JPY": 103.3},
	})
	assert.Nil(t, err)

	m, err := bank.NewMoney(100, "EUR")
	assert.Nil(t, err)
	assert.Equal(t, 100, m.Cents)
	assert.Equal(t, "EUR", m.Currency)

	exchanged, err := m.ExchangeTo("EUR")
	assert.Nil(t, err)
	assert.Equal(t, 100, exchanged.Cents)
	assert.Equal(t, "EUR", exchanged.Currency)

	exchanged, err = m.ExchangeTo("USD")
	assert.Nil(t, err)
	assert.Equal(t, 120, exchanged.Cents)
	assert.Equal(t, "USD", exchanged.Currency)

	_, err = exchanged.ExchangeTo("XLN")
	assert.NotNil(t, err)
	assert.Equal(t, "bank does not support exchange from USD to XLN", err.Error())

	exchanged, err = exchanged.ExchangeTo("JPY")
	assert.Nil(t, err)
	assert.Equal(t, 124, exchanged.Cents)
	assert.Equal(t, "JPY", exchanged.Currency)

	_, err = exchanged.ExchangeTo("EUR")
	assert.NotNil(t, err)
	assert.Equal(t, "bank does not support exchange from JPY to EUR", err.Error())
}

func ExampleMoney_ExchangeTo() {
	bank, _ := money.NewBankFromStaticExchangeRatesTable([]money.Currency{money.EUR, money.USD}, money.ExchangeRatesTable{
		"EUR": {"USD": 1.154321},
		"USD": {"EUR": 0.86702},
	})
	usd, _ := bank.NewMoney(100, "USD")
	eur, _ := usd.ExchangeTo("EUR")
	fmt.Println(eur.Format())
	// Output: €0,87
}

func TestExchangeToWithRate(t *testing.T) {
	m, err := money.NewMoney(100, "EUR")
	assert.Nil(t, err)

	_, err = m.ExchangeToWithRate("XLN", 1.2)
	assert.NotNil(t, err)
	assert.Equal(t, "bank does not support XLN currency", err.Error())

	ex, err := m.ExchangeToWithRate("EUR", 1.2)
	assert.Nil(t, err)
	assert.Equal(t, "€1,00", ex.Format())

	ex, err = m.ExchangeToWithRate("USD", 1.2)
	assert.Nil(t, err)
	assert.Equal(t, "$1.20", ex.Format())
}

func TestEquals(t *testing.T) {
	bank, err := money.NewBankFromStaticExchangeRatesTable([]money.Currency{money.EUR, money.USD}, money.ExchangeRatesTable{"USD": {"EUR": 0.8}})
	assert.Nil(t, err)

	m1, err := bank.NewMoney(100, "EUR")
	assert.Nil(t, err)
	m2, err := bank.NewMoney(100, "EUR")
	assert.Nil(t, err)
	eq, err := m1.Equals(m2)
	assert.Nil(t, err)
	assert.True(t, eq)

	m2, err = bank.NewMoney(90, "EUR")
	assert.Nil(t, err)
	eq, err = m1.Equals(m2)
	assert.Nil(t, err)
	assert.False(t, eq)

	m2, err = bank.NewMoney(125, "USD")
	assert.Nil(t, err)
	eq, err = m1.Equals(m2)
	assert.Nil(t, err)
	assert.True(t, eq)

	m2, err = bank.NewMoney(124, "USD")
	assert.Nil(t, err)
	eq, err = m1.Equals(m2)
	assert.Nil(t, err)
	assert.False(t, eq)

	m2, err = money.NewMoney(100, "USD")
	assert.Nil(t, err)
	_, err = m1.Equals(m2)
	assert.NotNil(t, err)
	assert.Equal(t, "currencies have different banks: operation between currencies can be done only between currencies of the same bank", err.Error())
}

func TestGreaterThan(t *testing.T) {
	bank, err := money.NewBankFromStaticExchangeRatesTable([]money.Currency{money.EUR, money.USD}, money.ExchangeRatesTable{"USD": {"EUR": 0.8}})
	assert.Nil(t, err)

	m1, err := bank.NewMoney(100, "EUR")
	assert.Nil(t, err)
	m2, err := bank.NewMoney(90, "EUR")
	assert.Nil(t, err)
	eq, err := m1.GreaterThan(m2)
	assert.Nil(t, err)
	assert.True(t, eq)

	m2, err = bank.NewMoney(100, "EUR")
	assert.Nil(t, err)
	eq, err = m1.GreaterThan(m2)
	assert.Nil(t, err)
	assert.False(t, eq)

	m2, err = bank.NewMoney(110, "EUR")
	assert.Nil(t, err)
	eq, err = m1.GreaterThan(m2)
	assert.Nil(t, err)
	assert.False(t, eq)

	m2, err = bank.NewMoney(124, "USD")
	assert.Nil(t, err)
	eq, err = m1.GreaterThan(m2)
	assert.Nil(t, err)
	assert.True(t, eq)

	m2, err = bank.NewMoney(125, "USD")
	assert.Nil(t, err)
	eq, err = m1.GreaterThan(m2)
	assert.Nil(t, err)
	assert.False(t, eq)

	m2, err = money.NewMoney(100, "USD")
	assert.Nil(t, err)
	_, err = m1.GreaterThan(m2)
	assert.NotNil(t, err)
	assert.Equal(t, "currencies have different banks: operation between currencies can be done only between currencies of the same bank", err.Error())
}

func TestGreaterThanOrEqual(t *testing.T) {
	bank, err := money.NewBankFromStaticExchangeRatesTable([]money.Currency{money.EUR, money.USD}, money.ExchangeRatesTable{"USD": {"EUR": 0.8}})
	assert.Nil(t, err)

	m1, err := bank.NewMoney(100, "EUR")
	assert.Nil(t, err)
	m2, err := bank.NewMoney(90, "EUR")
	assert.Nil(t, err)
	eq, err := m1.GreaterThanOrEqual(m2)
	assert.Nil(t, err)
	assert.True(t, eq)

	m2, err = bank.NewMoney(100, "EUR")
	assert.Nil(t, err)
	eq, err = m1.GreaterThanOrEqual(m2)
	assert.Nil(t, err)
	assert.True(t, eq)

	m2, err = bank.NewMoney(110, "EUR")
	assert.Nil(t, err)
	eq, err = m1.GreaterThanOrEqual(m2)
	assert.Nil(t, err)
	assert.False(t, eq)

	m2, err = bank.NewMoney(125, "USD")
	assert.Nil(t, err)
	eq, err = m1.GreaterThanOrEqual(m2)
	assert.Nil(t, err)
	assert.True(t, eq)

	m2, err = bank.NewMoney(126, "USD")
	assert.Nil(t, err)
	eq, err = m1.GreaterThanOrEqual(m2)
	assert.Nil(t, err)
	assert.False(t, eq)

	m2, err = money.NewMoney(100, "USD")
	assert.Nil(t, err)
	_, err = m1.GreaterThanOrEqual(m2)
	assert.NotNil(t, err)
	assert.Equal(t, "currencies have different banks: operation between currencies can be done only between currencies of the same bank", err.Error())
}

func TestLessThan(t *testing.T) {
	bank, err := money.NewBankFromStaticExchangeRatesTable([]money.Currency{money.EUR, money.USD}, money.ExchangeRatesTable{"USD": {"EUR": 0.8}})
	assert.Nil(t, err)

	m1, err := bank.NewMoney(100, "EUR")
	assert.Nil(t, err)
	m2, err := bank.NewMoney(110, "EUR")
	assert.Nil(t, err)
	eq, err := m1.LessThan(m2)
	assert.Nil(t, err)
	assert.True(t, eq)

	m2, err = bank.NewMoney(100, "EUR")
	assert.Nil(t, err)
	eq, err = m1.LessThan(m2)
	assert.Nil(t, err)
	assert.False(t, eq)

	m2, err = bank.NewMoney(90, "EUR")
	assert.Nil(t, err)
	eq, err = m1.LessThan(m2)
	assert.Nil(t, err)
	assert.False(t, eq)

	m2, err = bank.NewMoney(126, "USD")
	assert.Nil(t, err)
	eq, err = m1.LessThan(m2)
	assert.Nil(t, err)
	assert.True(t, eq)

	m2, err = bank.NewMoney(125, "USD")
	assert.Nil(t, err)
	eq, err = m1.LessThan(m2)
	assert.Nil(t, err)
	assert.False(t, eq)

	m2, err = money.NewMoney(100, "USD")
	assert.Nil(t, err)
	_, err = m1.LessThan(m2)
	assert.NotNil(t, err)
	assert.Equal(t, "currencies have different banks: operation between currencies can be done only between currencies of the same bank", err.Error())
}

func TestLessThanOrEqual(t *testing.T) {
	bank, err := money.NewBankFromStaticExchangeRatesTable([]money.Currency{money.EUR, money.USD}, money.ExchangeRatesTable{"USD": {"EUR": 0.8}})
	assert.Nil(t, err)

	m1, err := bank.NewMoney(100, "EUR")
	assert.Nil(t, err)
	m2, err := bank.NewMoney(110, "EUR")
	assert.Nil(t, err)
	eq, err := m1.LessThanOrEqual(m2)
	assert.Nil(t, err)
	assert.True(t, eq)

	m2, err = bank.NewMoney(100, "EUR")
	assert.Nil(t, err)
	eq, err = m1.LessThanOrEqual(m2)
	assert.Nil(t, err)
	assert.True(t, eq)

	m2, err = bank.NewMoney(90, "EUR")
	assert.Nil(t, err)
	eq, err = m1.LessThanOrEqual(m2)
	assert.Nil(t, err)
	assert.False(t, eq)

	m2, err = bank.NewMoney(125, "USD")
	assert.Nil(t, err)
	eq, err = m1.LessThanOrEqual(m2)
	assert.Nil(t, err)
	assert.True(t, eq)

	m2, err = bank.NewMoney(124, "USD")
	assert.Nil(t, err)
	eq, err = m1.LessThanOrEqual(m2)
	assert.Nil(t, err)
	assert.False(t, eq)

	m2, err = money.NewMoney(100, "USD")
	assert.Nil(t, err)
	_, err = m1.LessThanOrEqual(m2)
	assert.NotNil(t, err)
	assert.Equal(t, "currencies have different banks: operation between currencies can be done only between currencies of the same bank", err.Error())
}

func TestIsZero(t *testing.T) {
	m, err := money.NewMoney(1, "EUR")
	assert.Nil(t, err)
	assert.False(t, m.IsZero())

	m, err = money.NewMoney(-1, "EUR")
	assert.Nil(t, err)
	assert.False(t, m.IsZero())

	m, err = money.NewMoney(0, "EUR")
	assert.Nil(t, err)
	assert.True(t, m.IsZero())
}

func TestIsNegative(t *testing.T) {
	m, err := money.NewMoney(1, "EUR")
	assert.Nil(t, err)
	assert.False(t, m.IsNegative())

	m, err = money.NewMoney(-1, "EUR")
	assert.Nil(t, err)
	assert.True(t, m.IsNegative())

	m, err = money.NewMoney(0, "EUR")
	assert.Nil(t, err)
	assert.False(t, m.IsNegative())
}

func TestIsPositive(t *testing.T) {
	m, err := money.NewMoney(1, "EUR")
	assert.Nil(t, err)
	assert.True(t, m.IsPositive())

	m, err = money.NewMoney(-1, "EUR")
	assert.Nil(t, err)
	assert.False(t, m.IsPositive())

	m, err = money.NewMoney(0, "EUR")
	assert.Nil(t, err)
	assert.False(t, m.IsPositive())
}

func TestAbsolute(t *testing.T) {
	m1, err := money.NewMoney(-100, "EUR")
	assert.Nil(t, err)
	m2 := m1.Absolute()
	assert.Equal(t, "€1,00", m2.Format())
}

func TestAdd(t *testing.T) {
	bank, err := money.NewBankFromStaticExchangeRatesTable([]money.Currency{money.EUR, money.USD}, money.ExchangeRatesTable{"USD": {"EUR": 0.8}})
	assert.Nil(t, err)

	m1, err := bank.NewMoney(100, "EUR")
	assert.Nil(t, err)
	m2, err := bank.NewMoney(100, "EUR")
	assert.Nil(t, err)
	m3, err := m1.Add(m2)
	assert.Nil(t, err)
	assert.Equal(t, "€2,00", m3.Format())
	assert.NotEqual(t, m1, m3)
	assert.NotEqual(t, m2, m3)

	m2, err = bank.NewMoney(100, "USD")
	assert.Nil(t, err)
	m3, err = m1.Add(m2)
	assert.Nil(t, err)
	assert.Equal(t, "€1,80", m3.Format())
	assert.NotEqual(t, m1, m3)
	assert.NotEqual(t, m2, m3)

	m2, err = money.NewMoney(100, "USD")
	assert.Nil(t, err)
	_, err = m1.Add(m2)
	assert.NotNil(t, err)
	assert.Equal(t, "currencies have different banks: operation between currencies can be done only between currencies of the same bank", err.Error())
}

func TestSubtract(t *testing.T) {
	bank, err := money.NewBankFromStaticExchangeRatesTable([]money.Currency{money.EUR, money.USD}, money.ExchangeRatesTable{"USD": {"EUR": 0.8}})
	assert.Nil(t, err)

	m1, err := bank.NewMoney(100, "EUR")
	assert.Nil(t, err)
	m2, err := bank.NewMoney(110, "EUR")
	assert.Nil(t, err)
	m3, err := m1.Subtract(m2)
	assert.Nil(t, err)
	assert.Equal(t, "€-0,10", m3.Format())
	assert.NotEqual(t, m1, m3)
	assert.NotEqual(t, m2, m3)

	m2, err = bank.NewMoney(100, "USD")
	assert.Nil(t, err)
	m3, err = m1.Subtract(m2)
	assert.Nil(t, err)
	assert.Equal(t, "€0,20", m3.Format())
	assert.NotEqual(t, m1, m3)
	assert.NotEqual(t, m2, m3)

	m2, err = money.NewMoney(100, "USD")
	assert.Nil(t, err)
	_, err = m1.Subtract(m2)
	assert.NotNil(t, err)
	assert.Equal(t, "currencies have different banks: operation between currencies can be done only between currencies of the same bank", err.Error())
}

func TestMultiply(t *testing.T) {
	m1, err := money.NewMoney(100, "EUR")
	assert.Nil(t, err)
	m2 := m1.Multiply(-3)
	assert.Equal(t, "€-3,00", m2.Format())
}

func TestSplit(t *testing.T) {
	m, err := money.NewMoney(101, "EUR")
	assert.Nil(t, err)

	_, err = m.Split(0)
	assert.NotNil(t, err)
	assert.Equal(t, "split must be higher than zero", err.Error())

	r, err := m.Split(3)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(r))
	assert.Equal(t, "€0,34", r[0].Format())
	assert.Equal(t, "€0,34", r[1].Format())
	assert.Equal(t, "€0,33", r[2].Format())
}

func ExampleMoney_Split() {
	m, _ := money.NewMoney(101, "EUR")
	parts, _ := m.Split(3)
	fmt.Println(parts[0].Format())
	fmt.Println(parts[1].Format())
	fmt.Println(parts[2].Format())
	// Output: €0,34
	// €0,34
	// €0,33
}

func TestAmount(t *testing.T) {
	m, err := money.NewMoney(1234, "EUR")
	assert.Nil(t, err)
	assert.Equal(t, 12.34, m.Amount())
}

func ExampleMoney_Amount() {
	m, _ := money.NewMoney(123, "USD")
	fmt.Println(m.Amount())
	// Output: 1.23
}

func TestFormat(t *testing.T) {
	m, err := money.NewMoney(1200000034, "EUR")
	assert.Nil(t, err)
	assert.Equal(t, "€12.000.000,34", m.Format())

	m, err = money.NewMoney(1200000034, "JPY")
	assert.Nil(t, err)
	assert.Equal(t, "¥1,200,000,034", m.Format())

	m, err = money.NewMoney(1200000034, "JOD")
	assert.Nil(t, err)
	assert.Equal(t, "د.ا1,200,000.034", m.Format())

	bank, err := money.NewBank([]money.Currency{
		{
			IsoCode:            "PIO",
			Name:               "Pioz Dollar",
			ThousandsSeparator: '^',
			DecimalMark:        '/',
			Symbol:             "#",
			SymbolFirst:        false,
			SubunitToUnit:      1000,
		},
	}, nil, nil)
	assert.Nil(t, err)

	m, err = bank.NewMoney(-1200000034, "PIO")
	assert.Nil(t, err)
	assert.Equal(t, "-1^200^000/034#", m.Format())
}
