package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

func commandRETR(c *Client) (err error) {
	c.writeResponse("200 Command okay.")
	return
}

func commandSTOR(c *Client) (err error) {
	c.writeResponse("200 Command okay.")
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
	c.writeResponse("200 Command okay.")
	c.writeResponse(fmt.Sprintf("%s", c.cwd))
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
	files, err := ioutil.ReadDir(c.cwd)
	if err != nil {
		c.writeResponse("550 Requested action not taken.")
		return
	}
	c.writeResponse("200 Command okay.")
	var str string
	for _, file := range files {
		str += "\n" + file.Name()
	}
	c.writeResponse(str)
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
	c.writeResponse("System: FTP Server system ex02\n")
	c.writeResponse("200 Command okay.")
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
