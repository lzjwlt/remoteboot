package rb

import (
	"fmt"
	"log"
	"net"
)

const (
	udpPort int = 9
)

var broadcastIPs = []string{"255.255.255.255", "192.168.199.255"}

func sendBroadcast(bt []byte, broadcastIP string) error {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.ParseIP(broadcastIP), Port: udpPort})
	if err != nil {
		return err
	}
	defer conn.Close()

	n, err := conn.Write(bt)
	if err == nil && n != 102 {
		err = fmt.Errorf("magic packet sent was %d bytes (expected 102 bytes sent)", n)
	}
	if err != nil {
		return err
	}
	return nil
}

func wake(mac string) error {
	mp, err := New(mac)
	if err != nil {
		return err
	}
	bt, err := mp.Marshal()
	if err != nil {
		return err
	}
	for _, ip := range broadcastIPs {
		if err := sendBroadcast(bt, ip); err != nil {
			log.Println(err)
		}
	}
	return nil
}
