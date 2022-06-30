# expression

Another expression parser and executor
Better version of https://github.com/neonxp/lexpr

## usage

```go
import "go.neonxp.dev/expression"
...
e := expression.New()
result, err := e.Eval(`2 + 2`) // 4, nil
...
```

## defaults

Default operators and functions: [/defaults.go](/defaults.go)
