package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/khatibomar/fulus"
	"github.com/khatibomar/fulus/currency"
)

type Product struct {
	ID    string                    `json:"id"`
	Name  string                    `json:"name"`
	Price fulus.Money[currency.USD] `json:"price"`
}

type CheckoutRequest struct {
	Cart []Product `json:"cart"`
}

type CheckoutResponse struct {
	Subtotal fulus.Money[currency.USD] `json:"subtotal"`
	Tax      fulus.Money[currency.USD] `json:"tax"`
	Total    fulus.Money[currency.USD] `json:"total"`
	Message  string                    `json:"message"`
}

func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Calculate subtotal
	subtotal := fulus.NewMoney[currency.USD](0)
	for _, item := range req.Cart {
		var err error
		subtotal, err = subtotal.Add(item.Price)
		if err != nil {
			http.Error(w, "Error calculating subtotal: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Calculate 8.5% tax
	// 8.5% = 85 / 1000
	tax, err := subtotal.Mul(85)
	if err != nil {
		http.Error(w, "Error calculating tax", http.StatusInternalServerError)
		return
	}
	tax, err = tax.Div(1000, fulus.RoundHalfUp)
	if err != nil {
		http.Error(w, "Error rounding tax", http.StatusInternalServerError)
		return
	}

	// Calculate total
	total, err := subtotal.Add(tax)
	if err != nil {
		http.Error(w, "Error calculating total", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CheckoutResponse{
		Subtotal: subtotal,
		Tax:      tax,
		Total:    total,
		Message:  fmt.Sprintf("Checkout successful. Total: %s", total.String()),
	})
}

func main() {
	http.HandleFunc("/checkout", checkoutHandler)

	fmt.Println("Server starting on :8080...")
	fmt.Println(`Try testing with:
curl -X POST -H "Content-Type: application/json" \
  -d '{"cart": [{"id": "1", "name": "Gopher Plush", "price": {"amount": "1500", "currency": "USD"}}, {"id": "2", "name": "Go Sticker", "price": {"amount": "300", "currency": "USD"}}]}' \
  http://localhost:8080/checkout`)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
