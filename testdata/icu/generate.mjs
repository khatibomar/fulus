// generate.mjs writes golden.tsv with currency amounts formatted by ICU through the Intl API of Node.js.
// Run it from the repository root: node testdata/icu/generate.mjs > testdata/icu/golden.tsv
import fs from 'node:fs';

const root = new URL('../../', import.meta.url);
const locales = [...fs.readFileSync(new URL('locale/gen_locale.go', root), 'utf8')
  .matchAll(/^\t\{code: "([^"]+)"/gm)].map(m => m[1]);
const minorUnits = new Map([...fs.readFileSync(new URL('currency/gen_currencies.go', root), 'utf8')
  .matchAll(/^func \(([A-Z]{3})\) MinorUnits\(\) int \{ return (\d+) \}/gm)].map(m => [m[1], Number(m[2])]));

// Every locale formats these currencies. They cover symbols, codes, 0, 2, 3 and 4 minor units and CLDR overrides.
const sample = ['AED', 'BHD', 'CHF', 'CLF', 'CVE', 'EUR', 'GBP', 'INR', 'JPY', 'TRY', 'USD'];
// These locales format all currencies.
const allCurrencies = new Set(['ar', 'de', 'en', 'es', 'fr', 'hi', 'ja', 'pt', 'ru', 'zh']);
// The amounts are in minor units. They cover the minus sign, digit grouping and fraction padding.
const amounts = [5n, 100000n, -123456789n];

function decimal(amount, digits) {
  const sign = amount < 0n ? '-' : '';
  const s = (amount < 0n ? -amount : amount).toString().padStart(digits + 1, '0');
  return digits === 0 ? sign + s : `${sign}${s.slice(0, -digits)}.${s.slice(-digits)}`;
}

const lines = [`# ICU ${process.versions.icu}, CLDR ${process.versions.cldr}. Columns: locale, currency, amount in minor units, formatted amount.`];
for (const locale of locales) {
  // Skip the locales that ICU does not have, because ICU uses the data of another locale for them.
  if (new Intl.NumberFormat(locale).resolvedOptions().locale !== locale) {
    continue;
  }
  const codes = allCurrencies.has(locale) ? [...minorUnits.keys()] : sample;
  for (const code of codes) {
    const digits = minorUnits.get(code);
    const format = new Intl.NumberFormat(`${locale}-u-nu-latn`, {
      style: 'currency',
      currency: code,
      minimumFractionDigits: digits,
      maximumFractionDigits: digits,
    });
    for (const amount of amounts) {
      lines.push([locale, code, amount, format.format(decimal(amount, digits))].join('\t'));
    }
  }
}
process.stdout.write(lines.join('\n') + '\n');
