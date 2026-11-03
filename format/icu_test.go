package format

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/khatibomar/fulus"
	"github.com/khatibomar/fulus/currency"
	"github.com/khatibomar/fulus/locale"
)

// icuDifferences maps a locale, or a locale and currency pair, where ICU does not agree with CLDR to the reason.
var icuDifferences = map[string]string{
	"en-BE/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-CH/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-CZ/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-DE/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-DK/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-EE/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-ES/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-FI/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-FR/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-GE/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-HU/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-IT/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-LT/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-LV/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-NL/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-NO/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-PL/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-PT/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-RO/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-SE/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-SI/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-SK/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
	"en-UA/EUR": "ICU uses the separators of en-150 with the EUR pattern of en-150",
}

type icuCase struct {
	line     int
	locale   string
	code     string
	amount   int64
	expected string
}

func readICUGolden(t *testing.T) []icuCase {
	t.Helper()
	data, err := os.ReadFile("testdata/icu/golden.tsv")
	if err != nil {
		t.Fatal(err)
	}

	var cases []icuCase
	line := 0
	for text := range strings.Lines(string(data)) {
		line++
		text = strings.TrimSuffix(text, "\n")
		if strings.HasPrefix(text, "#") {
			continue
		}
		fields := strings.Split(text, "\t")
		if len(fields) != 4 {
			t.Fatalf("line %d: want 4 fields, got %d", line, len(fields))
		}
		amount, err := strconv.ParseInt(fields[2], 10, 64)
		if err != nil {
			t.Fatalf("line %d: %v", line, err)
		}
		cases = append(cases, icuCase{line: line, locale: fields[0], code: fields[1], amount: amount, expected: fields[3]})
	}
	return cases
}

func (c icuCase) skipReason() string {
	if reason, ok := icuDifferences[c.locale]; ok {
		return reason
	}
	if reason, ok := icuDifferences[c.locale+"/"+c.code]; ok {
		return reason
	}
	return ""
}

// TestFormatMatchesICU compares AnyMoney and ParseAny with ICU, an independent CLDR implementation. See testdata/icu/README.md.
func TestFormatMatchesICU(t *testing.T) {
	t.Parallel()

	cases := readICUGolden(t)
	if len(cases) < 20000 {
		t.Fatalf("golden file has only %d cases", len(cases))
	}

	checked, failures := 0, 0
	for _, c := range cases {
		if c.skipReason() != "" {
			continue
		}
		loc, ok := locale.ParseLocale(c.locale)
		if !ok {
			t.Fatalf("line %d: unknown locale %q", c.line, c.locale)
		}
		cur, ok := currency.ByCode(c.code)
		if !ok {
			t.Fatalf("line %d: unknown currency %q", c.line, c.code)
		}
		checked++

		m := fulus.NewAnyMoney(c.amount, cur)
		if got := AnyMoney(m, loc); got != c.expected {
			t.Errorf("line %d: AnyMoney(%s, %s, %d) = %q, ICU gives %q", c.line, c.locale, c.code, c.amount, got, c.expected)
			failures++
		}
		if got, err := ParseAny(c.expected, cur, loc); err != nil || got.Decimal() != m.Decimal() {
			t.Errorf("line %d: ParseAny(%q, %s) = %s, %v, want %s", c.line, c.expected, c.locale, got.Decimal(), err, m.Decimal())
			failures++
		}
		if failures >= 20 {
			t.Fatal("too many failures")
		}
	}
	t.Logf("checked %d of %d cases", checked, len(cases))
}
