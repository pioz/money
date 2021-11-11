package money_test

import (
	"os"
	"testing"

	"github.com/pioz/money"
	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	table := money.ExchangeRatesTable{
		"EUR": {"USD": 1.2},
		"USD": {"EUR": 0.8},
	}

	brokenFileCache := money.ExchangeRatesTableFileCache{FilePath: "/tmp/not-found/not-found"}
	err := brokenFileCache.Write(table)
	assert.NotNil(t, err)
	assert.Equal(t, "open /tmp/not-found/not-found: no such file or directory", err.Error())

	_, err = brokenFileCache.Read()
	assert.NotNil(t, err)
	assert.Equal(t, "open /tmp/not-found/not-found: no such file or directory", err.Error())

	fileCache := money.ExchangeRatesTableFileCache{FilePath: "/tmp/go-money-exchange-rates-table-cache"}
	defer os.RemoveAll(fileCache.FilePath)

	err = fileCache.Write(table)
	assert.Nil(t, err)

	reloadedRates, err := fileCache.Read()
	assert.Nil(t, err)

	assert.Equal(t, 1.2, reloadedRates["EUR"]["USD"])
	assert.Equal(t, 0.8, reloadedRates["USD"]["EUR"])
}
