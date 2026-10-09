package locale

import (
	"slices"
	"strings"
)

// Match returns the supported locale that best fits a BCP 47 or POSIX tag,
// for example "en_US.UTF-8", "de-ch" or "fr-CA-u-nu-latn".
// It removes subtags from the end of the tag until it finds a supported locale,
// so "en-US" gives EN and "fr-CA-x-private" gives FR_CA.
// It does not use CLDR likely subtags, so "zh-TW" gives ZH and not ZH_HANT.
func Match(tag string) (Locale, bool) {
	tag, _, _ = strings.Cut(tag, ".")
	tag, _, _ = strings.Cut(tag, "@")
	subtags := strings.Split(strings.ReplaceAll(tag, "_", "-"), "-")

	// A singleton such as "u" or "x" starts an extension, which has no effect on the locale.
	if i := slices.IndexFunc(subtags, func(s string) bool { return len(s) == 1 }); i >= 0 {
		subtags = subtags[:i]
	}

	for i, subtag := range subtags {
		switch {
		case i == 0:
			subtags[i] = strings.ToLower(subtag)
		case len(subtag) == 4 && isAlpha(subtag):
			subtags[i] = strings.ToUpper(subtag[:1]) + strings.ToLower(subtag[1:])
		case len(subtag) == 2 || len(subtag) == 3 && !isAlpha(subtag):
			subtags[i] = strings.ToUpper(subtag)
		default:
			subtags[i] = strings.ToLower(subtag)
		}
	}

	for n := len(subtags); n > 0; n-- {
		if l, ok := ParseLocale(strings.Join(subtags[:n], "-")); ok {
			return l, true
		}
	}
	return Locale{}, false
}

func isAlpha(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}
