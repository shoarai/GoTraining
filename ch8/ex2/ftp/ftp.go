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

	fmt.Println("Connected client")
	ftpConn.reply("220", "FTP server ready")

	input := bufio.NewScanner(conn)
	for input.Scan() {
		text := input.Text()
		err := ftpConn.execute(text)
		if err != nil {
			ftpConn.error(err)
		}
	}
}

func (ftpConn *FtpConn) execute(text string) error {
	texts := strings.Split(text, " ")
	cmd := texts[0]
	args := texts[1:]
	// NOTE: Print log
	fmt.Printf("Command: %s, Args: %s\n", cmd, args)

	switch cmd {
	case "USER":
		ftpConn.login(args)
	case "CWD": // cd
		ftpConn.cd(args)
	case "PORT": // ls
		ftpConn.list(args)
	// case "LIST": // ls
	// None
	case "RETR": // get
		ftpConn.get(args)
	case "QUIT": // close
		ftpConn.close()
	default:
		return fmt.Errorf("Invalid command: %s", cmd)
	}

	return nil
}

func (ftpConn *FtpConn) login(args []string) {
	ftpConn.reply("230", "User logged in, proceed.")
}

func (ftpConn *FtpConn) cd(args []string) {
	if len(args) == 0 {
		args = []string{"."}
	}

	path := ftpConn.currentDir + args[0]
	info, err := os.Stat(path)
	if err != nil {
		ftpConn.error(err)
		return
	}

	if info.IsDir() {
		ftpConn.currentDir = path + "/"
	} else {
		ftpConn.reply("200", "not a directory")
	}

	ftpConn.reply("200", "Command okay.")
}

func (ftpConn *FtpConn) list(args []string) {
	path := ftpConn.currentDir
	info, err := os.Stat(path)
	if err != nil {
		// No such file or directory
		ftpConn.error(err)
		return
	}

	var fileNames []string
	if info.IsDir() {
		infos, err := ioutil.ReadDir(path)
		if err != nil {
			ftpConn.error(err)
			return
		}
		for _, info := range infos {
			fileNames = append(fileNames, info.Name())
		}
	} else {
		fileNames = append(fileNames, info.Name())
	}

	ftpConn.reply("200", strings.Join(fileNames, ", "))
}

func (ftpConn *FtpConn) get(args []string) {
	if len(args) == 0 {
		ftpConn.error(fmt.Errorf("no argument"))
		return
	}

	b, err := ioutil.ReadFile(args[0])
	if err != nil {
		ftpConn.error(err)
		return
	}
	ftpConn.reply("200", string(b))
}

func (ftpConn *FtpConn) close() {
	ftpConn.conn.Close()
}

func (ftpConn *FtpConn) reply(code, msg string) {
	fmt.Fprintf(ftpConn.conn, "%s %s\n", code, msg)
}

func (ftpConn *FtpConn) error(err error) {
	fmt.Fprintf(ftpConn.conn, "%s\n", err)
}
