// Copyright © 2016 shoarai

package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"

	"github.com/shoarai/GoTraining/ch10/ex2/archive"
)

func init() {
	archive.RegistorFormat("zip", "\x50\x4b", read)
}

func read(file *os.File, out io.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}

	// Open a zip archive for reading.
	r, err := zip.NewReader(file, info.Size())
	if err != nil {
		return err
	}

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			return err
		}
		_, err = io.CopyN(out, rc, 68)
		if err != nil {
			return err
		}
		rc.Close()
		fmt.Println()
	}

	return nil
}
