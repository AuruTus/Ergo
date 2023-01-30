package main

import (
	"errors"
	"fmt"

	"github.com/AuruTus/Ergo/pkg/utils"
)

func main() {
	err := utils.WithRecover(func() {
		panic("test package visibility")
	})
	p := &utils.WrappedPanic{}
	ok := errors.As(err, &p)
	fmt.Printf("%t\n%s", ok, p)

	q := &utils.PathError{}
	ok = errors.As(err, &q)
	fmt.Printf("%t\n", ok)
}
