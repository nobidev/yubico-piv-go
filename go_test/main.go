package main

import (
	"fmt"

	ykpiv "github.com/nobidev/yubico-piv-go"
)

func main() {
	y := ykpiv.MustNew()

	y.MustRun(func() error {
		fmt.Printf("OK")
		return nil
	})
}
