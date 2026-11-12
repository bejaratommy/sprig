# String Slice Functions

These function operate on or generate slices of strings. In Go, a slice is a
growable array. In Sprig, it's a special case of a `list`.

## join

Join a list of strings into a single string, with the given separator.

```
list "hello" "world" | join "_"
```

The above will produce `hello_world`

`join` will try to convert non-strings to a string value:

```
list 1 2 3 | join "+"
```

The above will produce `1+2+3`

## splitList and split

Split a string into a list of strings:

```
splitList "$" "foo$bar$baz"
```

The above will return `[foo bar baz]`

The older `split` function splits a string into a `dict`. It is designed to make
it easy to use template dot notation for accessing members:

```
$a := split "$" "foo$bar$baz"
```

The above produces a map with index keys. `{_0: foo, _1: bar, _2: baz}`

```
$a._0
```

The above produces `foo`

The result is a `dict` (`map[string]interface{}`), so it also works with dictionary
functions such as `get` and `hasKey`:

```
get $a "_0"
hasKey $a "_1"
```

These produce `foo` and `true`, respectively.

## splitn

`splitn` function splits a string into a `dict`. It is designed to make
it easy to use template dot notation for accessing members:

```
$a := splitn "$" 2 "foo$bar$baz"
```

The above produces a map with index keys. `{_0: foo, _1: bar$baz}`

```
$a._0
```

The above produces `foo`

The result is a `dict` (`map[string]interface{}`), so it also works with dictionary
functions such as `get` and `hasKey`:

```
get $a "_0"
hasKey $a "_1"
```

These produce `foo` and `true`, respectively.

## sortAlpha

The `sortAlpha` function sorts a list of strings into alphabetical (lexicographical)
order.

It does _not_ sort in place, but returns a sorted copy of the list, in keeping
with the immutability of lists.
