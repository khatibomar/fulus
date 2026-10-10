package fulus

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/khatibomar/fulus/currency"
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
	_, err = As[currency.USD](NewAnyMoney(10000, fourDigitUSD{}))
	if !errors.Is(err, ErrCurrencyMismatch) {
		t.Fatalf("As[USD]() with other minor units error = %v, want %v", err, ErrCurrencyMismatch)
	}
	if want := "currency mismatch: expected USD with 2 minor units, got USD with 4 minor units"; err.Error() != want {
		t.Errorf("As[USD]() error = %q, want %q", err, want)
	}
}

// fourDigitUSD has the code of USD but other minor units.
type fourDigitUSD struct{}

func (fourDigitUSD) Code() string    { return "USD" }
func (fourDigitUSD) MinorUnits() int { return 4 }

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
		{name: "add other minor units", op: func() (AnyMoney, error) { return usd(100).Add(NewAnyMoney(1, fourDigitUSD{})) }, wantErr: ErrCurrencyMismatch},
		{name: "mul", op: func() (AnyMoney, error) { return usd(-25).Mul(4) }, want: -100},
		{name: "mul overflow", op: func() (AnyMoney, error) { return maxMoney[currency.USD]().Any().Mul(2) }, wantErr: ErrOverflow},
		{name: "div", op: func() (AnyMoney, error) { return usd(100).Div(3, RoundHalfUp) }, want: 33},
		{name: "div rounds", op: func() (AnyMoney, error) { return usd(-5).Div(2, RoundFloor) }, want: -3},
		{name: "div by zero", op: func() (AnyMoney, error) { return usd(100).Div(0, RoundHalfUp) }, wantErr: ErrDivisionByZero},
		{name: "div invalid mode", op: func() (AnyMoney, error) { return usd(100).Div(3, 0) }, wantErr: ErrInvalidRoundingMode},
		{name: "mul factor", op: func() (AnyMoney, error) { return usd(1000).MulFactor(Percent(15), RoundHalfEven) }, want: 150},
		{name: "mul zero factor", op: func() (AnyMoney, error) { return usd(1000).MulFactor(Factor{}, RoundHalfEven) }, wantErr: ErrInvalidFactor},
		{name: "mul factor inexact", op: func() (AnyMoney, error) { return usd(1).MulFactor(Percent(50), RoundUnnecessary) }, wantErr: ErrInexact},
		{name: "round cash", op: func() (AnyMoney, error) { return NewAnyMoney(1003, currency.CHF{}).RoundCash(RoundHalfUp) }, want: 1005},
		{name: "round cash without rule", op: func() (AnyMoney, error) { return usd(1003).RoundCash(RoundHalfUp) }, want: 1003},
		{name: "round cash zero value", op: func() (AnyMoney, error) { return AnyMoney{}.RoundCash(RoundHalfUp) }, want: 0},
		{name: "round cash invalid mode", op: func() (AnyMoney, error) { return usd(1).RoundCash(0) }, wantErr: ErrInvalidRoundingMode},
		{name: "abs", op: func() (AnyMoney, error) { return usd(-7).Abs() }, want: 7},
		{name: "abs positive", op: func() (AnyMoney, error) { return usd(7).Abs() }, want: 7},
		{name: "neg", op: func() (AnyMoney, error) { return usd(7).Neg() }, want: -7},
		{name: "neg overflow", op: func() (AnyMoney, error) { return minMoney[currency.USD]().Any().Neg() }, wantErr: ErrOverflow},
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
	if got, err := usd(-7).Neg(); err != nil || got.Currency() != (currency.USD{}) {
		t.Errorf("Neg() = %v, %v; want the currency USD", got, err)
	}
}

func TestAnyMoneySign(t *testing.T) {
	t.Parallel()

	tests := []struct {
		amount             int64
		sign               int
		positive, negative bool
	}{
		{amount: -3, sign: -1, negative: true},
		{amount: 0, sign: 0},
		{amount: 3, sign: 1, positive: true},
	}
	for _, tt := range tests {
		m := NewAnyMoney(tt.amount, currency.EUR{})
		if m.Sign() != tt.sign || m.IsPositive() != tt.positive || m.IsNegative() != tt.negative {
			t.Errorf("%d: Sign() = %d, IsPositive() = %v, IsNegative() = %v", tt.amount, m.Sign(), m.IsPositive(), m.IsNegative())
		}
	}
}

func TestAnyMoneyAllocate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		amount  int64
		ratios  []int64
		want    []int64
		wantErr error
	}{
		{name: "thirds", amount: 100, ratios: []int64{1, 1, 1}, want: []int64{34, 33, 33}},
		{name: "zero ratio", amount: -100, ratios: []int64{0, 1, 1}, want: []int64{0, -50, -50}},
		{name: "no ratios", amount: 100, wantErr: ErrNoRatios},
		{name: "all zero", amount: 100, ratios: []int64{0}, wantErr: ErrInvalidRatios},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			parts, err := NewAnyMoney(tt.amount, currency.EUR{}).Allocate(tt.ratios...)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Allocate() error = %v, want %v", err, tt.wantErr)
			}
			if len(parts) != len(tt.want) {
				t.Fatalf("Allocate() = %v, want %v", parts, tt.want)
			}
			for i, p := range parts {
				if p.amount64() != tt.want[i] || p.Currency() != (currency.EUR{}) {
					t.Errorf("part %d = %v, want EUR %d", i, p, tt.want[i])
				}
			}
		})
	}
}

func TestAnyMoneyScanColumns(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		amount    any
		code      any
		codeFirst bool
		want      int64
		wantCode  string
		wantErr   error
	}{
		{name: "text", amount: []byte("-10.50"), code: []byte("eur"), want: -1050, wantCode: "EUR"},
		{name: "code first", amount: "0.007", code: "KWD", codeFirst: true, want: 7, wantCode: "KWD"},
		{name: "int64", amount: int64(3), code: "JPY", wantErr: ErrInvalidAmountFormat},
		{name: "float64", amount: 2.5, code: "USD", want: 250, wantCode: "USD"},
		{name: "too many fraction digits", amount: "1.005", code: "USD", wantErr: ErrScaleMismatch},
		{name: "unknown code", amount: "1", code: "XYZ", wantErr: ErrUnknownCurrency},
		{name: "NULL amount", amount: nil, code: "USD", wantErr: ErrInvalidAmountFormat},
		{name: "NULL code", amount: "1", code: nil, wantErr: ErrUnknownCurrency},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := NewAnyMoney(99, currency.CHF{})
			amount, code := m.ScanColumns()
			first, second := func() error { return amount.Scan(tt.amount) }, func() error { return code.Scan(tt.code) }
			if tt.codeFirst {
				first, second = second, first
			}
			err := first()
			if err == nil {
				if m.amount64() != 99 {
					t.Fatalf("m changed after one column: %v", m)
				}
				err = second()
			}
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Scan() error = %v, want %v", err, tt.wantErr)
			}
			if err != nil {
				if m.amount64() != 99 || m.Currency() != (currency.CHF{}) {
					t.Errorf("m changed after an error: %v", m)
				}
				return
			}
			if m.amount64() != tt.want || m.Currency().Code() != tt.wantCode {
				t.Errorf("Scan() = %v, want %d %s", m, tt.want, tt.wantCode)
			}
		})
	}
}

func TestAnyMoneyScanColumnsNextRow(t *testing.T) {
	t.Parallel()

	var m AnyMoney
	amount, code := m.ScanColumns()
	rows := []struct {
		amount, code string
		want         string
	}{
		{amount: "1.5", code: "XYZ"},
		{amount: "1.5", code: "USD", want: "USD 1.50"},
		{amount: "7", code: "EUR", want: "EUR 7.00"},
	}
	for _, row := range rows {
		err := amount.Scan(row.amount)
		if err == nil {
			err = code.Scan(row.code)
		}
		if row.want == "" {
			if err == nil {
				t.Fatalf("Scan(%q, %q) succeeded", row.amount, row.code)
			}
			continue
		}
		if err != nil || m.String() != row.want {
			t.Errorf("Scan(%q, %q) = %v, %v; want %s", row.amount, row.code, m, err, row.want)
		}
	}
}

func TestAnyMoneyJSON(t *testing.T) {
	t.Parallel()

	b, err := json.Marshal(NewAnyMoney(1050, currency.EUR{}))
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	if string(b) != `{"amount":"10.50","currency":"EUR"}` {
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
		{name: "valid", input: `{"amount":"-0.007","currency":"KWD"}`, want: -7, code: "KWD"},
		{name: "unknown currency", input: `{"amount":"1","currency":"XYZ"}`, wantErr: ErrUnknownCurrency},
		{name: "invalid amount", input: `{"amount":"1.5x","currency":"USD"}`, wantErr: ErrInvalidAmountFormat},
		{name: "too many fraction digits", input: `{"amount":"1.555","currency":"USD"}`, wantErr: ErrScaleMismatch},
		{name: "null does not change the value", input: `null`, want: 5, code: "CHF"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := NewAnyMoney(5, currency.CHF{})
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

func TestAnyMoneyAccessors(t *testing.T) {
	t.Parallel()

	m := NewAnyMoney(-1050, currency.EUR{})
	tests := []struct {
		name string
		got  any
		want any
	}{
		{"BigInt", m.BigInt().String(), "-1050"},
		{"IsZero", m.IsZero(), false},
		{"zero IsZero", NewAnyMoney(0, currency.EUR{}).IsZero(), true},
		{"Decimal", m.Decimal(), "-10.50"},
		{"String", m.String(), "EUR -10.50"},
		{"zero value Decimal", AnyMoney{}.Decimal(), "0"},
		{"zero value String", AnyMoney{}.String(), "0"},
	}
	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf("%s = %v, want %v", tt.name, tt.got, tt.want)
		}
	}

	b, err := json.Marshal(AnyMoney{})
	if err != nil || string(b) != `null` {
		t.Errorf("Marshal() of the zero value = %s, %v", b, err)
	}
	var zero AnyMoney
	if err := json.Unmarshal(b, &zero); err != nil || zero != (AnyMoney{}) {
		t.Errorf("Unmarshal() of the zero value = %v, %v", zero, err)
	}
	if _, err := json.Marshal(NewAnyMoney(5, nil)); !errors.Is(err, ErrUnknownCurrency) {
		t.Errorf("Marshal() of an amount without a currency error = %v", err)
	}
	if _, err := NewAnyMoneyFromDecimal("1", nil); !errors.Is(err, ErrUnknownCurrency) {
		t.Errorf("NewAnyMoneyFromDecimal() with nil currency error = %v", err)
	}
}
