package rb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"regexp"
)

var (
	delims = ":-"
	reMAC  = regexp.MustCompile(`^([0-9a-fA-F]{2}[` + delims + `]){5}([0-9a-fA-F]{2})$`)
)

type MACAddress [6]byte

type MagicPacket struct {
	header [6]byte
	body   [16]MACAddress
}

// New returns packet
func New(mac string) (*MagicPacket, error) {
	var packet MagicPacket
	var macAddr MACAddress

	hwAddr, err := net.ParseMAC(mac)

	if err != nil {
		return nil, err
	}

	if !reMAC.MatchString(mac) {
		return nil, fmt.Errorf("%s is not a MAC address", mac)
	}

	for idx := range macAddr {
		macAddr[idx] = hwAddr[idx]
	}

	for idx := range packet.body {
		packet.body[idx] = macAddr
	}

	return &packet, nil
}

func (mp *MagicPacket) Marshal() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, mp); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
