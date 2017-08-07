package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

func commandRETR(c *Client, fileName, fileType string) (err error) {
	c.dAccept()
	c.writeResponse("150 File status okay; about to open data connection.")
	file, err := os.Open(path.Join(c.cwd, fileName))
	if err != nil {
		c.writeResponse("500 File not found.")
		log.Println(err)
		return fmt.Errorf("%v", err)
	}
	if _, err := io.Copy(c.dconn, file); err != nil {
		log.Println(err)
		return fmt.Errorf("%v", err)
	}
	c.dClose("RETR")
	return nil
}

func commandSTOR(c *Client, fileName, fileType string) (err error) {
	c.dAccept()
	c.writeResponse("150 File status okay; about to open data connection.")
	file, err := os.Create(path.Join(c.cwd, fileName))
	if err != nil {
		log.Println(err)
		return fmt.Errorf("%v", err)
	}
	if _, err := io.Copy(file, c.dconn); err != nil {
		log.Println(err)
		return fmt.Errorf("%v", err)
	}
	c.dClose("STOR")
	return
}

func parseIPPort(arg string) (ip, port string, err error) {
	trimed := strings.Trim(arg, " ")
	splited := strings.Split(trimed, ",")
	if len(splited) != 6 {
		return "", "", fmt.Errorf("sytax error")
	}
	ip = fmt.Sprintf("%s.%s.%s.%s", splited[0], splited[1], splited[2], splited[3])
	a, err := strconv.ParseInt(splited[4], 10, 64)
	if err != nil {
		return "", "", fmt.Errorf("sytax error")
	}
	b, err := strconv.ParseInt(splited[5], 10, 64)
	if err != nil {
		return "", "", fmt.Errorf("sytax error")
	}
	port = fmt.Sprintf("%d", a*256+b)
	return ip, port, nil
}

func commandPORT(c *Client, ip, port string) (err error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		fmt.Println(err)
		return err
	}
	c.dconn = conn
	c.writeResponse("200 Command okay.")
	return
}

func commandPASV(c *Client) (err error) {
	if c.dlistener == nil {
		listener, err := net.Listen("tcp", "localhost:12345")
		if err != nil {
			return err
		}
		c.dlistener = listener
	}
	c.writeResponse("227 Entering Passive Mode (127,0,0,1,48,57).")
	return
}

func commandSTOU(c *Client) (err error) {
	c.writeResponse("200 Command okay.")
	return
}

func commandAPPE(c *Client) (err error) {
	c.writeResponse("200 Command okay.")
	return
}

func commandALLO(c *Client) (err error) {
	c.writeResponse("200 Command okay.")
	return
}

func commandREST(c *Client) (err error) {
	c.writeResponse("200 Command okay.")
	return
}

func commandRNTO(c *Client, src, dst string) (err error) {
	if err = os.Rename(path.Join(c.cwd, src), path.Join(c.cwd, dst)); err != nil {
		c.writeResponse("550 Requested action not taken.")
		return
	}
	c.writeResponse("200 Command okay.")
	return
}

func commandABOR(c *Client) (err error) {
	c.writeResponse("200 Command okay.")
	return
}

func commandDELE(c *Client, file string) (err error) {
	if err = os.Remove(file); err != nil {
		c.writeResponse("550 Requested action not taken.")
		return
	}
	c.writeResponse("200 Command okay.")
	return
}

func commandRMD(c *Client, dir string) (err error) {
	if err = os.Remove(path.Join(c.cwd, dir)); err != nil {
		c.writeResponse("550 Requested action not taken.")
		return
	}
	c.writeResponse("200 Command okay.")
	return
}

func commandMKD(c *Client, dir string) (err error) {
	if err = os.MkdirAll(path.Join(c.cwd, dir), 0777); err != nil {
		c.writeResponse("550 Requested action not taken.")
		return
	}
	c.writeResponse("200 Command okay.")
	return
}

func commandPWD(c *Client) (err error) {
	c.writeResponse(fmt.Sprintf("257 \"%s\" created.", c.cwd))
	return
}

func commandCWD(c *Client, p string) (err error) {
	// absolute path or relative path
	if strings.HasPrefix(p, "/") {
		if _, err = os.Stat(p); err != nil {
			c.writeResponse("550 Requested action not taken.")
			return
		}
		c.cwd = p
		c.cwd = path.Join(c.cwd, p)
		c.writeResponse("200 Command okay.")
	} else {
		if _, err = os.Stat(path.Join(c.cwd, p)); err != nil {
			c.writeResponse("550 Requested action not taken.")
			return
		}
		c.cwd = path.Join(c.cwd, p)
		c.writeResponse("200 Command okay.")
	}
	return
}

func commandLIST(c *Client) (err error) {
	c.dAccept()
	c.writeResponse("150 File status okay; about to open data connection.")
	files, err := ioutil.ReadDir(c.cwd)
	if err != nil {
		c.writeResponse("550 Requested action not taken.")
		return
	}
	var str string
	for i, file := range files {
		str += " " + file.Name()
		if i%5 == 4 {
			str += "\n"
		}
	}
	log.Println(str)
	c.writeDResponse(str)
	c.dClose("LIST")
	return
}

func commanNLST(c *Client) (err error) {
	c.writeResponse("200 Command okay.")
	return
}

func commandSITE(c *Client, cmds []string) (err error) {
	if len(cmds) == 1 {
		_, err = exec.Command(cmds[0]).Output()
	} else {
		_, err = exec.Command(cmds[0], cmds[1:]...).Output()
	}
	if err != nil {
		c.writeResponse("500 Syntax error, command unrecognized")
		return
	}
	c.writeResponse("200 Command okay.")
	return
}

func commandSYST(c *Client) (err error) {
	c.writeResponse("215 UNIX system type.")
	return
}

func commandSTAT(c *Client) (err error) {
	// c.writeResponse("200 Command okay.")
	return
}

func commandHELP(c *Client) (err error) {
	c.writeResponse("200 Command okay.")

	commands := []string{
		"RETR",
		"STOR",
		"STOU",
		"APPE",
		"ALLO",
		"REST",
		"RNFR",
		"RNTO",
		"ABOR",
		"DELE",
		"RMD",
		"MKD",
		"PWD",
		"LIST",
		"NLST",
		"SITE",
		"SYST",
		"STAT",
		"HELP",
		"QUIT",
		"NOOP",
	}
	c.writeResponse(strings.Join(commands, "\n"))
	return
}

func commandNOOP(c *Client) (err error) {
	c.writeResponse("200 Command okay.")
	return
}
