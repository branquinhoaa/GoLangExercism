package secret

import (
	"reflect"
	"testing"
)

const targetTestVersion = 1

var tests = []struct {
	code uint
	h    []string
}{
	{1, []string{"wink"}},                                             //1
	{2, []string{"double blink"}},                                     //10
	{4, []string{"close your eyes"}},                                  //100
	{8, []string{"jump"}},                                             //1000
	{3, []string{"wink", "double blink"}},                             //11
	{19, []string{"double blink", "wink"}},                            //10011
	{31, []string{"jump", "close your eyes", "double blink", "wink"}}, //11111
	{0, nil},               //0
	{32, nil},              //100000
	{33, []string{"wink"}}, //100001
}

func TestHandshake(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
	for _, test := range tests {
		h := Handshake(test.code)
		// use len() to allow either nil or empty list, because
		// they are not equal by DeepEqual
		if len(h) == 0 && len(test.h) == 0 {
			continue
		}
		if !reflect.DeepEqual(h, test.h) {
			t.Fatalf("Handshake(%d) = %q, want %q.", test.code, h, test.h)
		}
	}
}

func BenchmarkHandshake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Handshake(31)
	}
}
