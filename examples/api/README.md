# Fulus API Example

This example demonstrates how to use `fulus` in a real-world HTTP API context.

It implements a simple `/checkout` endpoint that accepts a JSON payload representing a shopping cart. The payload includes prices for each item. 

## Key features demonstrated:
- Unmarshaling JSON request payloads directly into `fulus.Money` types securely.
- Safely adding up multiple money amounts in a loop to calculate a subtotal.
- Applying a percentage tax (8.5%) using `Mul` and `Div` with explicit rounding (`fulus.RoundHalfUp`).
- Marshaling `fulus.Money` types securely back into a JSON response.

## How to run

Start the server:

```bash
go run main.go
```

Test it with a curl command:

```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"cart": [{"id": "1", "name": "Gopher Plush", "price": {"amount": "1500", "currency": "USD"}}, {"id": "2", "name": "Go Sticker", "price": {"amount": "300", "currency": "USD"}}]}' \
  http://localhost:8080/checkout
```

### Expected Output:

```json
{
  "subtotal": {
    "amount": "1800",
    "currency": "USD"
  },
  "tax": {
    "amount": "153",
    "currency": "USD"
  },
  "total": {
    "amount": "1953",
    "currency": "USD"
  },
  "message": "Checkout successful. Total: $19.53"
}
```
