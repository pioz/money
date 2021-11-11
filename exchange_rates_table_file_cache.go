package money

import (
	"io"
	"os"
)

// ExchangeRatesTableFileCache implements the ExchangeRatesTableCache
// interface to cache an exchange rates table into a file.
type ExchangeRatesTableFileCache struct {
	// File path where to store the cache.
	FilePath string
}

// Read implements the Read method of ExchangeRatesTableCache interface.
func (c ExchangeRatesTableFileCache) Read() (ExchangeRatesTable, error) {
	var table ExchangeRatesTable
	f, err := os.Open(c.FilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	_, err = io.Copy(&table, f)
	if err != nil {
		return nil, err
	}
	return table, nil
}

// Write implements the Write method of ExchangeRatesTableCache interface.
func (c ExchangeRatesTableFileCache) Write(table ExchangeRatesTable) error {
	f, err := os.OpenFile(c.FilePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, &table)
	if err != nil {
		return err
	}
	return nil
}
