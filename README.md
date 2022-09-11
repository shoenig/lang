# lang

Package `lang` provides some helpful or at least amusing language features
not included in the standard Go language or libraries.

[![Run CI Tests](https://github.com/shoenig/lang/actions/workflows/ci.yaml/badge.svg)](https://github.com/shoenig/lang/actions/workflows/ci.yaml)

# Motivation

The helpers provided by package `lang` are inspired by the lack of some basic
language features in the Go language or libraries.

### If/Else

Say we want to conditionally set the value of a string - one that could be a const
in another language, written only on initialization, but conditionally.

e.g.

```C#
final String s = (true) ? "foo" : "bar"
```

In Go we cannot do this, instead we must write something like,

```golang
s := "foo"
if true {
    s = "bar"
}
```

This grinds my gears. Why waste 4 lines of code on something that could be 1 line?

Now, we can!

```golang
s := lang.If(true, "foo").Else("bar")
```

## Contributing

Contributions are welcome! Feel free to help make `lang` better by filing an
issue or opening a PR.

## License

The `github.com/shoenig/lang` module is open source under the [MPL-2.0](LICENSE) license.
