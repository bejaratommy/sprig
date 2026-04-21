# Integer Math Functions

The following math functions operate on `int64` values.

## add

Sum numbers with `add`. Accepts two or more inputs.

```
add 1 2 3
```

## add1

To increment by 1, use `add1`

## sub

To subtract, use `sub`

## div

Perform integer division with `div`

## mod

Perform integer modulo with `mod`. The first argument is the dividend and the
second is the divisor.

`mod 10 3` returns `1` (10 % 3).

When used in a pipeline, the piped value becomes the divisor:

```
{{ 3 | mod 10 }}
```

The above returns `1` (10 % 3).

## mul

Multiply with `mul`. Accepts two or more inputs.

```
mul 1 2 3
```

## max

Return the largest of a series of integers:

This will return `3`:

```
max 1 2 3
```

## min

Return the smallest of a series of integers.

`min 1 2 3` will return `1`

## floor

Returns the greatest float value less than or equal to input value

`floor 123.9999` will return `123.0`

## ceil

Returns the greatest float value greater than or equal to input value

`ceil 123.001` will return `124.0`

## round

Returns a float value with the remainder rounded to the given number to digits after the decimal point.

`round 123.555555 3` will return `123.556`

## randInt
Returns a random integer value from min (inclusive) to max (exclusive).

```
randInt 12 30
```

The above will produce a random number in the range [12,30].
