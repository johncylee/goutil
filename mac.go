package goutil

import (
	"fmt"
	"net"
)

// Wraps around net.HardwareAddr with JSON marshal/unmarshal support.
type MAC net.HardwareAddr

func (t *MAC) UnmarshalJSON(b []byte) error {
	if !IsJSONString(b) {
		return fmt.Errorf("Malformed JSON string: %s", string(b))
	}
	s := string(b[1 : len(b)-1])
	mac, err := net.ParseMAC(s)
	if err != nil {
		goto Error
	}
	*t = MAC(mac)
	return nil
Error:
	return fmt.Errorf("Malformed MAC: %s", s)
}

func (t MAC) MarshalJSON() (b []byte, err error) {
	return []byte(`"` + net.HardwareAddr(t).String() + `"`), nil
}

func (t *MAC) String() string {
	return net.HardwareAddr(*t).String()
}
