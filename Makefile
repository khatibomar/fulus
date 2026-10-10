CLDR_VERSION = $(shell sed -n 's|^// CLDR Version: ||p' locale/gen_locale.go)
FUZZ_TIME ?= 10s

gen:
	go run generator.go

# check-gen generates the data again with the same CLDR version and fails if a generated file changes.
check-gen:
	go run generator.go -cldr $(CLDR_VERSION)
	git diff --exit-code -I '^// ISO 4217 Data Last Updated:' -- locale/gen_locale.go currency/gen_currencies.go

fuzz:
	for pkg in $$(go list ./...); do \
		for target in $$(go test -list '^Fuzz' $$pkg | grep '^Fuzz'); do \
			go test -run '^$$' -fuzz "^$$target$$" -fuzztime $(FUZZ_TIME) $$pkg || exit 1; \
		done; \
	done

clean:
	find . -name "gen_*.go" -type f -delete

.PHONY: gen check-gen fuzz clean
