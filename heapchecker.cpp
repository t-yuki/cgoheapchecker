#include <google/heap-checker.h>
#include "heapchecker.h"

int NoGlobalLeaks() {
	return HeapLeakChecker::NoGlobalLeaks();
}

void CancelGlobalCheck() {
	HeapLeakChecker::CancelGlobalCheck();
}
