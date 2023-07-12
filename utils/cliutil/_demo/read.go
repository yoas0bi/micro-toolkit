package main

import (
	"fmt"

	"github.com/yoas0bi/micro-toolkit/utils/cliutil"
)

// go run ./_demo/read.go
func main() {
	ans, err := cliutil.ReadFirst("hi?")
	if err != nil {
		panic(err)
	}
	fmt.Println("ans:", ans)
}
