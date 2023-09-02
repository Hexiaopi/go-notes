package base

import "testing"

func ExamplePrintRecover() {
	PrintRecover()
	// Output:
	// a
	// some thing
}

func TestFailRecoverWithoutDefer(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()
	FailRecoverWithoutDefer()
}

func TestFailRecoverWithoutLevelDefer(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()
	FailRecoverWithoutLevelDefer()
}
