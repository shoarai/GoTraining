// Copyright © 2016 shoarai

package tar

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/shoarai/GoTraining/ch10/ex2/archive"
)

func init() {
	archive.RegistorFormat("tar", "\x75\x73\x74\x61\x72", read)
}

func read(file *os.File, out io.Writer) error {
	tr := tar.NewReader(file)

	// Iterate through the files in the archive.
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatalln(err)
		}
		fmt.Println()
	}

	return nil
}
