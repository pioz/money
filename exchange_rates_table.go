package money

import (
	"bytes"
	"database/sql/driver"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// ExchangeRates is a map of currency ISO code to the exchange rate. See
// ExchangeRatesTable type for more details.
type ExchangeRates map[string]float64

// ExchangeRatesTable represents an exchange rates table, a map of currency ISO
// code to a map of currency ISO code to the exchange rate.
//
//   table := map[string]map[string]float64{
//     "USD": { "EUR": 0.87, "GBP": 0.74 },
//     "EUR": { "USD": 1.16, "GBP": 0.86 },
//     "GBP": { "USD": 1.35, "EUR": 1.17 },
//   }
//
// In this example to convert 1 USD to 1 EUR you have to multiply 1 * 0.87. So
// to get the exchange rate to convert c1 into c2 you have to access the map
// with
//    table[c1][c2]
type ExchangeRatesTable map[string]ExchangeRates

// Scan scan value into Json, implements sql.Scanner interface.
func (rates *ExchangeRates) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal ExchangeRates JSON value:", value))
	}
	return json.Unmarshal(bytes, rates)
}

// Value return json value, implement driver.Valuer interface.
func (rates ExchangeRates) Value() (driver.Value, error) {
	if len(rates) == 0 {
		return nil, nil
	}
	return json.Marshal(rates)
}

// Read implements the standard Read interface: fill b from table.
func (table *ExchangeRatesTable) Read(b []byte) (int, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(table)
	if err != nil {
		return 0, err
	}
	n := copy(b, buf.Bytes())
	return n, io.EOF
}

// Write implements the standard Write interface: load table from b.
func (table *ExchangeRatesTable) Write(b []byte) (int, error) {
	reader := bytes.NewReader(b)
	dec := gob.NewDecoder(reader)
	err := dec.Decode(table)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}

// ExchangeRatesTableCache is the interface that can be implemented to cache an
// exchange rates table.
type ExchangeRatesTableCache interface {
	Read() (ExchangeRatesTable, error)
	Write(table ExchangeRatesTable) error
}
