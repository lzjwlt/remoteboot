package rb

import (
	"net"
)

func tcpServer(msg string, ss ServerSocket) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ss.Addr+":"+string(ss.Port))
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		conn.Write([]byte(msg))
		conn.Close()
	}

}
