// Copyright © 2016 shoarai

package ftp

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

// type FtpConnection struct {
// 	conn       net.Conn
// 	currentDir string
// }

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Fprintf(conn, "> ")

	input := bufio.NewScanner(conn)
	for input.Scan() {
		text := input.Text()
		err := execute(conn, text)
		if err != nil {
			fmt.Fprintln(conn, err)
		}

		fmt.Fprintf(conn, "> ")
	}
}

func execute(conn net.Conn, text string) error {
	texts := strings.Split(text, " ")
	cmd := texts[0]
	args := texts[1:]

	fmt.Println(cmd, args)
	switch cmd {
	case "ls":
		list(conn, args)
	default:
		return fmt.Errorf("Invalid command: %s", cmd)
	}

	return nil
}

func list(writer io.Writer, args []string) error {
	if len(args) == 0 {
		args = []string{"."}
	}

	for _, arg := range args {
		info, err := os.Stat(arg)
		if err != nil {
			// No such file or directory
			return err
		}

		if info.IsDir() {
			infos, err := ioutil.ReadDir(arg)
			if err != nil {
				return err
			}
			for _, info := range infos {
				fmt.Fprintf(writer, "%s\n", info.Name())
			}
		} else {
			fmt.Fprintf(writer, "%s\n", info.Name())
		}
	}
	return nil
}
