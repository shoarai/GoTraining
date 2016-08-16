// Copyright © 2016 shoarai

package ftp

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

type FtpConn struct {
	conn       net.Conn
	currentDir string
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	ftpConn := FtpConn{conn, "./"}

	fmt.Fprintf(conn, "> ")

	input := bufio.NewScanner(conn)
	for input.Scan() {
		text := input.Text()
		err := ftpConn.execute(text)
		if err != nil {
			fmt.Fprintln(conn, err)
		}

		fmt.Fprintf(conn, "> ")
	}
}

func (ftpConn *FtpConn) execute(text string) error {
	texts := strings.Split(text, " ")
	cmd := texts[0]
	args := texts[1:]

	fmt.Println(cmd, args)
	switch cmd {
	case "ls":
		ftpConn.list(args)
	case "cd":
		ftpConn.cd(args)
	default:
		return fmt.Errorf("Invalid command: %s", cmd)
	}

	return nil
}

func (ftpConn *FtpConn) cd(args []string) {
	if len(args) == 0 {
		args = []string{"."}
	}

	ftpConn.currentDir += args[0] + "/"
}

func (ftpConn *FtpConn) list(args []string) {
	if len(args) == 0 {
		args = []string{"."}
	}

	for _, arg := range args {
		path := ftpConn.currentDir + arg
		info, err := os.Stat(path)
		if err != nil {
			// No such file or directory
			fmt.Fprintf(ftpConn.conn, "%s\n", err)
			continue
		}

		if info.IsDir() {
			infos, err := ioutil.ReadDir(path)
			if err != nil {
				fmt.Fprintf(ftpConn.conn, "%s\n", err)
				continue
			}
			for _, info := range infos {
				fmt.Fprintf(ftpConn.conn, "%s\n", info.Name())
			}
		} else {
			fmt.Fprintf(ftpConn.conn, "%s\n", info.Name())
		}
	}
}
