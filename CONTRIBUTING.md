# Contributing

Thank you for your help. This page tells you how to make a change.

## Before you start

For a large change or a new API, open an issue first. Then we can agree on the design before you write the code.
Read [Design decisions](docs/design-decisions.md) to see why Fulus works as it does.

## Development

You need the Go version in `go.mod`.

```sh
go vet ./...
go test -race ./...
golangci-lint run ./...
make fuzz FUZZ_TIME=10s
(cd examples && go test ./...)
```

## Generated files

Do not edit `locale/gen_locale.go` or `currency/gen_currencies.go`. `generator.go` writes them from CLDR and ISO 4217 data.

```sh
make gen        # use the latest CLDR release
make check-gen  # generate again with the current CLDR version and show changes
```

When you change the generator or the formatting code, update the ICU golden data and run the tests.
You need Node.js, because it includes ICU:

```sh
node format/testdata/icu/generate.mjs > format/testdata/icu/golden.tsv
go test -run TestFormatMatchesICU ./format
```

If ICU does not agree with the CLDR data for a good reason, add the case to `icuDifferences` in `format/icu_test.go` with the reason.

### Minor units

`currency/testdata/minor_units.golden` holds the minor units of each generated currency.
`TestMinorUnitsDoNotChange` fails when the generated data does not agree with it.

- For a new currency, run `go test ./currency -run TestMinorUnitsDoNotChange -update`.
- A changed or removed line is a breaking change: an amount in minor units that is stored in a database or a queue
  changes its value. Tell this in the pull request, so that the release notes list it.
  CI fails for such a change after a stable release.

## Tests

- Use table tests.
- Call `t.Parallel()` in each test that can run in parallel.
- Do not test the standard library.
- Add an `Example` function for a new exported function.
- For arithmetic, extend `checkMoneyOps` in `differential_test.go`. It compares each operation and each rounding mode
  with an independent `math/big` reference. `FuzzDifferential` runs it with random values.

## Documentation and comments

Write documentation and comments in [ASD-STE100 Simplified Technical English](https://www.asd-ste100.org/):
short sentences, active voice, and one meaning for each word.
A comment tells what the code does now. It does not tell the history of the code.

## Commits and pull requests

- Write the commit subject in the imperative mood, for example "Add RoundCash".
- In the commit body, tell what changed and why.
- In the pull request, tell which changes a user can see, so that the release notes can list them.
- Keep a pull request to one subject.

## API changes

Read the [stability policy](README.md#stability). CI compares the API with the latest release.
A change that removes or changes an exported name needs a deprecation first, after v1.0.0.
