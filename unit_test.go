package goutil

import (
	"bytes"
	"encoding/json"
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

func TestMAC(t *testing.T) {
	b := []byte(`"01:23:45:67:89:ab"`)
	var mac MAC
	err := json.Unmarshal(b, &mac)
	if err != nil {
		t.Error(err)
		return
	}
	d, err := json.Marshal(mac)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(b, d) {
		t.Errorf("Not match: %s , %s", b, d)
	}
}
