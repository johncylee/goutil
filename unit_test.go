package goutil

import (
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
