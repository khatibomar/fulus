# ICU golden data

`golden.tsv` holds currency amounts that ICU4J formats.
ICU is an independent implementation of CLDR.
`TestFormatMatchesICU` compares `format.AnyMoney` and `format.ParseAny` with this data.

The file has these columns: locale, currency code, amount in minor units, formatted amount.
The first line gives the ICU and CLDR versions.

To update the file, run this command from the repository root. It downloads the ICU4J release and uses it to format the amounts:

```sh
./format/testdata/icu/generate.sh > format/testdata/icu/golden.tsv
```

Use the ICU release that has the same CLDR version as `locale/gen_locale.go`, because a different ICU release formats some locales differently.
ICU 79.1 has CLDR 49.0.0. Set `ICU_VERSION` to use another release, or set `ICU4J_JAR` to use a jar that you already have.
Node.js is not used because its bundled ICU version can be older than the CLDR data.

The script skips the locales that ICU does not have.
`icuDifferences` in `format/icu_test.go` lists the cases where ICU does not agree with the CLDR data, and the reasons.

The data comes from CLDR through ICU. Both use the Unicode License v3. See `LICENSE-UNICODE` in the repository root.
