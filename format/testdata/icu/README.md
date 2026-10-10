# ICU golden data

`golden.tsv` holds currency amounts that ICU formats.
ICU is an independent implementation of CLDR.
`TestFormatMatchesICU` compares `format.AnyMoney` and `format.ParseAny` with this data.

The file has these columns: locale, currency code, amount in minor units, formatted amount.
The first line gives the ICU and CLDR versions.

To update the file, run this command from the repository root with Node.js, which includes ICU:

```sh
node format/testdata/icu/generate.mjs > format/testdata/icu/golden.tsv
```

The script skips the locales that ICU does not have.
`icuDifferences` in `format/icu_test.go` lists the cases where ICU does not agree with the CLDR data, and the reasons.

The data comes from CLDR through ICU. Both use the Unicode License v3. See `LICENSE-UNICODE` in the repository root.
