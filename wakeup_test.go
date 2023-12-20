package wakeup

import (
	"testing"
	"time"
)

func requireTrue(t *testing.T, val bool) {
	if !val {
		t.Error("true expected")
	}
}

func requireNoError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func requireEqual(t *testing.T, leftAndRight ...int) {
	if len(leftAndRight)%2 != 0 {
		panic("len(leftAndRight)%2 != 0")
	}

	for i := 0; i < len(leftAndRight); i += 2 {
		if leftAndRight[i] != leftAndRight[i+1] {
			t.Errorf("left != right, %d != %d", leftAndRight[i], leftAndRight[i+1])
		}
	}
}

// NO JOKES, THIS TEST WORKS ONLY IF YOU RUN IT BEFORE ITS 23:22:21 ON YOUR CLOCK.
// IF IT'S TOO LATE, YOU SHOULD HAVE A REST. <3
// (actually, it would work either way, I just felt funny, sorry, it wouldn't happen again)
func TestDuration(t *testing.T) {

	h, m, s, err := Parse("23:22:21")
	requireNoError(t, err)

	now := time.Now()
	wakeup := Time(h, m, s)
	requireEqual(t, wakeup.Hour(), h, wakeup.Minute(), m, wakeup.Second(), s)

	requireTrue(t, Time(h, m, s).After(now))
}
