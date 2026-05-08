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

##### Critical

Use `lang.Critical` to hold a mutex lock and execute a function.

```go
lang.Critical(lock, func() {
  shared += 1
})
```

Use `lang.Critical` to hold a mutex lock and execute a function with a return
value.

```go
v := lang.CriticalGet[int](lock, func() int {
  shared += 1
  return shared
})
```

##### Map

Use `lang.Map` to apply a function across all elements of a slice.

```go
conversion := lang.Map[int, string](originals, func(i int) string {
  return strconv.Itoa(i)
})
```

##### Filter / FilterFunc

Use `lang.Filter` or `lang.FilterFunc` to remove elements from a slice.

```go
filtered := lang.Filter[string](items, unwanted...)
```

```go
filtered := lang.FilterFunc[string](items, func(s string) bool {
  return someCondition(s)
})
```

##### Maybe

Use `lang.Maybe` as the ternary-operator for Go.

```go
passing := lang.Maybe[string](check, "ok", "fail")
```

##### Head

Use `lang.Head` to get a sub-slice of the first N elements of another slice. `lang.Head` will
return _up to_ N elements, gracefully handling slices that may have fewer than N elements.

```go
best := lang.Head[string](items, 3)
```

#### Tail

Use `lang.Tail` to get a sub-slice of the last N elements of another slice. `lang.Tail` will
return _up to_ N elements, gracefully handling slices that may have fewer than N elements.

```go
worst := lang.Tail[string](items, 3)
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
