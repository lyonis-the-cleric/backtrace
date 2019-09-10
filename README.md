Parameters:
Function `backtrace.DebugBacktrace` can take limit int parameter (max depth)

```go
backtrace.DebugBacktrace(30)
// or
backtrace.DebugBacktrace()
```


Usage:

```go
b := backtrace.DebugPrintBacktrace(backtrace.DebugBacktrace())

fmt.Println(b)
```

Example output:
```
File: /home/lyonis/go/src/github.com/lyonis-the-cleric/backtrace/backtrace.go, Function: github.com/lyonis-the-cleric/backtrace.DebugBacktrace, Line: 29 
File: /home/lyonis/go/src/debug-backtrace/main.go, Function: main.f2, Line: 21 
File: /home/lyonis/go/src/debug-backtrace/main.go, Function: main.f1, Line: 16 
File: /home/lyonis/go/src/debug-backtrace/main.go, Function: main.main, Line: 11 
```
