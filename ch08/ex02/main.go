package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var protocolINterpreterClose chan struct{}
var dataTransProcessClose chan struct{}

func commandList() {
	// RETR 取得
	// STOR 保存
	// STOU 一位保存
	// APPE	追加（生成）
	// ALLO 割当
	// REST 再開
	// RNFR 名称変更元
	// RNTO 名称変更先
	// ABOR 中断
	// DELE 削除
	// RMD  ディレクトリ削除
	// MKD  ディレクトリ作成
	// PWD  作業ディレクトリ表示
	// LIST 一覧
	// NLST 名前一覧
	// SITE サイト固有パラメータ
	// SYST システム
	// STAT ステータス
	// HELP ヘルプ
	// NOOP NOOP
}

func protocolInterpreter() {

}

func dataTransferProcess() {

}

func handleConn(conn net.Conn) {
	fmt.Println("A client connected")
	fmt.Fprintf(conn, "Connected to FTP server\n")

	scanner := bufio.NewScanner(conn)
	fmt.Print("CMD>")
	for scanner.Scan() {
		cmd := scanner.Text()
		fmt.Printf("%v\n", cmd)
		fmt.Print("CMD>")
	}
}

func main() {
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
		go handleConn(conn)
	}
}
