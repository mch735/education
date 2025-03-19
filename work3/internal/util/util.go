package util

import (
	"fmt"
	"os"
)

func Fatal(err error) {
	fmt.Println(err) //nolint:forbidigo
	os.Exit(1)
}
