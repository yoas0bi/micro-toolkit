package main

import "github.com/yoas0bi/micro-toolkit/utils/dump"

// rum demo:
//
//	go run ./dump/_examples/slice.go
func main() {
	dump.P(
		[]byte("abc"),
		[]int{1, 2, 3},
		[]string{"ab", "cd"},
		[]any{
			"ab",
			234,
			[]int{1, 3},
			[]string{"ab", "cd"},
		},
	)
}
