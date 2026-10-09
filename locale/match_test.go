package locale

import "testing"

func TestMatch(t *testing.T) {
	t.Parallel()

	tests := []struct {
		tag  string
		want Locale
		ok   bool
	}{
		{tag: "en", want: EN, ok: true},
		{tag: "en-US", want: EN, ok: true},
		{tag: "en_US.UTF-8", want: EN, ok: true},
		{tag: "en_IN@euro", want: EN_IN, ok: true},
		{tag: "DE-ch", want: DE_CH, ok: true},
		{tag: "fr-CA-u-nu-latn", want: FR_CA, ok: true},
		{tag: "fr-CA-x-private", want: FR_CA, ok: true},
		{tag: "zh-hant-hk", want: ZH_HANT_HK, ok: true},
		{tag: "sr-Latn-XX", want: SR_LATN, ok: true},
		{tag: "es-419", want: ES_419, ok: true},
		{tag: "en-001", want: EN_001, ok: true},
		{tag: "zh-TW", want: ZH, ok: true},
		{tag: "xx-YY", ok: false},
		{tag: "", ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.tag, func(t *testing.T) {
			t.Parallel()

			got, ok := Match(tt.tag)
			if ok != tt.ok || got != tt.want {
				t.Errorf("Match(%q) = %v, %v; want %v, %v", tt.tag, got, ok, tt.want, tt.ok)
			}
		})
	}
}
