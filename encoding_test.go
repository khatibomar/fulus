package fulus

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"math"
	"testing"

	"github.com/khatibomar/fulus/currency"
)

func TestDecimalAndText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value interface {
			Decimal() string
			MarshalText() ([]byte, error)
		}
		want string
	}{
		{name: "usd", value: NewMoney[currency.USD](-123450), want: "-1234.50"},
		{name: "usd small", value: NewMoney[currency.USD](5), want: "0.05"},
		{name: "jpy", value: NewMoney[currency.JPY](1234), want: "1234"},
		{name: "bhd", value: NewMoney[currency.BHD](1), want: "0.001"},
		{name: "min int64", value: NewMoney[currency.USD](math.MinInt64), want: "-92233720368547758.08"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.value.Decimal(); got != tt.want {
				t.Errorf("Decimal() = %q, want %q", got, tt.want)
			}
			b, err := tt.value.MarshalText()
			if err != nil || string(b) != tt.want {
				t.Errorf("MarshalText() = %q, %v; want %q", b, err, tt.want)
			}
		})
	}

	var m Money[currency.USD]
	if err := m.UnmarshalText([]byte("-1234.50")); err != nil || m.amount64() != -123450 {
		t.Errorf("UnmarshalText() = %d, %v", m.amount64(), err)
	}
	if err := m.UnmarshalText([]byte("1.234")); !errors.Is(err, ErrScaleMismatch) {
		t.Errorf("UnmarshalText() error = %v, want %v", err, ErrScaleMismatch)
	}

	b, err := json.Marshal(map[Money[currency.USD]]string{NewMoney[currency.USD](150): "x"})
	if err != nil || string(b) != `{"1.50":"x"}` {
		t.Errorf("map key Marshal() = %s, %v", b, err)
	}
}

func TestLogValue(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	}))
	logger.Info("paid", "price", NewMoney[currency.EUR](1050))

	want := "level=INFO msg=paid price.amount=10.50 price.currency=EUR\n"
	if buf.String() != want {
		t.Errorf("log = %q, want %q", buf.String(), want)
	}
}

func TestSumMinMax(t *testing.T) {
	t.Parallel()

	usd := NewMoney[currency.USD]
	tests := []struct {
		name    string
		values  []Money[currency.USD]
		want    int64
		wantErr error
	}{
		{name: "empty", values: nil, want: 0},
		{name: "values", values: []Money[currency.USD]{usd(100), usd(-30), usd(5)}, want: 75},
		{name: "intermediate overflow", values: []Money[currency.USD]{usd(math.MaxInt64), usd(1), usd(-2)}, want: math.MaxInt64 - 1},
		{name: "intermediate overflow of 128 bits", values: []Money[currency.USD]{maxMoney[currency.USD](), usd(1), maxMoney[currency.USD]().MustMul(-1)}, want: 1},
		{name: "overflow", values: []Money[currency.USD]{maxMoney[currency.USD](), usd(1)}, wantErr: ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Sum(tt.values...)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Sum() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil && got.amount64() != tt.want {
				t.Errorf("Sum() = %d, want %d", got.amount64(), tt.want)
			}
		})
	}

	if got := Min(usd(3), usd(-1), usd(2)); got.amount64() != -1 {
		t.Errorf("Min() = %d, want -1", got.amount64())
	}
	if got := Max(usd(3), usd(-1), usd(7)); got.amount64() != 7 {
		t.Errorf("Max() = %d, want 7", got.amount64())
	}
	if got := Max(usd(3)); got.amount64() != 3 {
		t.Errorf("Max() with one value = %d, want 3", got.amount64())
	}
}

type quotedCodeCurrency struct{}

func (quotedCodeCurrency) Code() string    { return `Q"<&>` }
func (quotedCodeCurrency) MinorUnits() int { return 2 }

func TestMarshalJSONEscapesCode(t *testing.T) {
	t.Parallel()

	got, err := NewMoney[quotedCodeCurrency](-5).MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON() error = %v", err)
	}
	want, _ := json.Marshal(moneyJSON{Amount: "-0.05", Currency: `Q"<&>`})
	if string(got) != string(want) {
		t.Errorf("MarshalJSON() = %s, want %s", got, want)
	}

	var back Money[quotedCodeCurrency]
	if err := json.Unmarshal(got, &back); err != nil || back.amount64() != -5 {
		t.Errorf("Unmarshal() = %d, %v", back.amount64(), err)
	}
}
