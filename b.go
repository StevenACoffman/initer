package main

import (
	"fmt"
	"os"
)

var _ = func() error {
	fmt.Println("b", os.Args)
	os.Args = []string{"b"}
	return nil
}()
