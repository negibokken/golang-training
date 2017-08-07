package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

// Client is store connection and cwd
type Client struct {
	conn  *net.Conn
	cwd   string
	dconn net.Conn
	mode  string
}

// NewClient returns client
func NewClient(conn *net.Conn, cwd string) (c *Client) {
	c = new(Client)
	c.conn = conn
	c.cwd = cwd
	c.mode = "PASV"
	c.dconn = nil
	return c
}

func (c *Client) writeResponse(res string) {
	fmt.Fprintf(*c.conn, "%v\n", res)
}

func (c *Client) writeDResponse(res string) {
	fmt.Fprintf(c.dconn, "%v\n", res)
}

func (c *Client) dClose() {
	c.dconn.Close()
	c.writeResponse("226 Closing data connection. List successful")
}

func handleConn(c *Client) {
	log.Println("A client connected")

	c.writeResponse("220 Ready.")
	scanner := bufio.NewScanner(*c.conn)
	var src, dst string
	for scanner.Scan() {
		cmd := scanner.Text()
		log.Printf("%v\n", cmd)
		cmds := strings.Split(cmd, " ")

		// default value
		var fileType = "A"

		if len(cmds) == 0 {
			c.writeResponse("500 Syntax error, command unrecognized")
			continue
		}

		upperCMD := strings.ToUpper(cmds[0])
		var err error
		switch upperCMD {
		// USER 設定
		case "USER":
			if len(cmds) != 2 {
				c.writeResponse("500 Syntax error, command unrecognized")
				continue
			}
			c.writeResponse("230 Login successful.")
		// RETR 取得
		case "RETR":
			// err = commandRETR(c)
			if c.dconn == nil {
				c.writeResponse("503 Bad sequence of commands.")
				continue
			}
			if len(cmds) != 2 || cmds[1] == "" {
				c.writeResponse("500 Syntax error, command unrecognized")
				continue
			}
			err = commandRETR(c, cmds[1], fileType)
			// CDUP 一つ上のディレクトリへ
		case "CDUP":
			err = commandCWD(c, "..")
		case "TYPE":
			if len(cmds) != 2 {
				c.writeResponse("500 Syntax error, command unrecognized")
				continue
			}
			if cmds[1] == "A" {
				fileType = "A"
			} else if cmds[1] == "I" {
				fileType = "I"
			}
			c.writeResponse("200 Command okay.")
		case "PORT":
			if len(cmds) != 2 {
				c.writeResponse("500 Syntax error, command unrecognized")
				continue
			}
			ip, port, err := parseIPPort(cmds[1])
			if err != nil {
				c.writeResponse("500 Syntax error, command unrecognized")
				continue
			}
			err = commandPORT(c, ip, port)
		case "PASV":
			commandPASV(c)
		// STOR 保存
		case "STOR":
			if c.dconn == nil {
				c.writeResponse("503 Bad sequence of commands.")
				continue
			}
			if len(cmds) != 2 || cmds[1] == "" {
				c.writeResponse("500 Syntax error, command unrecognized")
				continue
			}
			err = commandSTOR(c, cmds[1], fileType)
		// STOU 一時保存
		case "STOU":
			// err = commandSTOU(c)
			c.writeResponse("502 Command not implemented.")
		// APPE	追加（生成）
		case "APPE":
			// err = commandAPPE(c)
			c.writeResponse("502 Command not implemented.")
		// ALLO 割当
		case "ALLO":
			c.writeResponse("502 Command not implemented.")
		// REST 再開
		case "REST":
			c.writeResponse("502 Command not implemented.")
		// RNFR 名称変更元
		case "RNFR":
			if len(cmds) != 2 || cmds[1] == "" {
				c.writeResponse("500 Syntax error, command unrecognized")
				continue
			}
			c.writeResponse("200 Command okay.")
			src = cmds[1]
		// RNTO 名称変更先
		case "RNTO":
			if len(cmds) != 2 || cmds[1] == "" {
				c.writeResponse("500 Syntax error, command unrecognized")
				continue
			}
			dst = cmds[1]
			err = commandRNTO(c, src, dst)
			// reset
			src, dst = "", ""
		// ABOR 中断
		case "ABOR":
			err = commandABOR(c)
		// DELE 削除
		case "DELE":
			if len(cmds) != 2 || cmds[1] == "" {
				c.writeResponse("500 Syntax error, command unrecognized")
				continue
			}
			err = commandDELE(c, cmds[1])
		// RMD  ディレクトリ削除
		case "RMD":
			if len(cmds) != 2 || cmds[1] == "" {
				c.writeResponse("500 Syntax error, command unrecognized")
				continue
			}
			err = commandRMD(c, cmds[1])
		// MKD  ディレクトリ作成
		case "MKD":
			if len(cmds) != 2 || cmds[1] == "" {
				c.writeResponse("500 Syntax error, command unrecognized")
				continue
			}
			err = commandMKD(c, cmds[1])
		// PWD  作業ディレクトリ表示
		case "PWD":
			err = commandPWD(c)
		case "CWD":
			if len(cmds) != 2 || cmds[1] == "" {
				c.writeResponse("500 Syntax error, command unrecognized")
				continue
			}
			err = commandCWD(c, cmds[1])
		// LIST 一覧
		case "LIST":
			err = commandLIST(c)
		// NLST 名前一覧
		case "NLST":
			// commanNLST(c)
			c.writeResponse("502 Command not implemented.")
		// SITE サイト固有パラメータ
		case "SITE":
			if len(cmds) <= 1 {
				c.writeResponse("500 Syntax error, command unrecognized")
				continue
			}
			err = commandSITE(c, cmds[1:])
		// SYST システム
		case "SYST":
			err = commandSYST(c)
		// STAT ステータス
		case "STAT":
			// err = commandSTAT(c)
			c.writeResponse("502 Command not implemented.")
		// HELP ヘルプ
		case "HELP":
			err = commandHELP(c)
		case "QUIT":
			c.writeResponse("221 Service closing control connection.")
			return
		// NOOP NOOP
		case "NOOP":
			err = commandNOOP(c)
		// other
		default:
			c.writeResponse(fmt.Sprintf("502 Command %q not implemented.", cmds[0]))
		}
		if err != nil {
			/* error process */
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "./ex02 <ftp_home_dir>")
		return
	}
	home := os.Args[1]
	absHome, err := filepath.Abs(home)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}
	log.Printf("HOME is set in %v\n", absHome)
	listener, err := net.Listen("tcp", "localhost:21")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("%v\n", err)
			conn.Close()
			return
		}
		defer conn.Close()
		c := NewClient(&conn, absHome)
		go handleConn(c)
	}
}
