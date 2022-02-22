# Numeral Systems

> Warning: It is not allowed to use any external long maths / big numbers libraries.

## Task 1: Signed Long Multiplication Base 2 (100%)

### Definition

Write a program (with tests) that implements "long" signed arithmetics multiplication.
The program may contain a `LongMul` function, which accepts two binary numbers encoded
as strings. Allowed sign literals are `-` or `+` or none. None defaults to `+`. See the
examples below for specs.

```go
func LongMul(a, b string) string { ... }
```

### Test examples

```go
// Checking signed operations.
{"10", "101", "1010"},
{"+10", "101", "1010"},
{"10", "+101", "1010"},
{"+10", "+101", "1010"},
{"-10", "101", "-1010"},
{"10", "-101", "-1010"},
{"-10", "-101", "1010"},
{"+10", "-101", "-1010"},
{"-10", "+101", "-1010"},
// Resulting zero must never contain sign literals.
{"0", "0", "0"},
{"-0", "0", "0"},
{"0", "-0", "0"},
{"1", "0", "0"},
{"1", "+0", "0"},
{"1", "-0", "0"},
{"+1", "-0", "0"},
{"-1", "+0", "0"},

// Other examples.
{"1101100100000011", "110111", "1011101001111110100101"},
```

## Extra Task: Benchmarks (+30%)

* Read about benchmarks (e.g. https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go).
* Implement benchmarks for your programs.
* Try to optimize your program.
* Try to make any meaningful conclusions.
* Have fun!
