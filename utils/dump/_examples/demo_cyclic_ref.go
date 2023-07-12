package main

import (
	"time"

	"github.com/yoas0bi/micro-toolkit/utils/dump"
)

// rum demo:
//
//	go run ./dump/_examples/demo_cyclic_ref.go
func main() {
	a := map[string]any{}
	a["circular"] = map[string]any{
		"a": a,
	}

	// TIP: will stack overflow
	// fmt.Println(a)
	dump.V(a)

	val := map[string]any{
		"bool":   true,
		"number": 1 + 1i,
		"bytes":  []byte{97, 98, 99},
		"lines":  "first line\nsecond line",
		"slice":  []any{1, 2},
		"time":   time.Now(),
		"struct": struct{ test int32 }{
			test: 13,
		},
	}
	val["slice"].([]any)[1] = val["slice"]
	dump.P(val)
}
