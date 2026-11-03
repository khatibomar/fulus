#!/bin/sh
# generate.sh writes golden.tsv with amounts formatted by ICU4J.
# Run it from the repository root:
#
#   ./format/testdata/icu/generate.sh > format/testdata/icu/golden.tsv
#
# Set ICU_VERSION to the ICU release that has the same CLDR version as locale/gen_locale.go.
# ICU 79.1 has CLDR 49.0.0. Set ICU4J_JAR to use an existing icu4j jar instead of downloading it.
set -eu

script_dir=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
cd "$script_dir/../../.."

version=${ICU_VERSION:-79.1}
jar=${ICU4J_JAR:-}

if [ -z "$jar" ]; then
	tmp_dir=$(mktemp -d)
	trap 'rm -rf "$tmp_dir"' EXIT
	jar=$tmp_dir/icu4j-$version.jar
	curl -fsSL "https://github.com/unicode-org/icu/releases/download/release-$version/icu4j-$version.jar" -o "$jar"
fi

# The source-file mode needs the compiler module, which some Java runtimes do not include by default.
java --add-modules jdk.compiler -cp "$jar" "$script_dir/Generate.java"
