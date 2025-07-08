package shscan

import (
	"io"
	"os"
)

func main() {
	os.Exit(scan(os.Stdin, os.Stderr))
}

func scan(w, ew io.Writer) int {
	return 0
}
