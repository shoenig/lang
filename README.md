# lang

Package `lang` provides some helpful language features excluded from the Go language.

### Getting Started

The `lang` package can be added to a Go project with `go get`.

```shell
go get github.com/shoenig/lang@latest
```

```go
import "github.com/shoenig/lang"
```

### Examples

##### Maybe

Use `lang.Maybe` as the ternary-operator for Go.

```go
passing := lang.Maybe[string](check, "ok", "fail")
```

##### Head

Use `lang.Head` to get a sub-slice of the first N elements of another slice.

```go
best := lang.Head(items, 3)
```

##### Pairs

Use `lang.Pair` to conveniently associate two types together.

```go
p := lang.Pair[string, int]{A: "bob", B: 42}
```

### Contributing

Contributions are welcome! Feel free to help make `lang` better by filing an
issue or opening a PR.

## License

The `github.com/shoenig/lang` module is open source under the [BSD-3-Clause](LICENSE) license.
