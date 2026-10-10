# Security policy

## Supported versions

Only the latest release gets security fixes.
Before v1.0.0, a fix can come with a breaking change. The release notes tell you about it.

## Report a vulnerability

Do not open a public issue for a vulnerability.

Use the private report form of GitHub:

1. Go to the [Security tab](https://github.com/khatibomar/fulus/security) of the repository.
2. Click **Report a vulnerability**.
3. Tell us the version, the steps to reproduce the problem, and the effect.

We reply in 7 days. When the fix is ready, we publish a release and a GitHub security advisory.

## Scope

These problems are in scope:

- An operation that returns a wrong amount and no error, for example after an overflow.
- A parse function that accepts an input and gives a different amount.
- A panic or an unbounded allocation from untrusted input, for example in `UnmarshalJSON`, `Scan` or `ParseFormatted`.

A difference between `Format` and CLDR data is a bug, not a vulnerability. Open a normal issue for it.
