package currency

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

type testCurrency struct {
	code, number string
}

func (c testCurrency) Code() string    { return c.code }
func (c testCurrency) Number() string  { return c.number }
func (c testCurrency) Name() string    { return "Test " + c.code }
func (c testCurrency) MinorUnits() int { return 2 }

func TestByCodeAndNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		lookup func() (Currency, bool)
		want   string
		ok     bool
	}{
		{name: "code", lookup: func() (Currency, bool) { return ByCode("USD") }, want: "USD", ok: true},
		{name: "lower case code", lookup: func() (Currency, bool) { return ByCode("chf") }, want: "CHF", ok: true},
		{name: "unknown code", lookup: func() (Currency, bool) { return ByCode("XYZ") }},
		{name: "number", lookup: func() (Currency, bool) { return ByNumber("978") }, want: "EUR", ok: true},
		{name: "unknown number", lookup: func() (Currency, bool) { return ByNumber("000") }},
		{name: "gold", lookup: func() (Currency, bool) { return ByCode("XAU") }, want: "XAU", ok: true},
		{name: "silver number", lookup: func() (Currency, bool) { return ByNumber("961") }, want: "XAG", ok: true},
		{name: "testing code is excluded", lookup: func() (Currency, bool) { return ByCode("XTS") }},
		{name: "no currency code is excluded", lookup: func() (Currency, bool) { return ByCode("XXX") }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c, ok := tt.lookup()
			if ok != tt.ok {
				t.Fatalf("ok = %v, want %v", ok, tt.ok)
			}
			if ok && c.Code() != tt.want {
				t.Errorf("code = %s, want %s", c.Code(), tt.want)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		currency Currency
		wantErr  error
	}{
		{name: "new currency", currency: testCurrency{code: "ZZTEST", number: "9001"}},
		{name: "new currency without number", currency: testCurrency{code: "ZZNONUM"}},
		{name: "second currency without number", currency: testCurrency{code: "ZZNONUM2"}},
		{name: "duplicate code", currency: testCurrency{code: "USD", number: "9002"}, wantErr: ErrDuplicateCurrency},
		{name: "duplicate number", currency: testCurrency{code: "ZZDUP", number: "840"}, wantErr: ErrDuplicateCurrency},
		{name: "empty code", currency: testCurrency{}, wantErr: ErrInvalidCurrency},
		{name: "lower case code", currency: testCurrency{code: "zzlow"}, wantErr: ErrInvalidCurrency},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := Register(tt.currency)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Register() error = %v, want %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}
			if c, ok := ByCode(tt.currency.Code()); !ok || c != tt.currency {
				t.Errorf("ByCode() = %v, %v after Register", c, ok)
			}
		})
	}
}

func TestAll(t *testing.T) {
	t.Parallel()

	all := All()
	if len(all) < len(builtin) {
		t.Fatalf("All() has %d currencies, want at least %d", len(all), len(builtin))
	}
	if !slices.IsSortedFunc(all, func(a, b Currency) int { return strings.Compare(a.Code(), b.Code()) }) {
		t.Error("All() is not sorted by code")
	}
}

func TestBuiltinIsUnique(t *testing.T) {
	t.Parallel()

	codes := make(map[string]bool)
	numbers := make(map[string]bool)
	for _, c := range builtin {
		if codes[c.Code()] || numbers[Number(c)] {
			t.Errorf("duplicate builtin currency %s %s", c.Code(), Number(c))
		}
		codes[c.Code()] = true
		numbers[Number(c)] = true
	}
}

func TestMinorUnitsWithoutISOValue(t *testing.T) {
	t.Parallel()

	// ISO 4217 defines no minor units for these codes. The generated value comes from CLDR.
	for _, c := range []Currency{XAU{}, XAG{}, XPT{}, XPD{}, XDR{}, XSU{}, XUA{}} {
		if got := c.MinorUnits(); got != 2 {
			t.Errorf("%s.MinorUnits() = %d, want 2", c.Code(), got)
		}
	}
}

type codeOnly struct{}

func (codeOnly) Code() string    { return "CODEONLY" }
func (codeOnly) MinorUnits() int { return 8 }

func TestRegisterWithoutNumber(t *testing.T) {
	t.Parallel()

	if err := Register(codeOnly{}); err != nil {
		t.Fatal(err)
	}
	if c, ok := ByCode("CODEONLY"); !ok || c.MinorUnits() != 8 {
		t.Errorf("ByCode() = %v, %v", c, ok)
	}
	if got := Number(codeOnly{}); got != "" {
		t.Errorf("Number() = %q, want empty", got)
	}
}

func TestRegistryInstance(t *testing.T) {
	t.Parallel()

	r, err := NewRegistry(USD{}, EUR{})
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Register(testCurrency{code: "LOCALONLY", number: "991"}); err != nil {
		t.Fatal(err)
	}
	if _, ok := r.ByCode("localonly"); !ok {
		t.Error("ByCode() does not find the registered currency")
	}
	if _, ok := ByCode("LOCALONLY"); ok {
		t.Error("the Default registry finds a currency of another registry")
	}
	if _, ok := r.ByCode("JPY"); ok {
		t.Error("ByCode() finds a currency that is not in the registry")
	}
	if c, ok := r.ByNumber("978"); !ok || c.Code() != "EUR" {
		t.Errorf("ByNumber(978) = %v, %v", c, ok)
	}
	if got := len(r.All()); got != 3 {
		t.Errorf("len(All()) = %d, want 3", got)
	}

	if _, err := NewRegistry(USD{}, USD{}); !errors.Is(err, ErrDuplicateCurrency) {
		t.Errorf("NewRegistry() with a duplicate error = %v", err)
	}

	var zero Registry
	if err := zero.Register(USD{}); err != nil {
		t.Errorf("zero Registry Register() error = %v", err)
	}
	if err := zero.Register(nil); !errors.Is(err, ErrInvalidCurrency) {
		t.Errorf("Register(nil) error = %v", err)
	}

	all, err := NewRegistry(Builtin()...)
	if err != nil || len(all.All()) != len(builtin) {
		t.Errorf("NewRegistry(Builtin()...) = %d currencies, %v", len(all.All()), err)
	}
}

type negativeUnits struct{}

func (negativeUnits) Code() string    { return "NEG" }
func (negativeUnits) MinorUnits() int { return -1 }

func TestRegisterNegativeMinorUnits(t *testing.T) {
	t.Parallel()

	var r Registry
	if err := r.Register(negativeUnits{}); !errors.Is(err, ErrInvalidCurrency) {
		t.Errorf("Register() error = %v", err)
	}
}
