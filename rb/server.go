package rb

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type lockedConnsMap struct {
	data map[string]net.Conn
	lock sync.Mutex
}

func (lcm lockedConnsMap) get(k string) net.Conn {
	lcm.lock.Lock()
	defer lcm.lock.Unlock()
	return lcm.data[k]
}
func (lcm lockedConnsMap) set(k string, v net.Conn) {
	lcm.lock.Lock()
	defer lcm.lock.Unlock()
	lcm.data[k] = v
}
func (lcm lockedConnsMap) del(k string) {
	lcm.lock.Lock()
	defer lcm.lock.Unlock()
	delete(lcm.data, k)
}

// 启动地址
var addrsToBoot []string

// Start Server
func StartServer(port int) {
	tcpServer(port)
}

func tcpServer(port int) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf(":%d", port))
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// 连接池
	conns := make(map[string]net.Conn)
	lcm := lockedConnsMap{data: conns}

	// 消息通道
	messageChan := make(chan string, 10)
	go input()

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		lcm.set(conn.RemoteAddr().String(), conn)
		//conns[conn.RemoteAddr().String()] = conn
		go handler(conn, &lcm, messageChan)
		go broadcast(&lcm)
		go printClientChan(messageChan)
	}
}
func printClientChan(messageChan chan string) {
	for {
		msg := <-messageChan
		if len(msg) > 1 {
			fmt.Println(msg)
		}
	}
}

func broadcast(lcm *lockedConnsMap) {

	for {
		for key, conn := range lcm.data {
			for _, mac := range addrsToBoot {
				_, err := conn.Write([]byte("WAKE:" + mac))
				if err != nil {
					fmt.Fprintf(os.Stderr, "broadcast to %s faild %s: %s\n", key, mac, err)
					lcm.del(key)
				}
			}
		}
		time.Sleep(10 * time.Second)
		fmt.Println("***** to wake list ****")
		for k, v := range addrsToBoot {
			fmt.Printf("** %d  ---- %s **\n", k, v)
		}
		for k, v := range lcm.data {
			fmt.Printf("** CONN ---- %s **\n", k)
			_ = v
		}

	}

}

func handler(conn net.Conn, lcm *lockedConnsMap, messages chan string) {
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			//delete(*conns, conn.RemoteAddr().String())
			lcm.del(conn.RemoteAddr().String())
			conn.Close()
		}
		rcvString := string(buf[0:n])
		if strings.HasPrefix(rcvString, "OK:") {
			mac := rcvString[3:]
			fmt.Printf("wait %s success, remove it from todos..\n", mac)
			for k, v := range addrsToBoot {
				if strings.Contains(mac, v) {
					addrsToBoot = append(addrsToBoot[:k], addrsToBoot[k+1:]...)
					fmt.Printf("delete %s finished\n", v)
				}
			}

		}
		messages <- fmt.Sprintf("~~%s Say: %s", conn.RemoteAddr().String(), rcvString)
	}
}

func input() {
	var opt string
	var obj string
	for {
		fmt.Scanln(&opt, &obj)
		switch opt {
		case "my":
			if obj == "1" {
				mac := "00:e0:70:1b:77:92"
				addrsToBoot = append(addrsToBoot, mac)
				fmt.Printf("to wake %s\n", mac)
			}

		case "wake":

			addrsToBoot = append(addrsToBoot, obj)
			fmt.Printf("to wake %s\n", obj)

		case "delete":
			for k, v := range addrsToBoot {
				if strings.Contains(v, obj) {
					addrsToBoot = append(addrsToBoot[:k], addrsToBoot[k+1:]...)
					fmt.Printf("delete %s finished\n", v)
					break
				}
			}
		default:
			fmt.Fprintf(os.Stderr, "error: %s, unknow cmd\n", opt)
		}

	}

}
