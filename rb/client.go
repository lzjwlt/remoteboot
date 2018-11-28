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
		fmt.Fprintf(os.Stderr, "error: %s\n", e.Error())
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

	messageChan := make(chan string, 10)

	go printServerChan(messageChan)
	go recvMessage(conn, messageChan, ip, port)
	for {
		time.Sleep(time.Second * 10)
	}

}

func printServerChan(messageChan chan string) {
	for {
		msg := <-messageChan
		if len(msg) > 1 {
			fmt.Println("~~Server Say:" + msg)
		}
	}
}

func recvMessage(conn net.Conn, messageChan chan string, ip string, port int) {
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			time.Sleep(time.Second * 10)
			tcpClient(ip, port)
		}
		msg := string(buf[0:n])
		messageChan <- msg

		if strings.HasPrefix(msg, "WAKE:") {
			mac := msg[5:]
			go handWake(conn, mac)
		}

	}
}

func handWake(conn net.Conn, mac string) {
	for i := 0; i < wakeTimes; i++ {
		err := wake(mac)
		checkError(err)
		fmt.Printf("waking: %s\n", mac)
	}
	_, err := conn.Write([]byte("OK:" + mac))
	checkError(err)
}
