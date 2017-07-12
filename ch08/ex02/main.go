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
	conn *net.Conn
	cwd  string
}

// NewClient returns client
func NewClient(conn *net.Conn, cwd string) (c *Client) {
	c = new(Client)
	c.conn = conn
	c.cwd = cwd
	return c
}

func (c *Client) writeResponse(res string) {
	fmt.Fprintf(*c.conn, "%v\n", res)
}

func handleConn(c *Client) {
	fmt.Println("A client connected")
	fmt.Fprintf(*c.conn, "Connected to FTP server\n")

	scanner := bufio.NewScanner(*c.conn)

	var src, dst string
	for scanner.Scan() {
		cmd := scanner.Text()
		fmt.Printf("%v\n", cmd)
		cmds := strings.Split(cmd, " ")

		if len(cmds) == 0 {
			c.writeResponse("500 Syntax error, command unrecognized")
			continue
		}

		var err error
		switch cmds[0] {
		// RETR 取得
		case "RETR":
			// err = commandRETR(c)
			c.writeResponse("502 Command not implemented.")
		// STOR 保存
		case "STOR":
			// err = commandSTOR(c)
			c.writeResponse("502 Command not implemented.")
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
				break
			}
			src = cmds[1]
		// RNTO 名称変更先
		case "RNTO":
			if len(cmds) != 2 || cmds[1] == "" {
				c.writeResponse("500 Syntax error, command unrecognized")
				break
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
				break
			}
			err = commandDELE(c, cmds[1])
		// RMD  ディレクトリ削除
		case "RMD":
			if len(cmds) != 2 || cmds[1] == "" {
				c.writeResponse("500 Syntax error, command unrecognized")
				break
			}
			err = commandRMD(c, cmds[1])
		// MKD  ディレクトリ作成
		case "MKD":
			if len(cmds) != 2 || cmds[1] == "" {
				c.writeResponse("500 Syntax error, command unrecognized")
				break
			}
			err = commandMKD(c, cmds[1])
		// PWD  作業ディレクトリ表示
		case "PWD":
			err = commandPWD(c)
		// LIST 一覧
		case "LIST":
			err = commandLIST(c)
		// NLST 名前一覧
		case "NLST":
			commanNLST(c)
			c.writeResponse("502 Command not implemented.")
		// SITE サイト固有パラメータ
		case "SITE":
			err = commandSITE(c)
			c.writeResponse("502 Command not implemented.")
		// SYST システム
		case "SYST":
			err = commandSYST(c)
			c.writeResponse("502 Command not implemented.")
		// STAT ステータス
		case "STAT":
			err = commandSTAT(c)
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
		}
		// break if got error
		if err != nil {
			return
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
	fmt.Printf("HOME is set in %v\n", absHome)
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
