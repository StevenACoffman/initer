package e

import (
	"fmt"
	"os"
)

var _ = func() error {
	fmt.Println("e", os.Args)
	os.Args = []string{"e"}
	return nil
}()
