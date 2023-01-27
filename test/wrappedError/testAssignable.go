package main

import (
	"errors"
	"fmt"

	"github.com/AuruTus/Ergo/tools"
)

func main() {
	err := tools.WithRecover(func() {
		panic("test package visibility")
	})
	p := &tools.WrappedPanic{}
	ok := errors.As(err, &p)
	fmt.Printf("%t\n%s", ok, p)

	q := &tools.PathError{}
	ok = errors.As(err, &q)
	fmt.Printf("%t\n", ok)
}
