package goutil

import (
	"encoding/binary"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"math"
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

var bsonBinaryKind byte = 0x05

// Save as binary
func (t MAC) GetBSON() (interface{}, error) {
	mac := []byte(t)
	l := len(mac)
	if l > math.MaxInt32 {
		return nil, fmt.Errorf("Length overflow int32: %d", l)
	}
	data := make([]byte, 4, 4+1+l)
	binary.LittleEndian.PutUint32(data, uint32(l))
	data = append(data, byte(0x00))
	data = append(data, mac...)
	return bson.Raw{
		Kind: bsonBinaryKind,
		Data: data,
	}, nil
}

func (t *MAC) SetBSON(raw bson.Raw) error {
	if raw.Kind != bsonBinaryKind {
		return fmt.Errorf("Kind 0x%x, expect 0x05", raw.Kind)
	}
	totalLen := uint64(len(raw.Data))
	if totalLen < 6 {
		return fmt.Errorf("Data too short: %d, at least 6", totalLen)
	}
	l := binary.LittleEndian.Uint32(raw.Data[0:4])
	if totalLen != uint64(4+1+l) {
		return fmt.Errorf("Incorrect length %d, expect %d", totalLen, 4+1+l)
	}
	*t = MAC(raw.Data[5:])
	return nil
}
