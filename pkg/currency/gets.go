package currency

import (
	"fmt"
	"io"
	"net/http"
)

const (
	CURRENCY_LIST_URL     = "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies.json"
	CURRENCY_EXCHANGE_URL = "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/%s.json"
)

var client = http.DefaultClient

func GetAll() (string, error) {
	resp, err := client.Get(CURRENCY_LIST_URL)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error unexpected response: %v (%v) %v", resp.StatusCode, resp.Status, string(data))
	}

	return string(data), nil
}

func GetExchanges(currency string) (string, error) {
	resp, err := client.Get(fmt.Sprintf(CURRENCY_EXCHANGE_URL, currency))
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error unexpected response: %v (%v) %v", resp.StatusCode, resp.Status, string(data))
	}

	return string(data), nil
}
