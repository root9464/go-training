package util

import (
	"fmt"
	"runtime"
	"time"
)

func CallerName(skip int) string {
	const unknown = "unknown"
	pcs := make([]uintptr, 1)
	n := runtime.Callers(skip+2, pcs)
	if n < 1 {
		return unknown
	}
	frame, _ := runtime.CallersFrames(pcs).Next()
	if frame.Function == "" {
		return unknown
	}
	return frame.Function
}

func Timer(fn func()) func() {
	name := CallerName(1)
	start := time.Now()
	return func() {
		fn()
		fmt.Printf("%s Время затраченное на выполнение: %v\n", name, time.Since(start))
	}
}

func Pench(fn func()) {
	defer Timer(fn)()
	fn()
}
