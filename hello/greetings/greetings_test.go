package greetings

import "testing"

func TestHello(t *testing.T) {
	want := "Hi, Telman. Welcome!"
	if got, err := Hello("Telman"); got != want || err != nil {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
