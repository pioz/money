# Money

Money is a Go package that helps you to work with currencies. It is based on
[Fowler's Money pattern](https://martinfowler.com/eaaCatalog/money.html) and
provides the ability to work with monetary value using a currency's smallest
unit. So with this package, you can perform precise operations without being
afraid of losing a single cent.

Money supports **currency exchange** so you can convert easily, for example, 1
dollar in 1 euro.

## Usage

```go
package main

import (
  "fmt"
  "github.com/pioz/money"
)

func main() {
  m1, _ := money.NewMoney(100, "USD")
  m2, _ := money.NewMoneyFromAmount(4.00, "USD")
  m, _ := m1.Add(m2)
  fmt.Println(m.Format()) // Output: $5.00
  m = m.Multiply(2)
  fmt.Println(m.Format()) // Output: $10.00

  parts, _ := m.Split(3)
  fmt.Println(parts[0].Format()) // Output: $3.34
  fmt.Println(parts[1].Format()) // Output: $3.33
  fmt.Println(parts[2].Format()) // Output: $3.33
}
```

## Bank

A bank is a type in the money package that allows you to create money. A bank
define in which currencies you can create money and the exchange rates table to
exchange monetary value in a currency to another. A sample exchange rates table
is this:

```go
exchangeRatesTable := map[string]map[string]float64{
  "USD": { "EUR": 0.87, "GBP": 0.74 },
  "EUR": { "USD": 1.16, "GBP": 0.86 },
  "GBP": { "USD": 1.35, "EUR": 1.17 },
}
```

And here is an example of how to use a bank:

```go
package main

import (
  "fmt"
  "log"
  "github.com/pioz/money"
)

func main() {
  fetchExchangeRatesTable := func() (money.Rates, error) {
    var table money.ExchangeRatesTable
    // fetch table from Internet
    return table, nil
  }
  bank, _ := money.NewBank([]money.Currency{money.EUR, money.USD}, fetchExchangeRatesTable, nil)
  usd, _ := bank.NewMoney(100, "USD")
  eur, _ := usd.ExchangeTo("EUR")
  fmt.Println(eur.Format())

  bank.UpdateExchangeRatesTable() // Update exchange rates table

  eur, _ = usd.ExchangeTo("EUR")
  fmt.Println(eur.Format())

  _, err := bank.NewMoney(1, "JPY")
  if err != nil {
    log.Println(err.Error()) // Output: Bank does not support JPY currency
  }
}
```

You can create a bank with an `ExchangeRatesTableCache` that load the exchange
rates table from this cache, while a go routine in background updates the
exchange rates table.

```go
  fetchExchangeRatesTable := func() (money.Rates, error) {
    var table money.ExchangeRatesTable
    // fetch table from Internet
    return table, nil
  }
  fileCache := money.RatesFileCache{FilePath: "/tmp/go-money-exchange-rates-table-cache"}
  bank, _ := money.NewBank(money.AvailableCurrencies, fetchExchangeRatesTable, fileCache)
  usd, _ := bank.NewMoney(100, "USD")
  eur, _ := usd.ExchangeTo("EUR")
  fmt.Println(eur.Format())
```

Money package provides a simple file cache type that read and write the exchange
rates table in a file, but you can create your custom type simply by
implementing the `ExchangeRatesTableCache` interface, for example, to use Redis
or something else.

### Freecurrency bank

Money package allows you to create out of the box a bank that can fetch the
exchange rates table from https://freecurrencyapi.net/

```go
import "github.com/pioz/money/banks"

func main() {
  bank, err := banks.NewFreecurrencyBank([]money.Currency{money.EUR, money.USD}, "YOUR-API-KEY", nil)
  m, _ := bank.NewMoney(100, "EUR")
	ex, _ := m.ExchangeTo("USD")
}
```

## Currencies

The money package has all real-life currencies pre-defined. But you can also
define a new currency for your realm :) See the go doc to see how the currency
type is like.

## Contributing

**HELP**: all contributions to improve the go doc are very very welcome!

Bug reports and pull requests are welcome on GitHub at
https://github.com/pioz/money/issues.

## License

The package is available as open source under the terms of the [MIT
License](http://opensource.org/licenses/MIT).
