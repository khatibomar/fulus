// Generate.java writes golden.tsv with currency amounts formatted by ICU4J.
// Run it with generate.sh from anywhere in the repository:
//
//   ./format/testdata/icu/generate.sh > format/testdata/icu/golden.tsv
//
// ICU4J is the reference implementation of CLDR. Use the ICU release that has the same
// CLDR version as locale/gen_locale.go. ICU 79.1 has CLDR 49.0.0.
import com.ibm.icu.number.LocalizedNumberFormatter;
import com.ibm.icu.number.NumberFormatter;
import com.ibm.icu.number.Precision;
import com.ibm.icu.text.NumberFormat;
import com.ibm.icu.util.Currency;
import com.ibm.icu.util.ULocale;
import com.ibm.icu.util.VersionInfo;

import java.io.IOException;
import java.math.BigDecimal;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Generate {
    private static final String LOCALE_PATTERN = "(?m)^\\t\\{code: \"([^\"]+)\"";
    private static final String MINOR_UNITS_PATTERN =
            "(?m)^func \\(([A-Z]{3})\\) MinorUnits\\(\\) int \\{ return (\\d+) \\}";

    // Every locale formats these currencies. They cover symbols, codes, 0, 2, 3 and 4 minor units and CLDR overrides.
    private static final List<String> SAMPLE = List.of(
            "AED", "BHD", "CHF", "CLF", "CVE", "EUR", "GBP", "INR", "JPY", "TRY", "USD");
    // These locales format all currencies.
    private static final Set<String> ALL_CURRENCIES = Set.of(
            "ar", "de", "en", "es", "fr", "hi", "ja", "pt", "ru", "zh");
    // The amounts are in minor units. They cover the minus sign, digit grouping and fraction padding.
    private static final long[] AMOUNTS = {5L, 100000L, -123456789L};

    public static void main(String[] args) throws IOException {
        String localeSource = Files.readString(Path.of("locale/gen_locale.go"));
        String currencySource = Files.readString(Path.of("currency/gen_currencies.go"));

        List<String> locales = matches(LOCALE_PATTERN, localeSource);
        Map<String, Integer> minorUnits = new LinkedHashMap<>();
        Matcher currencyMatcher = Pattern.compile(MINOR_UNITS_PATTERN).matcher(currencySource);
        while (currencyMatcher.find()) {
            minorUnits.put(currencyMatcher.group(1), Integer.parseInt(currencyMatcher.group(2)));
        }

        String cldrVersion = matches("(?m)^// CLDR Version: (.+)$", localeSource).get(0);
        int cldrMajor = Integer.parseInt(cldrVersion.split("\\.", 2)[0]);
        if (VersionInfo.ICU_VERSION.getMajor() != cldrMajor + 30) {
            throw new IllegalStateException("ICU " + VersionInfo.ICU_VERSION + " does not match CLDR " + cldrVersion
                    + ": expected an ICU " + (cldrMajor + 30) + " release");
        }
        StringBuilder out = new StringBuilder();
        out.append("# ICU ").append(VersionInfo.ICU_VERSION).append(", CLDR ").append(cldrVersion)
                .append(". Columns: locale, currency, amount in minor units, formatted amount.\n");

        for (String locale : locales) {
            ULocale ulocale = ULocale.forLanguageTag(locale);
            ULocale resolved = NumberFormat.getCurrencyInstance(ulocale).getLocale(ULocale.VALID_LOCALE);
            // Skip the locales that ICU does not have, because ICU uses the data of another locale for them.
            if (!resolved.toLanguageTag().equals(ulocale.toLanguageTag())) {
                continue;
            }

            List<String> codes = ALL_CURRENCIES.contains(locale)
                    ? new ArrayList<>(minorUnits.keySet())
                    : SAMPLE;
            for (String code : codes) {
                int digits = minorUnits.get(code);
                LocalizedNumberFormatter format = NumberFormatter.withLocale(ULocale.forLanguageTag(locale + "-u-nu-latn"))
                        .unit(Currency.getInstance(code))
                        .precision(Precision.fixedFraction(digits));
                for (long amount : AMOUNTS) {
                    out.append(locale).append('\t').append(code).append('\t').append(amount).append('\t')
                            .append(format.format(new BigDecimal(decimal(amount, digits)))).append('\n');
                }
            }
        }
        System.out.print(out);
    }

    private static List<String> matches(String pattern, String source) {
        List<String> result = new ArrayList<>();
        Matcher matcher = Pattern.compile(pattern).matcher(source);
        while (matcher.find()) {
            result.add(matcher.group(1));
        }
        return result;
    }

    private static String decimal(long amount, int digits) {
        String sign = amount < 0 ? "-" : "";
        String s = Long.toString(Math.abs(amount));
        while (s.length() < digits + 1) {
            s = "0" + s;
        }
        return digits == 0 ? sign + s : sign + s.substring(0, s.length() - digits) + "."
                + s.substring(s.length() - digits);
    }
}
