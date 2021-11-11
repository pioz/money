// Money is a Go package that helps you to work with currencies. It is based on
// Fowler's Money pattern and provides the ability to work with monetary value
// using a currency's smallest unit. Money supports currency exchange so you can
// convert easily, for example, 1 dollar in 1 euro.
package money

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Money represents a monetary value in a specific currency.
type Money struct {
	// Factional value of the monetary value.
	Cents int
	// ISO code of the currency of the monetary value.
	Currency string
	bank     *Bank
}

// ExchangeTo creates a new money in the currency with ISO code currencyIsoCode
// converted from m, using the exchange rate value stored in the bank exchange
// rates table. Returns an error if the currencyIsoCode is not supported by the
// current bank or if the bank is not able to exchange m.Currency to
// currencyIsoCode.
func (m *Money) ExchangeTo(currencyIsoCode string) (*Money, error) {
	if m.Currency == currencyIsoCode {
		return m.bank.NewMoney(m.Cents, m.Currency)
	}
	rate, err := m.bank.GetExchangeRate(m.Currency, currencyIsoCode)
	if err != nil {
		return nil, err
	}
	fromCurrency := m.bank.Currencies[m.Currency]
	toCurrency := m.bank.Currencies[currencyIsoCode]
	fractional := float64(m.Cents) / (float64(fromCurrency.SubunitToUnit) / float64(toCurrency.SubunitToUnit))
	exchangedCents := int(math.Round(fractional * rate))
	return m.bank.NewMoney(exchangedCents, currencyIsoCode)
}

// ExchangeTo creates a new money in the currency with ISO code currencyIsoCode
// converted from m, using the given exchange rate. Returns an error if the
// currencyIsoCode is not supported by the current bank.
func (m *Money) ExchangeToWithRate(currencyIsoCode string, rate float64) (*Money, error) {
	if m.Currency == currencyIsoCode {
		return m.bank.NewMoney(m.Cents, m.Currency)
	}
	fromCurrency := m.bank.Currencies[m.Currency]
	toCurrency := m.bank.Currencies[currencyIsoCode]
	fractional := float64(m.Cents) / (float64(fromCurrency.SubunitToUnit) / float64(toCurrency.SubunitToUnit))
	exchangedCents := int(math.Round(fractional * rate))
	return m.bank.NewMoney(exchangedCents, currencyIsoCode)
}

// Equals returns true if m1 and m2 have the same monetaty value. If m2's
// currency is different from m1's currency, m2 will be exchanged to m1's
// currency. Returns error if the bank that generated m1 is different from the
// bank that generated m2.
func (m1 *Money) Equals(m2 *Money) (bool, error) {
	result, err := prepareOperation(m1, m2)
	if err != nil {
		return false, err
	}
	return m1.Cents == result.Cents, nil
}

// GreaterThan returns true if m1 is greater than m2. If m2's currency is
// different from m1's currency, m2 will be exchanged to m1's currency. Returns
// error if the bank that generated m1 is different from the bank that generated
// m2.
func (m1 *Money) GreaterThan(m2 *Money) (bool, error) {
	result, err := prepareOperation(m1, m2)
	if err != nil {
		return false, err
	}
	return m1.Cents > result.Cents, nil
}

// GreaterThan returns true if m1 is greater than or equal to m2. If m2's
// currency is different from m1's currency, m2 will be exchanged to m1's
// currency. Returns error if the bank that generated m1 is different from the
// bank that generated m2.
func (m1 *Money) GreaterThanOrEqual(m2 *Money) (bool, error) {
	result, err := prepareOperation(m1, m2)
	if err != nil {
		return false, err
	}
	return m1.Cents >= result.Cents, nil
}

// GreaterThan returns true if m1 is less than m2. If m2's currency is different
// from m1's currency, m2 will be exchanged to m1's currency. Returns error if
// the bank that generated m1 is different from the bank that generated m2.
func (m1 *Money) LessThan(m2 *Money) (bool, error) {
	result, err := prepareOperation(m1, m2)
	if err != nil {
		return false, err
	}
	return m1.Cents < result.Cents, nil
}

// GreaterThan returns true if m1 is less than or equal to m2. If m2's currency
// is different from m1's currency, m2 will be exchanged to m1's currency.
// Returns error if the bank that generated m1 is different from the bank that
// generated m2.
func (m1 *Money) LessThanOrEqual(m2 *Money) (bool, error) {
	result, err := prepareOperation(m1, m2)
	if err != nil {
		return false, err
	}
	return m1.Cents <= result.Cents, nil
}

// IsZero returns true if m moneraty value is zero.
func (m *Money) IsZero() bool {
	return m.Cents == 0
}

// IsNegative returns true if m moneraty value is negative.
func (m *Money) IsNegative() bool {
	return m.Cents < 0
}

// IsPositive returns true if m moneraty value is positive.
func (m *Money) IsPositive() bool {
	return m.Cents > 0
}

// Absolute returns the absolute moneraty value of m.
func (m *Money) Absolute() *Money {
	cents := int(math.Abs(float64(m.Cents)))
	absMoney, _ := m.bank.NewMoney(cents, m.Currency)
	return absMoney
}

// Add returns a new money with monetary value equals to m1 + m2. If m2's
// currency is different from m1's currency, m2 will be exchanged to m1's
// currency. Returns error if the bank that generated m1 is different from the
// bank that generated m2.
func (m1 *Money) Add(m2 *Money) (*Money, error) {
	result, err := prepareOperation(m1, m2)
	if err != nil {
		return nil, err
	}
	result.Cents += m1.Cents
	return result, nil
}

// Subtract returns a new money with monetary value equals to m1 - m2. If m2's
// currency is different from m1's currency, m2 will be exchanged to m1's
// currency. Returns error if the bank that generated m1 is different from the
// bank that generated m2.
func (m1 *Money) Subtract(m2 *Money) (*Money, error) {
	result, err := prepareOperation(m1, m2)
	if err != nil {
		return nil, err
	}
	result.Cents = m1.Cents - result.Cents
	return result, nil
}

// Multiply returns a new money with monetary value equals to m1 * mul.
func (m *Money) Multiply(mul int) *Money {
	result, _ := m.bank.NewMoney(m.Cents*mul, m.Currency)
	return result
}

// Split returns a slice of money with split monetary value in the given number.
// After division leftover pennies will be distributed round-robin amongst the
// parties. This means that parties listed first will likely receive more
// pennies than ones that are listed later.
func (m *Money) Split(parts int) ([]*Money, error) {
	if parts <= 0 {
		return nil, errors.New("split must be higher than zero")
	}
	cents := m.Cents / parts
	remainder := m.Cents % parts
	results := make([]*Money, 0, parts)
	for i := 0; i < parts; i++ {
		moneyPart, _ := m.bank.NewMoney(cents, m.Currency)
		results = append(results, moneyPart)
	}
	for i := 0; i < remainder; i++ {
		results[i].Cents += 1
	}
	return results, nil
}

// Amount returns the numerical value of the money.
func (m *Money) Amount() float64 {
	currency := m.bank.Currencies[m.Currency]
	return float64(m.Cents) / float64(currency.SubunitToUnit)
}

// Format creates a formatted price string according to m's currency fields.
func (m *Money) Format() string {
	currency := m.bank.Currencies[m.Currency]
	amountToS := commaf(m.Amount(), currency.ThousandsSeparator, currency.DecimalMark, int(math.Log10(float64(currency.SubunitToUnit))))
	if currency.SymbolFirst {
		return fmt.Sprintf("%s%s", currency.Symbol, amountToS)
	}
	return fmt.Sprintf("%s%s", amountToS, currency.Symbol)
}

// Private functions

func prepareOperation(m1, m2 *Money) (*Money, error) {
	if m1.bank != m2.bank {
		return nil, fmt.Errorf("currencies have different banks: operation between currencies can be done only between currencies of the same bank")
	}
	return m2.ExchangeTo(m1.Currency)
}

func commaf(v float64, thousandsSeparator, decimalMark rune, precision int) string {
	buf := &bytes.Buffer{}
	if v < 0 {
		buf.Write([]byte{'-'})
		v = 0 - v
	}

	comma := []byte(string(thousandsSeparator))

	parts := strings.Split(strconv.FormatFloat(v, 'f', precision, 64), ".")
	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.Write(comma)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.Write(comma)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.Write([]byte(string(decimalMark)))
		buf.WriteString(parts[1])
	}
	return buf.String()
}
