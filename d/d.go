package d

import (
	"fmt"
	"os"
)

var _ = func() error {
	fmt.Println("d", os.Args)
	os.Args = []string{"d"}
	return nil
}()
