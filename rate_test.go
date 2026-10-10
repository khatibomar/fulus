package fulus

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/khatibomar/fulus/currency"
)

func TestNewRate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		num, den int64
		wantNum  int64
		wantDen  int64
		wantErr  error
	}{
		{name: "reduces", num: 10, den: 4, wantNum: 5, wantDen: 2},
		{name: "zero numerator", num: 0, den: 1, wantErr: ErrInvalidExchangeRate},
		{name: "zero denominator", num: 1, den: 0, wantErr: ErrInvalidExchangeRate},
		{name: "negative", num: -1, den: 2, wantErr: ErrInvalidExchangeRate},
		{name: "both negative", num: -1, den: -2, wantErr: ErrInvalidExchangeRate},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r, err := NewRate[currency.EUR, currency.USD](tt.num, tt.den)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("NewRate() error = %v, expected %v", err, tt.wantErr)
			}
			if n, d := r.Fraction(); err == nil && (n != tt.wantNum || d != tt.wantDen) {
				t.Errorf("Fraction() = %d/%d, expected %d/%d", n, d, tt.wantNum, tt.wantDen)
			}
		})
	}
}

func TestParseRate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		rate    string
		wantNum int64
		wantDen int64
		wantErr error
	}{
		{name: "integer", rate: "2", wantNum: 2, wantDen: 1},
		{name: "decimal", rate: "1.5", wantNum: 3, wantDen: 2},
		{name: "market quote", rate: "1.07203", wantNum: 107203, wantDen: 100000},
		{name: "small", rate: "0.00001", wantNum: 1, wantDen: 100000},
		{name: "fraction", rate: "2/6", wantNum: 1, wantDen: 3},
		{name: "negative", rate: "-1.5", wantErr: ErrInvalidExchangeRate},
		{name: "zero", rate: "0", wantErr: ErrInvalidExchangeRate},
		{name: "invalid", rate: "abc", wantErr: ErrInvalidExchangeRate},
		{name: "too many digits", rate: "0.0000000000000000000001", wantErr: ErrOverflow},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r, err := ParseRate[currency.EUR, currency.USD](tt.rate)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ParseRate() error = %v, expected %v", err, tt.wantErr)
			}
			if n, d := r.Fraction(); err == nil && (n != tt.wantNum || d != tt.wantDen) {
				t.Errorf("Fraction() = %d/%d, expected %d/%d", n, d, tt.wantNum, tt.wantDen)
			}
		})
	}
}

func TestRateFromFloat64(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		rate    float64
		wantNum int64
		wantDen int64
		wantErr error
	}{
		{name: "decimal", rate: 1.5, wantNum: 3, wantDen: 2},
		{name: "shortest decimal", rate: 0.1, wantNum: 1, wantDen: 10},
		{name: "market quote", rate: 1.07203, wantNum: 107203, wantDen: 100000},
		{name: "negative", rate: -1.5, wantErr: ErrInvalidExchangeRate},
		{name: "zero", rate: 0, wantErr: ErrInvalidExchangeRate},
		{name: "NaN", rate: math.NaN(), wantErr: ErrInvalidExchangeRate},
		{name: "infinity", rate: math.Inf(1), wantErr: ErrInvalidExchangeRate},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r, err := RateFromFloat64[currency.EUR, currency.USD](tt.rate)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("RateFromFloat64() error = %v, expected %v", err, tt.wantErr)
			}
			if n, d := r.Fraction(); err == nil && (n != tt.wantNum || d != tt.wantDen) {
				t.Errorf("Fraction() = %d/%d, expected %d/%d", n, d, tt.wantNum, tt.wantDen)
			}
		})
	}
}

func TestRateString(t *testing.T) {
	t.Parallel()

	tests := map[string]string{
		"1.07203": "1.07203",
		"160":     "160",
		"1/4":     "0.25",
		"1/40":    "0.025",
		"1/3":     "1/3",
		"7/12":    "7/12",
	}
	for in, want := range tests {
		if got := MustParseRate[currency.EUR, currency.USD](in).String(); got != want {
			t.Errorf("ParseRate(%q).String() = %q, expected %q", in, got, want)
		}
	}
	if got := (Rate[currency.EUR, currency.USD]{}).String(); got != "0" {
		t.Errorf("zero Rate String() = %q, expected 0", got)
	}
}

func TestRateText(t *testing.T) {
	t.Parallel()

	type config struct {
		Rate Rate[currency.EUR, currency.USD] `json:"rate"`
	}
	var c config
	if err := json.Unmarshal([]byte(`{"rate":"1.0720"}`), &c); err != nil {
		t.Fatal(err)
	}
	out, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != `{"rate":"1.072"}` {
		t.Errorf("Marshal() = %s", out)
	}
	if err := json.Unmarshal([]byte(`{"rate":"-1"}`), &c); !errors.Is(err, ErrInvalidExchangeRate) {
		t.Errorf("Unmarshal() negative rate error = %v", err)
	}
	if _, err := json.Marshal(config{}); !errors.Is(err, ErrInvalidExchangeRate) {
		t.Errorf("Marshal() zero rate error = %v", err)
	}
}

func TestInvertAndCross(t *testing.T) {
	t.Parallel()

	eurUSD := MustParseRate[currency.EUR, currency.USD]("1.25")
	usdJPY := MustParseRate[currency.USD, currency.JPY]("150")

	usdEUR := eurUSD.Invert()
	if n, d := usdEUR.Fraction(); n != 4 || d != 5 {
		t.Errorf("Invert() = %d/%d, expected 4/5", n, d)
	}

	eurJPY, err := Cross(eurUSD, usdJPY)
	if err != nil {
		t.Fatal(err)
	}
	if n, d := eurJPY.Fraction(); n != 375 || d != 2 {
		t.Errorf("Cross() = %d/%d, expected 375/2", n, d)
	}

	if _, err := Cross(Rate[currency.EUR, currency.USD]{}, usdJPY); !errors.Is(err, ErrInvalidExchangeRate) {
		t.Errorf("Cross() with zero rate error = %v", err)
	}

	big1, _ := NewRate[currency.EUR, currency.USD](math.MaxInt64, 1)
	big2, _ := NewRate[currency.USD, currency.JPY](3, 1)
	if _, err := Cross(big1, big2); !errors.Is(err, ErrOverflow) {
		t.Errorf("Cross() overflow error = %v", err)
	}

	// Cross cancels common factors before it multiplies.
	r1, _ := NewRate[currency.EUR, currency.USD](math.MaxInt64, 3)
	r2, _ := NewRate[currency.USD, currency.JPY](3, math.MaxInt64)
	same, err := Cross(r1, r2)
	if n, d := same.Fraction(); err != nil || n != 1 || d != 1 {
		t.Errorf("Cross() = %d/%d, %v, expected 1/1", n, d, err)
	}
}

func TestConvert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		amount   int64
		rate     string
		mode     RoundingMode
		expected int64
		wantErr  error
	}{
		{name: "simple", amount: 10000, rate: "1.07203", mode: RoundTruncate, expected: 10720},
		{name: "zero amount", amount: 0, rate: "1.07203", mode: RoundTruncate},
		{name: "overflow", amount: math.MaxInt64, rate: "2", mode: RoundTruncate, wantErr: ErrOverflow},
		{name: "truncate positive half", amount: 1, rate: "1/2", mode: RoundTruncate, expected: 0},
		{name: "half up positive tie", amount: 5, rate: "1/2", mode: RoundHalfUp, expected: 3},
		{name: "half even positive tie", amount: 5, rate: "1/2", mode: RoundHalfEven, expected: 2},
		{name: "half up negative tie", amount: -1, rate: "1/2", mode: RoundHalfUp, expected: -1},
		{name: "half even negative tie", amount: -1, rate: "1/2", mode: RoundHalfEven, expected: 0},
		{name: "unnecessary inexact", amount: 1, rate: "1/2", mode: RoundUnnecessary, wantErr: ErrInexact},
		{name: "invalid mode", amount: 1, rate: "1/2", mode: RoundingMode(99), wantErr: ErrInvalidRoundingMode},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := MustParseRate[currency.EUR, currency.USD](tt.rate)
			got, err := Convert(NewMoney[currency.EUR](tt.amount), r, tt.mode)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Convert() error = %v, expected %v", err, tt.wantErr)
			}
			if err == nil && got.Amount() != tt.expected {
				t.Errorf("Convert() = %d, expected %d", got.Amount(), tt.expected)
			}
		})
	}

	if _, err := Convert(NewMoney[currency.EUR](1), Rate[currency.EUR, currency.USD]{}, RoundHalfEven); !errors.Is(err, ErrInvalidExchangeRate) {
		t.Errorf("Convert() with zero rate error = %v", err)
	}
}

func TestConvertDifferentMinorUnits(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		got  func() (int64, error)
		want int64
	}{
		{
			name: "EUR 1.00 to JPY at 160.25",
			got: func() (int64, error) {
				m, err := Convert(NewMoney[currency.EUR](100), MustParseRate[currency.EUR, currency.JPY]("160.25"), RoundHalfEven)
				return m.Amount(), err
			},
			want: 160,
		},
		{
			name: "USD 1.00 to BHD at 0.376",
			got: func() (int64, error) {
				m, err := Convert(NewMoney[currency.USD](100), MustParseRate[currency.USD, currency.BHD]("0.376"), RoundHalfEven)
				return m.Amount(), err
			},
			want: 376,
		},
		{
			name: "JPY 1000 to USD at 1/150",
			got: func() (int64, error) {
				m, err := Convert(NewMoney[currency.JPY](1000), MustParseRate[currency.JPY, currency.USD]("1/150"), RoundHalfUp)
				return m.Amount(), err
			},
			want: 667,
		},
		{
			name: "scaled rate does not fit int64",
			got: func() (int64, error) {
				r, _ := NewRate[currency.JPY, currency.CLF](math.MaxInt64, math.MaxInt64-1)
				m, err := Convert(NewMoney[currency.JPY](3), r, RoundHalfUp)
				return m.Amount(), err
			},
			want: 30000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.got()
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("Convert() = %d, expected %d", got, tt.want)
			}
		})
	}
}

func ExampleConvert() {
	eur := NewMoney[currency.EUR](500) // €5.00
	rate, err := ParseRate[currency.EUR, currency.USD]("1.04565")
	if err != nil {
		panic(err)
	}
	usd, err := Convert(eur, rate, RoundHalfEven)
	if err != nil {
		panic(err)
	}
	fmt.Println(usd)
	// Output: $5.23
}

func ExampleCross() {
	eurUSD := MustParseRate[currency.EUR, currency.USD]("1.25")
	usdJPY := MustParseRate[currency.USD, currency.JPY]("150")

	// Cross(usdJPY, eurUSD) does not compile, because the middle currencies do not agree.
	eurJPY, err := Cross(eurUSD, usdJPY)
	if err != nil {
		panic(err)
	}
	fmt.Println(eurJPY, eurJPY.Invert())
	// Output: 187.5 2/375
}

func FuzzConvert(f *testing.F) {
	f.Add(int64(100), int64(1), int64(3))
	f.Add(int64(-100), int64(1), int64(2))

	f.Fuzz(func(t *testing.T, amount, numerator, denominator int64) {
		r, err := NewRate[currency.EUR, currency.JPY](numerator, denominator)
		if err != nil {
			return
		}
		got, err := Convert(NewMoney[currency.EUR](amount), r, RoundHalfEven)

		// EUR has 2 minor units and JPY has 0, so the exact result is amount*n/(d*100).
		num := new(big.Int).Mul(big.NewInt(amount), big.NewInt(numerator))
		den := new(big.Int).Mul(big.NewInt(denominator), big.NewInt(100))
		want, refErr := divideWithRounding(num, den, RoundHalfEven)
		if refErr != nil {
			t.Fatal(refErr)
		}
		if !want.IsInt64() {
			if !errors.Is(err, ErrOverflow) {
				t.Fatalf("Convert() error = %v, expected ErrOverflow", err)
			}
			return
		}
		if err != nil || got.Amount() != want.Int64() {
			t.Fatalf("Convert(%d, %d/%d) = %d, %v; expected %d", amount, numerator, denominator, got.Amount(), err, want)
		}
	})
}
