// Package cgoheapchecker provides simple wrapper functions of HeapLeakChecker in google perftools.
// It is intended to use with CGO library, not for pure golang library.
// Since the program that is imported this package also links tcmalloc, it should not be imported except tests.
// To use this package, install perftools in your system such as `apt-get install libgoogle-perftools-dev`.
//
// Note that, it is not guaranteed that HeapLeakChecker detects all memory leaks and it will miss small objects sometimes.
// For more details, see https://google-perftools.googlecode.com/svn/trunk/doc/heap_checker.html
package cgoheapchecker

// #cgo LDFLAGS: -Bstatic -ltcmalloc -Bdynamic
// #include "heapchecker.h"
import "C"
import (
	"os"
	"os/exec"
)

// ForkTest is a helper function of unit testing.
// It executes another test `name` using `go test -v -run=^TestNAME$` in sub-process with heap checker.
// If the child calls this function, it returns with false and nil without any operations.
// If the child reports errors, it returns with true and error.
func ForkTest(name string) (isparent bool, err error) {
	if os.Getenv("cgoheapchecker_"+name) != "" {
		return false, nil
	}
	os.Setenv("cgoheapchecker_"+name, "1")
	os.Setenv("HEAPCHECK", "normal")
	cmd := exec.Command(os.Args[0], "-test.v", "-test.run=^Test"+name+"$")
	_, err = cmd.CombinedOutput()
	if err != nil {
		return true, err
	}
	return true, nil
}

// NoGlobalLeaks calls HeapLeakChecker::NoGlobalLeaks.
func NoGlobalLeaks() bool {
	return C.NoGlobalLeaks() != 0
}

// CancelGlobalCheck calls HeapLeakChecker::CancelGlobalCheck.
func CancelGlobalCheck() {
	C.CancelGlobalCheck()
}
