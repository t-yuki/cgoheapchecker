package cgo

import (
	"testing"

	"github.com/t-yuki/cgoheapchecker"
)

func TestHeapCheck_noop(t *testing.T) {
	if parent, err := cgoheapchecker.ForkTest("HeapCheck_noop"); parent {
		if err != nil {
			t.Fatal(err)
		}
		return
	}
	if ok := cgoheapchecker.NoGlobalLeaks(); !ok {
		t.Fatal("opps, leaked!")
	}
	cgoheapchecker.CancelGlobalCheck()
}

func TestHeapCheck_leak(t *testing.T) {
	if parent, err := cgoheapchecker.ForkTest("HeapCheck_leak"); parent {
		if err != nil {
			t.Fatal(err)
		}
		return
	}
	CallMalloc() // 1 call not detected sometimes so call twice.
	CallMalloc()
	if ok := cgoheapchecker.NoGlobalLeaks(); ok {
		t.Fatal("want: leaked but not")
	}
	cgoheapchecker.CancelGlobalCheck()
}

func TestHeapCheck_noleak(t *testing.T) {
	if parent, err := cgoheapchecker.ForkTest("HeapCheck_noleak"); parent {
		if err != nil {
			t.Fatal(err)
		}
		return
	}
	CallMalloc()
	CallFree()
	if ok := cgoheapchecker.NoGlobalLeaks(); !ok {
		t.Fatal("opps, leaked!")
	}
	cgoheapchecker.CancelGlobalCheck()
}
