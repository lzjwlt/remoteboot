package daemon

import (
	"fmt"
	"log"
	"net"
)

const (
	udpPort int = 9
)

func getLocalIPs() []string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	var IPs []string
	for _, address := range addrs {

		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				gInnerIP := ipnet.IP.String()
				IPs = append(IPs, gInnerIP)
			}
		}
	}
	return IPs
}

const broadcastIP = "255.255.255.255"

func sendBroadcast(bt []byte) error {
	for _, ip := range getLocalIPs() {
		conn, err := net.DialUDP("udp", &net.UDPAddr{IP: net.ParseIP(ip)}, &net.UDPAddr{IP: net.ParseIP(broadcastIP), Port: udpPort})
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
		log.Println(err)
	}

	return nil
}
