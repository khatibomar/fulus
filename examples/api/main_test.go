package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckoutHandler(t *testing.T) {
	reqBody := []byte(`{
		"cart": [
			{"id": "1", "name": "Gopher Plush", "price": {"amount": "1500", "currency": "USD"}},
			{"id": "2", "name": "Go Sticker", "price": {"amount": "300", "currency": "USD"}}
		]
	}`)

	req, err := http.NewRequest(http.MethodPost, "/checkout", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(checkoutHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resp CheckoutResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}

	expectedTotalAmount := int64(1953) // 1500 + 300 = 1800. Tax = 1800 * 0.085 = 153. Total = 1953
	if resp.Total.Amount() != expectedTotalAmount {
		t.Errorf("handler returned unexpected total: got %v want %v", resp.Total.Amount(), expectedTotalAmount)
	}
}
