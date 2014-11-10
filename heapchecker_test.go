package cgoheapchecker

import "testing"

func TestHeapChecker(t *testing.T) {
	if ok := NoGlobalLeaks(); !ok {
		t.Fatal("opps")
	}
	CancelGlobalCheck()
}

func TestHeapChecker_noop(t *testing.T) {
	if parent, err := ForkTest("HeapChecker_noop"); parent {
		if err != nil {
			t.Fatal(err)
		}
		return
	}
	// DO SOME OPERATIONS
	if ok := NoGlobalLeaks(); !ok {
		t.Fatal("opps, leak reported! it's false-alarm")
	}
	CancelGlobalCheck()
}
