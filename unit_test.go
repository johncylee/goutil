package goutil

import (
	"bytes"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"net"
	"testing"
)

type testCase struct {
	data   []byte
	expect bool
}

func TestIsJSONString(t *testing.T) {
	tests := []testCase{
		{data: []byte(`foo`), expect: false},
		{data: []byte(`"foo"`), expect: true},
		{data: []byte(`""foo"`), expect: false},
		{data: []byte(`"\"foo"`), expect: true},
	}
	for _, test := range tests {
		if IsJSONString(test.data) != test.expect {
			t.Errorf("data `%s` expect %v", test.data, test.expect)
		}
	}
}

func TestMACJSON(t *testing.T) {
	b := []byte(`"01:23:45:67:89:ab"`)
	var m1 MAC
	err := json.Unmarshal(b, &m1)
	if err != nil {
		t.Fatal(err)
	}
	d, err := json.Marshal(m1)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, d) {
		t.Fatalf("Not equal: %v , %v", b, d)
	}
}

type MACstruct struct {
	MAC MAC
}

func TestMACBSON(t *testing.T) {
	mac, err := net.ParseMAC("01:23:45:67:89:ab")
	if err != nil {
		t.Fatal(err)
	}
	doc1 := MACstruct{
		MAC: MAC(mac),
	}
	b, err := bson.Marshal(doc1)
	if err != nil {
		t.Fatal(err)
	}
	var doc2 MACstruct
	err = bson.Unmarshal(b, &doc2)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal([]byte(doc1.MAC), []byte(doc2.MAC)) {
		t.Fatalf("Not equal: %v , %v", doc1.MAC, doc2.MAC)
	}
}
