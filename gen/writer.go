package gen

import (
	"bufio"
	"io"
	"os"
)

func writeToFile(file string, structs structs) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	defer w.Flush()

	return write(w, structs)
}

func write(w io.Writer, structs structs) error {
	return nil
}
