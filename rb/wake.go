package rb

import (
	"fmt"
	"net"
)

const (
	broadcastIP string = "255.255.255.255"
	udpPort     int    = 9
)

func sendBroadcast(bt []byte) error {
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
	if err := sendBroadcast(bt); err != nil {
		return err
	}
	return nil
}
