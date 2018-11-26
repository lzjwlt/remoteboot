package rb

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

type ServerSocket struct {
	Addr string
	Port int
}

func checkError(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "error: %s", e.Error())
	}
}

var testServer = ServerSocket{Addr: "111.111.111.111", Port: 1111}

func tcpClient(msg string, ss ServerSocket) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ss.Addr+":"+string(ss.Port))
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	checkError(err)
	_, err = conn.Write([]byte(msg))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	_ = result

}
