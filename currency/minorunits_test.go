package currency

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

var update = flag.Bool("update", false, "write testdata/minor_units.golden from the generated currencies")

const minorUnitsGolden = "testdata/minor_units.golden"

// TestMinorUnitsDoNotChange compares the minor units of the generated currencies with a golden file.
// Stored amounts in minor units change their value if the minor units change,
// so a change of the golden file is a breaking change. See CONTRIBUTING.md.
func TestMinorUnitsDoNotChange(t *testing.T) {
	if *update {
		var b strings.Builder
		for _, c := range builtin {
			fmt.Fprintf(&b, "%s %d\n", c.Code(), c.MinorUnits())
		}
		if err := os.WriteFile(minorUnitsGolden, []byte(b.String()), 0o644); err != nil {
			t.Fatal(err)
		}
		return
	}

	data, err := os.ReadFile(minorUnitsGolden)
	if err != nil {
		t.Fatal(err)
	}

	golden := make(map[string]int)
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		code, digits, ok := strings.Cut(scanner.Text(), " ")
		n, err := strconv.Atoi(digits)
		if !ok || err != nil {
			t.Fatalf("invalid line %q", scanner.Text())
		}
		golden[code] = n
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}

	for _, c := range builtin {
		want, ok := golden[c.Code()]
		switch {
		case !ok:
			t.Errorf("%s is new. Run go test ./currency -run TestMinorUnitsDoNotChange -update", c.Code())
		case c.MinorUnits() != want:
			t.Errorf("%s minor units changed from %d to %d. This is a breaking change", c.Code(), want, c.MinorUnits())
		}
		delete(golden, c.Code())
	}
	for code := range golden {
		t.Errorf("%s was removed. This is a breaking change", code)
	}
}
