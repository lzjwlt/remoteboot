package rb

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

const (
	wakeTimes   int           = 3
	wakeWaiting time.Duration = time.Second * 10
)

func checkError(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "error: %s", e.Error())
	}
}

// StartClient ...
func StartClient(ip string, port int) {
	tcpClient(ip, port)
}

func tcpClient(ip string, port int) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", ip, port))
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	checkError(err)
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			time.Sleep(time.Second * 10)
			tcpClient(ip,port)
		}
		result := string(buf[0:n])
		fmt.Println(string(buf[0:n]))
		if strings.HasPrefix(result, "WAKE:") {
			mac := result[5:]
			for i := 0; i < wakeTimes; i++ {
				wake(mac)
				_, err = conn.Write([]byte("waking:" + mac))
			}
			_, err = conn.Write([]byte("OK:" + mac))
			checkError(err)
		}

	}

}
