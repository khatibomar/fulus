CLDR_VERSION = $(shell sed -n 's|^// CLDR Version: ||p' locale/gen_locale.go)
FUZZ_TIME ?= 10s
COVER_MIN ?= 95

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

# cover fails if the statement coverage of the fulus and format packages is below COVER_MIN percent.
cover:
	go test -coverprofile=cover.out . ./format
	go tool cover -func=cover.out | awk -v min=$(COVER_MIN) '/^total:/ { sub("%", "", $$3); print "coverage " $$3 "%"; if ($$3 + 0 < min) { print "coverage is below " min "%"; exit 1 } }'

clean:
	find . -name "gen_*.go" -type f -delete

.PHONY: gen check-gen fuzz cover clean
