package backtrace

import (
	"html/template"
	"runtime"
	"strings"
)

const DefaultBacktraceLimit = 10

type Backtrace []Trace

type Trace struct {
	File     string
	Function string
	Line     int
}

func DebugBacktrace(options ...int) Backtrace {
	backtrace := make(Backtrace, 0)

	limit := DefaultBacktraceLimit

	if len(options) > 0 {
		limit = options[0]
	}

	pc := make([]uintptr, limit)
	n := runtime.Callers(0, pc)
	if n == 0 {
		return backtrace
	}

	pc = pc[:n]

	frames := runtime.CallersFrames(pc)

	for {
		frame, more := frames.Next()

		if strings.Contains(frame.File, "runtime/") {
			continue
		}

		if !more {
			break
		}

		backtrace = append(backtrace, Trace{
			File:     frame.File,
			Function: frame.Function,
			Line:     frame.Line,
		})
	}

	return backtrace
}

func DebugPrintBacktrace(backtrace Backtrace) string {
	tmpl, err := template.New("backtrace").Parse("File: {{.File}}, Function: {{.Function}}, Line: {{.Line}} \n")
	if err != nil {
		panic(err)
	}

	builder := &strings.Builder{}

	for _, trace := range backtrace {
		err = tmpl.Execute(builder, trace)
		if err != nil {
			panic(err)
		}
	}

	return builder.String()
}
