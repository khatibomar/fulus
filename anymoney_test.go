package fulus

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
)

func TestAnyMoneyConstructors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		create  func() (AnyMoney, error)
		want    int64
		code    string
		wantErr error
	}{
		{
			name:   "from currency",
			create: func() (AnyMoney, error) { return NewAnyMoney(1050, currency.EUR{}), nil },
			want:   1050,
			code:   "EUR",
		},
		{
			name:   "from code",
			create: func() (AnyMoney, error) { return NewAnyMoneyFromCode(-5, "jpy") },
			want:   -5,
			code:   "JPY",
		},
		{
			name:   "from typed money",
			create: func() (AnyMoney, error) { return NewMoney[currency.USD](99).Any(), nil },
			want:   99,
			code:   "USD",
		},
		{
			name:   "parse decimal",
			create: func() (AnyMoney, error) { return ParseAnyMoney("12.345", "BHD") },
			want:   12345,
			code:   "BHD",
		},
		{
			name:    "unknown code",
			create:  func() (AnyMoney, error) { return NewAnyMoneyFromCode(1, "XYZ") },
			wantErr: ErrUnknownCurrency,
		},
		{
			name:    "parse unknown code",
			create:  func() (AnyMoney, error) { return ParseAnyMoney("1", "XYZ") },
			wantErr: ErrUnknownCurrency,
		},
		{
			name:    "parse scale too large",
			create:  func() (AnyMoney, error) { return ParseAnyMoney("1.5", "JPY") },
			wantErr: ErrScaleMismatch,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.create()
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error = %v, want %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}
			if got.amount64() != tt.want || got.Currency().Code() != tt.code {
				t.Errorf("got %d %s, want %d %s", got.amount64(), got.Currency().Code(), tt.want, tt.code)
			}
		})
	}
}

func TestAs(t *testing.T) {
	t.Parallel()

	usd, err := As[currency.USD](NewAnyMoney(1050, currency.USD{}))
	if err != nil || usd.amount64() != 1050 {
		t.Fatalf("As[USD]() = %v, %v; want 1050, nil", usd.amount64(), err)
	}

	if _, err := As[currency.EUR](NewAnyMoney(1050, currency.USD{})); !errors.Is(err, ErrCurrencyMismatch) {
		t.Errorf("As[EUR]() error = %v, want %v", err, ErrCurrencyMismatch)
	}
	if _, err := As[currency.USD](AnyMoney{}); !errors.Is(err, ErrCurrencyMismatch) {
		t.Errorf("As[USD](zero value) error = %v, want %v", err, ErrCurrencyMismatch)
	}
}

func TestAnyMoneyArithmetic(t *testing.T) {
	t.Parallel()

	usd := func(a int64) AnyMoney { return NewAnyMoney(a, currency.USD{}) }
	eur := NewAnyMoney(100, currency.EUR{})

	tests := []struct {
		name    string
		op      func() (AnyMoney, error)
		want    int64
		wantErr error
	}{
		{name: "add", op: func() (AnyMoney, error) { return usd(100).Add(usd(50)) }, want: 150},
		{name: "sub", op: func() (AnyMoney, error) { return usd(100).Sub(usd(150)) }, want: -50},
		{name: "add mismatch", op: func() (AnyMoney, error) { return usd(100).Add(eur) }, wantErr: ErrCurrencyMismatch},
		{name: "add zero value", op: func() (AnyMoney, error) { return AnyMoney{}.Add(usd(1)) }, wantErr: ErrCurrencyMismatch},
		{name: "add overflow", op: func() (AnyMoney, error) { return maxMoney[currency.USD]().Any().Add(usd(1)) }, wantErr: ErrOverflow},
		{name: "sub overflow", op: func() (AnyMoney, error) { return minMoney[currency.USD]().Any().Sub(usd(1)) }, wantErr: ErrOverflow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.op()
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error = %v, want %v", err, tt.wantErr)
			}
			if err == nil && got.amount64() != tt.want {
				t.Errorf("amount = %d, want %d", got.amount64(), tt.want)
			}
		})
	}

	if c, err := usd(1).Cmp(usd(2)); err != nil || c != -1 {
		t.Errorf("Cmp() = %d, %v; want -1, nil", c, err)
	}
	if _, err := usd(1).Cmp(eur); !errors.Is(err, ErrCurrencyMismatch) {
		t.Errorf("Cmp() mismatch error = %v", err)
	}
}

func TestAnyMoneyFormat(t *testing.T) {
	t.Parallel()

	if got := NewAnyMoney(-123456, currency.CHF{}).Format(locale.DE_CH); got != "CHF-1'234.56" {
		t.Errorf("Format() = %q", got)
	}
	if got := (AnyMoney{amount: int128FromInt64(42)}).String(); got != "42" {
		t.Errorf("zero currency String() = %q, want %q", got, "42")
	}
}

func TestAnyMoneyJSON(t *testing.T) {
	t.Parallel()

	b, err := json.Marshal(NewAnyMoney(1050, currency.EUR{}))
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	if string(b) != `{"amount":"1050","currency":"EUR"}` {
		t.Fatalf("Marshal() = %s", b)
	}

	var typed Money[currency.EUR]
	if err := json.Unmarshal(b, &typed); err != nil || typed.amount64() != 1050 {
		t.Fatalf("Money JSON is not compatible: %v, %d", err, typed.amount64())
	}

	tests := []struct {
		name    string
		input   string
		want    int64
		code    string
		wantErr error
	}{
		{name: "valid", input: `{"amount":"-7","currency":"KWD"}`, want: -7, code: "KWD"},
		{name: "unknown currency", input: `{"amount":"1","currency":"XYZ"}`, wantErr: ErrUnknownCurrency},
		{name: "invalid amount", input: `{"amount":"1.5","currency":"USD"}`, wantErr: ErrInvalidAmountFormat},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var m AnyMoney
			err := json.Unmarshal([]byte(tt.input), &m)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Unmarshal() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil && (m.amount64() != tt.want || m.Currency().Code() != tt.code) {
				t.Errorf("Unmarshal() = %d %s", m.amount64(), m.Currency().Code())
			}
		})
	}
}
