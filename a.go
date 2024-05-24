package main

import (
	"fmt"
	_ "github.com/StevenACoffman/initer/d"
	_ "github.com/StevenACoffman/initer/e"
	"os"
)

var _ = func() error {
	fmt.Println("a", os.Args)
	os.Args = []string{"a"}
	return nil
}()

func main() {
	fmt.Println("main")
	fmt.Println("### os.Args ->", os.Args)
}
