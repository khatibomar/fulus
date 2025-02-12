# Money Experiment

A Go snippet for handling monetary values and currency operations with type safety.

## Overview

This snippet provides a robust solution for handling money and currency operations in Go applications. It implements type-safe money operations with configurable base currency per application instance.

it's not a library really because if I need it to be so, I need to export Amount so money_gen gets generated on client side, but I think it's unsafe to do so!

for now it's just a snippet, I don't know how far I can go with this.

## Features

- Type-safe money operations
- Configurable base currency per application instance
- Currency conversion support
- Prevention of invalid currency operations
- Precise decimal handling
