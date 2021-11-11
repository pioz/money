package banks

import "github.com/pioz/money"

func NewOneBank(currencies []money.Currency) (*money.Bank, error) {
	return money.NewBank(money.AllCurrencies, func() (money.ExchangeRatesTable, error) {
		table := make(money.ExchangeRatesTable)
		for _, fromCurrency := range money.AllCurrencies {
			table[fromCurrency.IsoCode] = make(money.ExchangeRates)
			for _, toCurrency := range money.AllCurrencies {
				if fromCurrency.IsoCode == toCurrency.IsoCode {
					continue
				}
				table[fromCurrency.IsoCode][toCurrency.IsoCode] = float64(fromCurrency.SubunitToUnit) / float64(toCurrency.SubunitToUnit)
			}
		}
		return table, nil
	}, nil)
}
